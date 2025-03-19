package API

import (
	"geo-query-aws-cache/aws"
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db"
)

type API struct {
	cfg *config.Config
	db  *db.DBRoot
	aws *aws.Aws
}

func NewAPI(cfg *config.Config, db *db.DBRoot, aws *aws.Aws) *API {
	api := &API{cfg, db, aws}
	return api
}
