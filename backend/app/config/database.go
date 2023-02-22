package config

import (
	"os"
)

var Env Evironment

type Evironment struct {
	Database map[string]string
}

func init() {
	Env = Evironment{
		Database: map[string]string{
			"db1": os.Getenv("DB1_URL"),
			"db2": os.Getenv("DB2_URL"),
		},
	}
}
