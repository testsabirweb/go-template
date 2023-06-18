package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() (*Config, error) {
	config := &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "192.168.1.7", //Your machine ip address
			Port:     3306,
			Username: "guest",
			Password: "Guest0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}

	return config, nil
}
