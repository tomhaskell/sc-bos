package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/smart-core-os/sc-golang/pkg/middleware/name"
	"github.com/timshannon/bolthold"
	"github.com/vanti-dev/sc-bos/pkg/auth/token"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/vanti-dev/sc-bos/internal/manage/devices"
	"github.com/vanti-dev/sc-bos/pkg/app/services"
	"github.com/vanti-dev/sc-bos/pkg/task/service"

	"github.com/vanti-dev/sc-bos/internal/util/pki"
	"github.com/vanti-dev/sc-bos/internal/util/pki/expire"
	"github.com/vanti-dev/sc-bos/pkg/auto"
	"github.com/vanti-dev/sc-bos/pkg/driver"
	"github.com/vanti-dev/sc-bos/pkg/manage/enrollment"
	"github.com/vanti-dev/sc-bos/pkg/node"
	"github.com/vanti-dev/sc-bos/pkg/system"
	"github.com/vanti-dev/sc-bos/pkg/task"

	"github.com/vanti-dev/sc-bos/pkg/auth/policy"
	"github.com/vanti-dev/sc-bos/pkg/gen"
)

const LocalConfigFileName = "area-controller.local.json"

type SystemConfig struct {
	Logger      zap.Config
	ListenGRPC  string
	ListenHTTPS string

	DataDir             string
	StaticDir           string // hosts static files from this directory over HTTP if StaticDir is non-empty
	LocalConfigFileName string // defaults to LocalConfigFileName
	CertConfig          CertConfig

	Policy        policy.Policy // Override the policy used for RPC calls. Defaults to policy.Default
	DisablePolicy bool          // Unsafe, disables any policy checking for the server

	DriverFactories map[string]driver.Factory // keyed by driver name
	AutoFactories   map[string]auto.Factory   // keyed by automation type
	SystemFactories map[string]system.Factory // keyed by system type
}

// CertConfig encapsulates different settings used for loading and present certificates to clients and servers.
type CertConfig struct {
	KeyFile   string
	CertFile  string
	RootsFile string

	HTTPCert     bool // have the https stack (grpc-web and hosting) use different pki.Source from the grpc stack
	HTTPKeyFile  string
	HTTPCertFile string
}

func (c CertConfig) FillDefaults() CertConfig {
	or := func(a *string, b string) {
		if *a == "" {
			*a = b
		}
	}

	or(&c.KeyFile, "grpc.key.pem")
	or(&c.CertFile, "grpc.cert.pem")
	or(&c.RootsFile, "grpc.roots.pem")
	// if the config specifies http key or cert file paths, assume they want to use it
	if c.HTTPKeyFile != "" || c.HTTPCertFile != "" {
		c.HTTPCert = true
	}
	or(&c.HTTPKeyFile, "https.key.pem")
	or(&c.HTTPCertFile, "https.cert.pem")
	return c
}

func (sc SystemConfig) LocalConfigPath() string {
	s := sc.LocalConfigFileName
	if s == "" {
		s = LocalConfigFileName
	}
	return filepath.Join(sc.DataDir, s)
}

