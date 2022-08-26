package tenantapi

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"regexp"

	"github.com/jackc/pgx/v4"
	"github.com/smart-core-os/sc-golang/pkg/masks"
	"github.com/vanti-dev/bsp-ew/internal/db"
	"github.com/vanti-dev/bsp-ew/internal/util/rpcutil"
	"github.com/vanti-dev/bsp-ew/pkg/gen"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

var (
	errDatabase       = status.Error(codes.Internal, "database transaction failed")
	errTenantNotFound = status.Error(codes.NotFound, "tenant not found")
	errSecretNotFound = status.Error(codes.NotFound, "secret not found")
)

type Server struct {
	gen.UnimplementedTenantApiServer
	dbConn *pgx.Conn
	logger *zap.Logger
}

type Option func(server *Server)

func WithLogger(logger *zap.Logger) Option {
	return func(server *Server) {
		server.logger = logger
	}
}

func NewServer(conn *pgx.Conn, options ...Option) *Server {
	s := &Server{
		dbConn: conn,
		logger: zap.NewNop(),
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

func (s *Server) ListTenants(ctx context.Context, request *gen.ListTenantsRequest) (*gen.ListTenantsResponse, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)

	var tenants []*gen.Tenant
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		tenants, err = db.ListTenants(ctx, tx)
		return
	})
	if err != nil {
		logger.Error("db.ListTenants failed", zap.Error(err))
		return nil, errDatabase
	}

	return &gen.ListTenantsResponse{Tenants: tenants}, nil
}

func (s *Server) PullTenants(request *gen.PullTenantsRequest, server gen.TenantApi_PullTenantsServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Server) CreateTenant(ctx context.Context, request *gen.CreateTenantRequest) (*gen.Tenant, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)

	var newTenant *gen.Tenant
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		newTenant, err = db.CreateTenant(ctx, tx, request.Tenant.Title)
		if err != nil {
			return
		}

		if zones := request.Tenant.GetZoneNames(); len(zones) > 0 {
			err = db.AddTenantZones(ctx, tx, newTenant.Id, zones)
		}

		return
	})
	if err != nil {
		logger.Error("tenant database transaction failed", zap.Error(err))
		return nil, errDatabase
	}

	return newTenant, nil
}

func (s *Server) GetTenant(ctx context.Context, request *gen.GetTenantRequest) (*gen.Tenant, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)

	var tenant *gen.Tenant
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) error {
		var err error
		tenant, err = db.GetTenant(ctx, tx, request.GetId())
		return err
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errTenantNotFound
	} else if err != nil {
		logger.Error("db transaction failed", zap.Error(err))
		return nil, errDatabase
	}
	return tenant, nil
}

func (s *Server) UpdateTenant(ctx context.Context, request *gen.UpdateTenantRequest) (*gen.Tenant, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger).With(zap.String("tenant", request.GetTenant().GetId()))
	tenant := request.Tenant
	updater := masks.NewFieldUpdater(
		masks.WithUpdateMask(request.UpdateMask),
		masks.WithUpdateMaskFieldName("update_mask"),
		masks.WithWritableFields(&fieldmaskpb.FieldMask{Paths: []string{
			"title", "zone_names",
		}}),
	)
	err := updater.Validate(tenant)
	if err != nil {
		logger.Error("mask validation failed", zap.Error(err), zap.Strings("paths", request.UpdateMask.GetPaths()))
		return nil, err
	}

	err = s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		if rpcutil.MaskContains(request.UpdateMask, "title") {
			err = db.UpdateTenantTitle(ctx, tx, tenant.Id, tenant.Title)
			if err != nil {
				return err
			}
		}

		if rpcutil.MaskContains(request.UpdateMask, "zone_names") {
			err = db.ReplaceTenantZones(ctx, tx, tenant.Id, tenant.ZoneNames)
			if err != nil {
				return err
			}
		}

		tenant, err = db.GetTenant(ctx, tx, tenant.Id)
		return
	})

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errTenantNotFound
	} else if err != nil {
		logger.Error("database transaction failed", zap.Error(err))
		return nil, errDatabase
	}

	return tenant, nil
}

func (s *Server) DeleteTenant(ctx context.Context, request *gen.DeleteTenantRequest) (*gen.DeleteTenantResponse, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger).With(zap.String("id", request.Id))

	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) error {
		return db.DeleteTenant(ctx, tx, request.Id)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errTenantNotFound
	} else if err != nil {
		logger.Error("db.DeleteTenant failed", zap.Error(err))
		return nil, errDatabase
	}

	return &gen.DeleteTenantResponse{}, nil
}

