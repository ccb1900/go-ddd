package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var Cfg AppConfig
var cfgOnce sync.Once

func Init(name string) {
	cfgOnce.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName(name + ".dev")
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("read config error: %v", err)
		}

		if err := viper.Unmarshal(&Cfg); err != nil {
			log.Fatalf("unmarshal config error: %v", err)
		}

		log.Println("config loaded:", name)
	})

}
