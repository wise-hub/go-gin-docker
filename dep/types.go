package dep

import "database/sql"

type Dependencies struct {
	Cfg *Config
	Db  *sql.DB
}

type Config struct {
	Environment string
	Port        string
	PemLoc      string
	KeyLoc      string
	Database    *Database
}

type Database struct {
	DoConnect bool
	Server    string
	Port      string
	Service   string
	Username  string
	Password  string
}