// Bootstrap will obtain a Controller in a ready-to-run state.
func Bootstrap(ctx context.Context, config SystemConfig) (*Controller, error) {
	logger, err := config.Logger.Build()
	if err != nil {
		return nil, err
	}

	// create data dir if it doesn't exist
	err = os.MkdirAll(config.DataDir, 0750)
	if err != nil {
		return nil, err
	}

	// load the local config file if possible
	// TODO: pull config from manager publication
	var localConfig ControllerConfig
	localConfigPath := config.LocalConfigPath()
	rawLocalConfig, err := os.ReadFile(localConfigPath)
	if err == nil {
		err = json.Unmarshal(rawLocalConfig, &localConfig)
		if err != nil {
			return nil, fmt.Errorf("local config JSON unmarshal: %w", err)
		}
	} else {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("failed to load local config from file", zap.Error(err),
				zap.String("path", localConfigPath))
		} else {
			logger.Debug("failed to load local config from file", zap.Error(err), zap.String("path", localConfigPath))
		}
	}

	// rootNode grants both local (in process) and networked (via grpc.Server) access to controller apis.
	// If you have a device you want expose via a Smart Core API, rootNode is where you'd do that via Announce.
	// If you need to know the brightness of another controller device, rootNode.Clients is how you'd do that.
	rootNode := node.New(localConfig.Name)
	rootNode.Logger = logger.Named("node")

	// Setup a local database for storing non-critical data.
	// This is made available to automations and systems as a local cache, for example for lighting reports.
	dbPath := filepath.Join(config.DataDir, "db.bolt")
	db, err := bolthold.Open(dbPath, 0750, nil)
	if err != nil {
		logger.Warn("failed to open local database file - some system components may fail", zap.Error(err),
			zap.String("path", dbPath))
	}

	certConfig := config.CertConfig.FillDefaults()
	// Create a private key if it doesn't exist.
	// This key is used by the controller for incoming and outgoing connections, and as part of the enrolment process.
	key, keyPEM, err := pki.LoadOrGeneratePrivateKey(filepath.Join(config.DataDir, certConfig.KeyFile), logger)
	if err != nil {
		return nil, err
	}

	// enrollServer manages this controllers participation in a cohort.
	// When registered with a grpc.Server, the enrollServer will accept requests like CreateEnrollment which gives
	// this controller a new certificate for use during outgoing TLS connections to other cohort members.
	// In addition the enrollment will also include details of the trusted root certs so this controller can validate
	// incoming connections that contain a client certificate.
	//
	// enrollServer implements pki.Source providing these features without any extra work to setup.
	enrollServer, err := enrollment.LoadOrCreateServer(filepath.Join(config.DataDir, "enrollment"), keyPEM, logger.Named("enrollment"))
	if err != nil {
		return nil, err
	}

	// fileSource attempts to load a certificate and trust roots from disk.
	// The certificates public key must be paired with private key `key` loaded above.
	fileSource := pki.CacheSource(
		pki.FuncSource(func() (*tls.Certificate, []*x509.Certificate, error) {
			certPath := filepath.Join(config.DataDir, certConfig.CertFile)
			rootsPath := certConfig.RootsFile
			if rootsPath != "" {
				rootsPath = filepath.Join(config.DataDir, rootsPath)
			}
			return pki.LoadCertAndRootsWithKey(certPath, rootsPath, key)
		}),
		expire.BeforeInvalid(time.Hour),
	)

	// selfSignedSource creates a self signed certificate.
	// The certificates public key will be paired with and signed by `key`.
	selfSignedSource := pki.CacheSource(
		pki.SelfSignedSource(key, pki.WithExpireAfter(30*24*time.Hour), pki.WithIfaces()),
		expire.AfterProgress(0.5),
		pki.WithFSCache(filepath.Join(config.DataDir, "self-signed.cert.pem"), "", key),
	)

	// certSource is used by both incoming and outgoing connections.
	// The server present the sources certificate and any intermediates between it and the roots to clients during TLS handshake.
	// If an incoming connection presents a client cert then it will be validated against the roots.
	// Outgoing connections will present the sources certificate as a client cert for validation by the remote party.
	// Config can indicate that different certs be used for grpc and https (inc grpc-web)
	certSource := pki.ChainSource(
		enrollServer,
		fileSource,
		selfSignedSource,
	)
	tlsGRPCServerConfig := pki.TLSServerConfig(certSource)
	tlsGRPCServerConfig.ClientAuth = tls.VerifyClientCertIfGiven
	tlsGRPCClientConfig := pki.TLSClientConfig(certSource)

	// Certs used for https (hosting and grpc-web) can be different from the Smart Core native grpc endpoint,
	// mostly to allow support for trusted certs on the https interface and cohort managed certs for grpc requests.
	httpCertSource := certSource
	if certConfig.HTTPCert {
		fileSource := pki.CacheSource(
			pki.FSSource(
				filepath.Join(config.DataDir, certConfig.HTTPCertFile),
				filepath.Join(config.DataDir, certConfig.HTTPKeyFile),
				""),
			expire.After(30*time.Minute),
		)
		httpCertSource = pki.ChainSource(
			fileSource,
			selfSignedSource, // reuse the same self signed cert from grpc requests
		)
	}
	tlsHTTPServerConfig := pki.TLSServerConfig(httpCertSource)

	// manager represents a delayed connection to the cohort manager.
	// Using the manager connection when we aren't enrolled will result in RPC calls returning 'not resolved' errors or similar.
	// As soon as we get enrolled those connections will be updated with the current manager address and will start to work.
	manager := node.DialChan(ctx, enrollServer.ManagerAddress(ctx),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsGRPCClientConfig)))

	mux := http.NewServeMux()
	if config.StaticDir != "" {
		mux.Handle("/", http.FileServer(http.Dir(config.StaticDir)))
	}

	var grpcOpts []grpc.ServerOption
	grpcOpts = append(grpcOpts, grpc.Creds(credentials.NewTLS(tlsGRPCServerConfig)))

	// tokenValidator validates access tokens as part of the authorisation of requests to our APIs.
	// Claims associated with the token are presented along with other information when processing policy files.
	// Systems contribute validators to this set supporting different sources of token.
	tokenValidator := &token.ValidatorSet{}

	// configure request authorisation, here we setup grpc interceptors that decide if a request is denied or not.
	if !config.DisablePolicy {
		pol := policy.Default(false)
		if config.Policy != nil {
			pol = config.Policy
		}

		interceptor := policy.NewInterceptor(pol,
			policy.WithLogger(logger.Named("policy")),
			policy.WithTokenVerifier(tokenValidator),
		)
		grpcOpts = append(grpcOpts,
			grpc.ChainUnaryInterceptor(interceptor.GRPCUnaryInterceptor()),
			grpc.ChainStreamInterceptor(interceptor.GRPCStreamingInterceptor()),
		)
	}

	if rootNode.Name() != "" {
		grpcOpts = append(grpcOpts,
			grpc.ChainUnaryInterceptor(name.IfAbsentUnaryInterceptor(rootNode.Name())),
			grpc.ChainStreamInterceptor(name.IfAbsentStreamInterceptor(rootNode.Name())),
		)
	}

	grpcServer := grpc.NewServer(grpcOpts...)
	reflection.Register(grpcServer)
	gen.RegisterEnrollmentApiServer(grpcServer, enrollServer)
	devices.NewServer(rootNode).Register(grpcServer)
	// support the services api for managing drivers, automations, and systems
	serviceRouter := gen.NewServicesApiRouter()
	rootNode.Support(node.Routing(serviceRouter), node.Clients(gen.WrapServicesApi(serviceRouter)))

	grpcWebServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))

	httpServer := &http.Server{
		Addr:      config.ListenHTTPS,
		TLSConfig: tlsHTTPServerConfig,
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if grpcWebServer.IsGrpcWebRequest(request) || grpcWebServer.IsAcceptableGrpcCorsRequest(request) {
				grpcWebServer.ServeHTTP(writer, request)
			} else {
				mux.ServeHTTP(writer, request)
			}
		}),
	}

	c := &Controller{
		SystemConfig:     config,
		ControllerConfig: localConfig,
		Enrollment:       enrollServer,
		Logger:           logger,
		Node:             rootNode,
		Tasks:            &task.Group{},
		Database:         db,
		TokenValidators:  tokenValidator,
		Mux:              mux,
		GRPC:             grpcServer,
		HTTP:             httpServer,
		ClientTLSConfig:  tlsGRPCClientConfig,
		ManagerConn:      manager,
	}
	c.Defer(manager.Close)
	return c, nil
}

