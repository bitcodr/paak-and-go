package config

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DB
	Service
}

func Load(ctx context.Context) (*Config, error) {
	v := viper.New()

	setDefaultServiceConfig(v)

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("/app/config.d")
	v.AddConfigPath("config.d")
	v.AddConfigPath(".\\config.d")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error in reading configs from file: %+v \n\n", err)
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config Config

	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	watchConfigChanges(ctx, v)

	return &config, nil
}

func watchConfigChanges(_ context.Context, v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		var config Config
		err := v.Unmarshal(&config)
		if err != nil {
			log.Fatalln("Fatal error when unmarshal config:", err)
		}
	})
}
