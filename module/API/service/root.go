package service

import (
	"geo-query-aws-cache/aws"
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db"
	"geo-query-aws-cache/module/API/types"
	"log"
	"mime/multipart"
)

type service struct {
	cfg *config.Config
	db  *db.DBRoot
	aws *aws.Aws
}

type ServiceImpl interface {
	RegisterUser(req types.RegisterUserReq) error
	UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error
	FindAroundUsers(userName string, searchRange, limit int64) (interface{}, error)
}

func NewService(
	cfg *config.Config,
	db *db.DBRoot,
	aws *aws.Aws,
) ServiceImpl {
	s := &service{cfg, db, aws}
	return s
}

func (s *service) RegisterUser(req types.RegisterUserReq) error {
	var retryCount = 0
createAgain:
	if err := s.db.MySQL.RegisterUser(req.UserName, req.Descsription, req.Hobby, req.Latitude, req.Hardness); err != nil {
		retryCount++
		if retryCount < 3 {
			goto createAgain
		} else {
			log.Println("Failed To Create User", "user", req.UserName, "err", err.Error())
			return err
		}
	} else {
		log.Println("Success To Create New User", "name", req.UserName)
		return nil
	}

	return nil
}

func (s *service) UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error {
	return nil
}

func (s *service) FindAroundUsers(userName string, searchRange, limit int64) (interface{}, error) {
	return nil, nil
}
