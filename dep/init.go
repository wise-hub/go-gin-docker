package dep

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

func Init() (*Dependencies, error) {
	cfg, err := loadCfg()
	if err != nil {
		return nil, err
	}

	db, err := connectDb(cfg.Database)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Cfg: cfg,
		Db:  db,
	}, nil
}

func loadCfg() (*Config, error) {
	file, err := os.Open("./config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func connectDb(cfg *Database) (*sql.DB, error) {
	con := fmt.Sprintf("oracle://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Server,
		cfg.Port,
		cfg.Service)
	db, err := sql.Open("oracle", con)

	if err != nil {
		return nil, err
	}

	return db, nil
}