func (s *Server) PullTenant(request *gen.PullTenantRequest, server gen.TenantApi_PullTenantServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Server) AddTenantZones(ctx context.Context, request *gen.AddTenantZonesRequest) (*gen.Tenant, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)

	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) error {
		return db.AddTenantZones(ctx, tx, request.TenantId, request.AddZoneNames)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errTenantNotFound
	} else if err != nil {
		logger.Error("db transaction failed", zap.Error(err))
		return nil, errDatabase
	}

	return s.GetTenant(ctx, &gen.GetTenantRequest{Id: request.TenantId})
}

func (s *Server) RemoveTenantZones(ctx context.Context, request *gen.RemoveTenantZonesRequest) (*gen.Tenant, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)

	var tenant *gen.Tenant
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		err = db.RemoveTenantZones(ctx, tx, request.TenantId, request.RemoveZoneNames)
		if err != nil {
			return err
		}

		tenant, err = db.GetTenant(ctx, tx, request.TenantId)
		return
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errTenantNotFound
	} else if err != nil {
		logger.Error("db transaction failed", zap.Error(err))
		return nil, errDatabase
	}

	return tenant, nil
}

func (s *Server) ListSecrets(ctx context.Context, request *gen.ListSecretsRequest) (*gen.ListSecretsResponse, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger).With(
		zap.Namespace("request"),
		zap.String("filter", request.GetFilter()),
		zap.Bool("include_hash", request.GetIncludeHash()),
	)

	var tenantID string
	if request.Filter != "" {
		groups := filterTenantRegexp.FindStringSubmatch(request.Filter)
		if groups == nil {
			return nil, status.Error(codes.InvalidArgument, "invalid filter")
		}
		tenantID = groups[1]
	}

	var secrets []*gen.Secret
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		secrets, err = db.ListTenantSecrets(ctx, tx, tenantID)
		return
	})
	if err != nil {
		logger.Error("db.ListTenantSecrets failed", zap.Error(err))
		return nil, errDatabase
	}
	// unless specifically requested, censor the hashes
	if !request.IncludeHash {
		for i := range secrets {
			secrets[i].SecretHash = nil
		}
	}

	return &gen.ListSecretsResponse{Secrets: secrets}, nil
}

func (s *Server) PullSecrets(request *gen.PullSecretsRequest, server gen.TenantApi_PullSecretsServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Server) CreateSecret(ctx context.Context, request *gen.CreateSecretRequest) (*gen.Secret, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)
	secret := request.Secret

	var err error
	secret.Secret, err = genSecret()
	if err != nil {
		logger.Error("secret generation failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "secret generation failed")
	}
	secret.SecretHash = hashSecret(secret.Secret)

	err = s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		secret, err = db.CreateTenantSecret(ctx, tx, secret)
		return
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "tenant not found")
	} else if err != nil {
		logger.Error("db transaction failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "database transaction failed")
	}
	return secret, nil
}

func (s *Server) GetSecret(ctx context.Context, request *gen.GetSecretRequest) (*gen.Secret, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger).With(zap.String("secret_id", request.GetId()))

	var secret *gen.Secret
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) (err error) {
		secret, err = db.GetTenantSecret(ctx, tx, request.Id)
		return
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "secret not found")
	} else if err != nil {
		logger.Error("db.GetTenantSecret failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "database transaction failed")
	}

	return secret, nil
}

func (s *Server) UpdateSecret(ctx context.Context, request *gen.UpdateSecretRequest) (*gen.Secret, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Server) DeleteSecret(ctx context.Context, request *gen.DeleteSecretRequest) (*gen.DeleteSecretResponse, error) {
	logger := rpcutil.ServerLogger(ctx, s.logger)
	err := s.dbConn.BeginFunc(ctx, func(tx pgx.Tx) error {
		return db.DeleteTenantSecret(ctx, tx, request.Id)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "secret not found")
	} else if err != nil {
		logger.Error("db transaction failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "database transaction failed")
	}
	return &gen.DeleteSecretResponse{}, nil
}

func (s *Server) PullSecret(request *gen.PullSecretRequest, server gen.TenantApi_PullSecretServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func (s *Server) RegenerateSecret(ctx context.Context, request *gen.RegenerateSecretRequest) (*gen.Secret, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func genSecret() (string, error) {
	secretBytes := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, secretBytes)
	if err != nil {
		return "", err
	}

	encoding := base64.URLEncoding
	encoded := make([]byte, encoding.EncodedLen(len(secretBytes)))
	encoding.Encode(encoded, secretBytes)

	return string(encoded), nil
}

func hashSecret(secret string) (hash []byte) {
	sum := sha256.Sum256([]byte(secret))
	return sum[:]
}

var filterTenantRegexp = regexp.MustCompile(
	`^\s*tenant\.id\s*=\s*"?([0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12})"?\s*$`,
)
