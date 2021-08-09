package config

import "github.com/spf13/viper"

type DB struct {
	Default     string
	Connections []*Connection
}

type Connection struct {
	Name     string
	Port     int
	Host     string
	Username string
	Password string
	Ssl      string
}

func setDefaultDBConfig(v *viper.Viper) {
	v.SetDefault("db.Default", "postgres")
}
