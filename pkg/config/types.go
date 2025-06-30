package config

type AppConfig struct {
	Debug  bool
	Server ServerConfig
	DB     DatabseConfig `mapstructure:"database" json:"database"`
	Log    LogConfig
}

type ServerConfig struct {
	Port string
}

type LogConfig struct {
	Filename   string
	MaxSize    int // MB
	MaxAge     int // days
	MaxBackups int
	Compress   bool
	Level      string // "debug", "info", etc.
}

type DatabseConfig struct {
	Default      string              `mapstructure:"default" json:"default"`
	BatchSize    int                 `mapstructure:"batchsize" json:"batchsize"`
	Databases    map[string]DBConfig `mapstructure:"databases" json:"databases"`
	DbPoolConfig DbPoolConfig        `mapstructure:"pool" json:"pool"`
}

type DbPoolConfig struct {
	MaxLifeTime int                 `mapstructure:"maxlifetime" json:"maxlifetime"`
	MaxIdleTime int                 `mapstructure:"maxidletime" json:"maxidletime"`
	MaxIdle     int                 `mapstructure:"maxidle" json:"maxidle"`
	MaxOpen     int                 `mapstructure:"maxopen" json:"maxopen"`
	Monitor     DbPoolMonitorConfig `mapstructure:"monitor" json:"monitor"`
}

type DbPoolMonitorConfig struct {
	Enabled bool `mapstructure:"enabled"`
	Period  int  `mapstructure:"period"`
}

type DBConfig struct {
	Name     string `mapstructure:"name"`
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}
