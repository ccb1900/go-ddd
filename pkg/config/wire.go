package config

func ProvideDbConfig() DbConfig {
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