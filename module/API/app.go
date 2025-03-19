package API

import (
	"geo-query-aws-cache/aws"
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db"
	"geo-query-aws-cache/module/API/network"
	"geo-query-aws-cache/module/API/service"
)

type API struct {
	cfg *config.Config
	db  *db.DBRoot
	aws *aws.Aws
}

func NewAPI(cfg *config.Config, db *db.DBRoot, aws *aws.Aws) *API {
	api := &API{cfg, db, aws}

	s := service.NewService(cfg, db, aws)
	n := network.NewNetwork(cfg, s)

	go func() {
		n.Start()
	}()

	return api
}
