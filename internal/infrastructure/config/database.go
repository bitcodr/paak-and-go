package config

type DB struct {
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