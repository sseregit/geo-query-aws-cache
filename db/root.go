package db

import (
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db/mysql"
)

type DBRoot struct {
	cfg   *config.Config
	MySQL *mysql.DB
}

func RootDB(cfg *config.Config) *DBRoot {
	root := &DBRoot{cfg: cfg}

	root.MySQL = mysql.NewDB(cfg)

	return root
}
