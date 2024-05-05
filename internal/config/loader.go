package config

func Get() *Config {
	return &Config{
		SRV: Server{
			Host: "localhost",
			Port: "8080",
		},
		DB: Database{
			Host:     "localhost",
			Port:     3630,
			User:     "mysql",
			Password: "mysql",
			Name:     "campsite",
		},
	}
}