type Controller struct {
	SystemConfig     SystemConfig
	ControllerConfig ControllerConfig
	Enrollment       *enrollment.Server

	// services for drivers/automations
	Logger          *zap.Logger
	Node            *node.Node
	Tasks           *task.Group
	Database        *bolthold.Store
	TokenValidators *token.ValidatorSet

	Mux  *http.ServeMux
	GRPC *grpc.Server
	HTTP *http.Server

	ClientTLSConfig *tls.Config
	ManagerConn     node.Remote

	deferred []Deferred
}

type Deferred func() error

// Defer indicates that the given Deferred should be executed when the Controllers Run method returns.
func (c *Controller) Defer(d Deferred) {
	c.deferred = append(c.deferred, d)
}

func (c *Controller) Run(ctx context.Context) (err error) {
	defer func() {
		for _, d := range c.deferred {
			err = multierr.Append(err, d())
		}
	}()

	addFactorySupport(c.Node, c.SystemConfig.DriverFactories)
	addFactorySupport(c.Node, c.SystemConfig.AutoFactories)
	addFactorySupport(c.Node, c.SystemConfig.SystemFactories)

	// we delay registering the node servers until now, so that the caller can call c.Node.Support in between
	// Bootstrap and Run and have all these added correctly.
	c.Node.Register(c.GRPC)

	group, ctx := errgroup.WithContext(ctx)
	if c.SystemConfig.ListenGRPC != "" {
		group.Go(func() error {
			return ServeGRPC(ctx, c.GRPC, c.SystemConfig.ListenGRPC, 15*time.Second, c.Logger.Named("server.grpc"))
		})
	}
	if c.SystemConfig.ListenHTTPS != "" {
		group.Go(func() error {
			return ServeHTTPS(ctx, c.HTTP, 15*time.Second, c.Logger.Named("server.https"))
		})
	}

	// load and start the systems
	systemServices, err := c.startSystems()
	if err != nil {
		return err
	}
	c.Node.Announce("systems", node.HasClient(gen.WrapServicesApi(services.NewApi(systemServices, services.WithKnownTypesFromMapKeys(c.SystemConfig.SystemFactories)))))
	// load and start the drivers
	driverServices, err := c.startDrivers()
	if err != nil {
		return err
	}
	c.Node.Announce("drivers", node.HasClient(gen.WrapServicesApi(services.NewApi(driverServices, services.WithKnownTypesFromMapKeys(c.SystemConfig.DriverFactories)))))
	// load and start the automations
	autoServices, err := c.startAutomations()
	if err != nil {
		return err
	}
	c.Node.Announce("automations", node.HasClient(gen.WrapServicesApi(services.NewApi(autoServices, services.WithKnownTypesFromMapKeys(c.SystemConfig.AutoFactories)))))

	err = multierr.Append(err, group.Wait())
	return
}

