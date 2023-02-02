package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hi20160616/voter/configs"
)

type Client struct {
	db             *sql.DB
	DatabaseClient *DatabaseClient
}

type DatabaseClient struct {
	db *sql.DB
}

func open(cfg *configs.Config) (*sql.DB, error) {
	return sql.Open(cfg.Database.Driver, cfg.Database.Source)
}

func NewClient() (*Client, error) {
	cfg := configs.NewConfig("voter")
	if cfg.Err != nil {
		return &Client{nil, nil}, cfg.Err
	}
	db, err := open(cfg)
	return &Client{db, &DatabaseClient{db}}, err
}
