package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DB struct {
	Connections map[string]*Connection
}

type Connection struct {
	Name     string
	Port     int
	Host     string
	Username string
	Password string
	Ssl      string
}

type Service struct {
	Host         string
	RestPort     string        `mapstructure:"rest_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

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

func setDefaultServiceConfig(v *viper.Viper) {
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.rest_port", "8081")
	v.SetDefault("server.read_timeout", 30*time.Second)
	v.SetDefault("server.write_timeout", 30*time.Second)
	v.SetDefault("server.idle_timeout", 30*time.Second)
}
