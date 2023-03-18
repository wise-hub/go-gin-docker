package config

import (
	"database/sql"
	"time"
)

type Dependencies struct {
	Cfg *Config
	Db  *sql.DB
}

type MainConfig struct {
	Environment string `validate:"required"`
	Config      []Config
}

type Config struct {
	EnvType             string   `json:"env_type" validate:"required"`
	Port                string   `json:"port" validate:"required"`
	SessionLifetimeMins string   `json:"session_lifetime_mins" validate:"required"`
	TokenDbCheck        string   `json:"token_db_check" validate:"required"`
	Database            Database `json:"database" validate:"required"`
	LDAP                LDAP     `json:"ldap" validate:"required"`
}

type Database struct {
	Server          string        `json:"server" validate:"required"`
	Port            string        `json:"port" validate:"required"`
	Service         string        `json:"service" validate:"required"`
	Username        string        `json:"username" validate:"required"`
	Password        string        `json:"password" validate:"required"`
	MaxOpenConns    int           `json:"max_open_conns" validate:"required"`
	MaxIdleConns    int           `json:"max_idle_conns" validate:"required"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" validate:"required"`
}

type LDAP struct {
	Server string `json:"server" validate:"required"`
	Port   string `json:"port" validate:"required"`
	UserDN string `json:"user_dn" validate:"required"`
}
