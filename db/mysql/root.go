package mysql

import (
	"database/sql"
	"geo-query-aws-cache/config"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	cfg *config.Config
	db  *sql.DB
}

func NewDB(cfg *config.Config) *DB {
	d := &DB{cfg: cfg}

	var err error

	if d.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		panic(err)
	} else if err = d.db.Ping(); err != nil {
		panic(err)
	} else {
		return d
	}

	return d
}
