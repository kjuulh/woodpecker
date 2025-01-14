package types

type Type string

const (
	TypeForge              Type = "forge"
	TypeConfigService      Type = "config_service"
	TypeSecretService      Type = "secret_service"
	TypeEnvironmentService Type = "environment_service"
	TypeRegistryService    Type = "registry_service"
)
