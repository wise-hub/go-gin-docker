package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() (*Dependencies, error) {
	cfg, err := loadCfg()
	if err != nil {
		return nil, err
	}

	var curCfg = cfg.Config[0] // test default

	if cfg.Environment == "PROD" {
		curCfg = cfg.Config[1]
		gin.SetMode(gin.ReleaseMode)
	}
	db, err := connectDb(&curCfg.Database)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return &Dependencies{
		Cfg: &curCfg,
		Db:  db,
	}, nil
}

func loadCfg() (*MainConfig, error) {
	file, err := os.Open("./config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg MainConfig

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
