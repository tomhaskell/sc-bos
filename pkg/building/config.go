package building

type SystemConfig struct {
	ListenGRPC             string `json:"listen-grpc"`
	ListenHTTPS            string `json:"listen-https"`
	SelfSignedHTTPS        bool   `json:"self-signed-https"`
	StaticDir              string `json:"static-dir"`
	DisableAuth            bool   `json:"unsafe-disable-auth"`
	DatabaseURL            string `json:"database-url"`
	DatabasePasswordFile   string `json:"database-password-file"`
	KeycloakAddress        string `json:"keycloak-address"`
	KeycloakRealm          string `json:"keycloak-realm"`
	EnrollmentValidityDays int    `json:"enrollment-validity-days"`
	CanonicalAddress       string `json:"canonical-address"`
}

func DefaultSystemConfig() SystemConfig {
	return SystemConfig{
		ListenGRPC:             ":23557",
		ListenHTTPS:            ":443",
		DatabaseURL:            "postgres://postgres:postgres@localhost/smart_core",
		KeycloakAddress:        "http://localhost:8888/",
		KeycloakRealm:          "smart-core",
		EnrollmentValidityDays: 60,
	}
}