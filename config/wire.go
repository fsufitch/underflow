package config

import "github.com/google/wire"

// EnvironmentProviderSet is a Wire provider set for environment configuration
var EnvironmentProviderSet = wire.NewSet(
	ProvideDebugModeFromEnvironment,
	ProvideUnderflowListenPortFromEnvironment,
	ProvideUnderflowMasterAddrFromEnvironment,
	ProvideUnderflowMinionAddrsFromEnvironment,
	ProvideUnderflowModeFromEnvironment,
)
