package config

const (
	POSTGRES = "postgres"
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
