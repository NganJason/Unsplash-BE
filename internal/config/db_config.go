package config

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/NganJason/BE-template/pkg/cerr"
)

type Database struct {
	Username       string `json:"db_username"`
	Password       string `json:"db_password"`
	Host           string `json:"host"`
	Port           string `json:"db_port"`
	DBName         string `json:"db_name"`
	PoolMaxOpen    int    `json:"pool_max_open"`
	PoolMaxIdle    int    `json:"pool_max_idle"`
	MaxIdleSeconds int    `json:"max_idle_seconds"`
	MaxLifeSeconds int    `json:"max_life_seconds"`
}

func initDB(cfg *Database) (*sql.DB, error) {
	pool, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?parseTime=true",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.DBName))
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("init DB err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if err = pool.Ping(); err != nil {
		return nil, cerr.New(
			fmt.Sprintf("ping DB err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	pool.SetMaxIdleConns(cfg.PoolMaxIdle)
	pool.SetMaxOpenConns(cfg.PoolMaxOpen)
	pool.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleSeconds) * time.Second)
	pool.SetConnMaxLifetime(time.Duration(cfg.MaxLifeSeconds) * time.Second)

	return pool, nil
}
