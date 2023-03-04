package config

import (
	"database/sql"
)

type Dependencies struct {
	Cfg *Config
	Db  *sql.DB
}

type MainConfig struct {
	Environment string
	Config      []Config
}

type Config struct {
	EnvType  string
	Port     string
	PemLoc   string
	KeyLoc   string
	Database Database
}

type Database struct {
	Server   string
	Port     string
	Service  string
	Username string
	Password string
}
