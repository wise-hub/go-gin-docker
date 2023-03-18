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
	EnvType         string
	SessionLifetime string
	Port            string
	Database        Database
	LDAP            LDAP
}

type Database struct {
	Server   string
	Port     string
	Service  string
	Username string
	Password string
}

type LDAP struct {
	Server string
	Port   string
	UserDN string
}