// addFactorySupport is used to register factories with a node to expose custom factory APIs.
// This checks each value in m and if that value has an API, via node.SelfSupporter, then it is registered with s.
func addFactorySupport[M ~map[K]F, K comparable, F any](s node.Supporter, m M) {
	for _, factory := range m {
		if api, ok := any(factory).(node.SelfSupporter); ok {
			api.AddSupport(s)
		}
	}
}

func (c *Controller) startDrivers() (*service.Map, error) {
	ctxServices := driver.Services{
		Logger:          c.Logger.Named("driver"),
		Node:            c.Node,
		ClientTLSConfig: c.ClientTLSConfig,
		HTTPMux:         c.Mux,
	}

	m := service.NewMap(func(kind string) (service.Lifecycle, error) {
		f, ok := c.SystemConfig.DriverFactories[kind]
		if !ok {
			return nil, fmt.Errorf("unsupported driver type %v", kind)
		}
		return f.New(ctxServices), nil
	}, service.IdIsRequired)

	var allErrs error
	for _, cfg := range c.ControllerConfig.Drivers {
		_, _, err := m.Create(cfg.Name, cfg.Type, service.State{Active: !cfg.Disabled, Config: cfg.Raw})
		allErrs = multierr.Append(allErrs, err)
	}
	return m, allErrs
}

func (c *Controller) startAutomations() (*service.Map, error) {
	ctxServices := auto.Services{
		Logger:       c.Logger.Named("auto"),
		Node:         c.Node,
		Database:     c.Database,
		GRPCServices: c.GRPC,
	}

	m := service.NewMap(func(kind string) (service.Lifecycle, error) {
		f, ok := c.SystemConfig.AutoFactories[kind]
		if !ok {
			return nil, fmt.Errorf("unsupported automation type %v", kind)
		}
		return f.New(ctxServices), nil
	}, service.IdIsRequired)

	var allErrs error
	for _, cfg := range c.ControllerConfig.Automation {
		_, _, err := m.Create(cfg.Name, cfg.Type, service.State{Active: !cfg.Disabled, Config: cfg.Raw})
		allErrs = multierr.Append(allErrs, err)
	}
	return m, allErrs
}

func (c *Controller) startSystems() (*service.Map, error) {
	ctxServices := system.Services{
		DataDir:         c.SystemConfig.DataDir,
		Logger:          c.Logger.Named("system"),
		Node:            c.Node,
		Database:        c.Database,
		HTTPMux:         c.Mux,
		TokenValidators: c.TokenValidators,
	}
	m := service.NewMap(func(kind string) (service.Lifecycle, error) {
		f, ok := c.SystemConfig.SystemFactories[kind]
		if !ok {
			return nil, fmt.Errorf("unsupported system type %v", kind)
		}
		return f.New(ctxServices), nil
	}, service.IdIsKind)

	var allErrs error
	for kind, cfg := range c.ControllerConfig.Systems {
		_, _, err := m.Create("", kind, service.State{Active: !cfg.Disabled, Config: cfg.Raw})
		allErrs = multierr.Append(allErrs, err)
	}
	return m, allErrs
}
