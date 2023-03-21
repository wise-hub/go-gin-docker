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
	Environment string   `validate:"required,oneof=DEV TEST PROD"`
	Config      []Config `validate:"required,dive"`
}

type Config struct {
	EnvType             string   `json:"env_type" validate:"required"`
	Port                string   `json:"port" validate:"required,min=2,max=5,numeric"`
	SessionLifetimeMins string   `json:"session_lifetime_mins" validate:"required,max=4,numeric"`
	TokenDbCheck        string   `json:"token_db_check" validate:"required,oneof=Y N"`
	Database            Database `json:"database" validate:"required"`
	LDAP                LDAP     `json:"ldap" validate:"required"`
}

type Database struct {
	Server          string        `json:"server" validate:"required,hostname|ip4_addr"`
	Port            string        `json:"port" validate:"required,min=2,max=5,numeric"`
	Service         string        `json:"service" validate:"required,max=20,alphanum"`
	Username        string        `json:"username" validate:"required,max=30,alphanum"`
	Password        string        `json:"password" validate:"required,max=50"`
	MaxOpenConns    int           `json:"max_open_conns" validate:"required,gte=1,lte=500"`
	MaxIdleConns    int           `json:"max_idle_conns" validate:"required,gte=1,lte=450"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" validate:"required"`
}

type LDAP struct {
	Server string `json:"server" validate:"required,hostname|ip4_addr"`
	Port   string `json:"port" validate:"required,min=2,max=5,numeric"`
	UserDN string `json:"user_dn" validate:"required,max=30"`
}
