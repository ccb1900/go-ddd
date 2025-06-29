package config

type AppConfig struct {
    Server ServerConfig
    DB     DbConfig
    Log    LogConfig
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	DSN string
}

type LogConfig struct {
	Filename   string
    MaxSize    int  // MB
    MaxAge     int  // days
    MaxBackups int
    Compress   bool
    Level      string // "debug", "info", etc.
}