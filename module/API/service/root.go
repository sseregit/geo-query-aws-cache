package service

import (
	"geo-query-aws-cache/aws"
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db"
)

type service struct {
	cfg *config.Config
	db  *db.DBRoot
	aws *aws.Aws
}

type ServiceImpl interface {
}

func NewService(
	cfg *config.Config,
	db *db.DBRoot,
	aws *aws.Aws,
) ServiceImpl {
	s := service{cfg, db, aws}
	return s
}
