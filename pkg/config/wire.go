package config

func ProvideDbConfig() DatabseConfig {
	return Cfg.DB
}
func ProvideLogConfig() LogConfig {
	return Cfg.Log
}
func ProvideServerConfig() ServerConfig {
	return Cfg.Server
}

func ProvideConfig() AppConfig {
	return Cfg
}
