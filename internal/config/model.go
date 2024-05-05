package config

type Config struct {
	SRV Server
	DB  Database
	JWT SecretKey
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type SecretKey struct {
	Key string
}
