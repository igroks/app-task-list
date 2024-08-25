package config

import (
	"os"
)

var Env Evironment

type Evironment struct {
	Database string
}

func init() {
	Env = Evironment {
		Database: os.Getenv("DB_URL"),
	}
}
