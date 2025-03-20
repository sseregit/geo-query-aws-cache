package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
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

func unMarshalToField(field []interface{}, to ...interface{}) error {
	if len(field) != len(to) {
		return errors.New("Field Length is not match")
	} else {
		for i, f := range field {
			if err := json.Unmarshal(f.([]byte), to[i]); err != nil {
				return err
			}
		}
		return nil
	}
}
