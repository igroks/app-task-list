package config

import (
	"net"
	"net/url"
	"os"
)

var cfg config

type config struct {
	Database map[string]DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func ParseUrl(databaseUrl string) DatabaseConfig {
	u, err := url.Parse(databaseUrl)
	if err != nil {
		panic(err)
	}

	host, port, _ := net.SplitHostPort(u.Host)
	user := u.User.Username()
	password, _ := u.User.Password()
	database := u.Path[1:]

	return DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
}

func init() {
	cfg = config{
		Database: map[string]DatabaseConfig{
			"db1": ParseUrl(os.Getenv("DB1_URL")),
			"db2": ParseUrl(os.Getenv("DB2_URL")),
		},
	}
}

func GetDatabaseConfig(key string) DatabaseConfig {
	return cfg.Database[key]
}
