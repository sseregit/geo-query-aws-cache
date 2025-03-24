package service

import (
	"errors"
	"fmt"
	"geo-query-aws-cache/aws"
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db"
	. "geo-query-aws-cache/db/mysql/types"
	"geo-query-aws-cache/module/API/types"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type service struct {
	cfg *config.Config
	db  *db.DBRoot
	aws *aws.Aws
}

type ServiceImpl interface {
	RegisterUser(req types.RegisterUserReq) error
	UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error
	FindAroundUsers(userName string, searchRange, limit int64) ([]*User, error)
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
	if err := s.db.MySQL.RegisterUser(req.UserName, req.Description, req.Hobby, req.Latitude, req.Hardness); err != nil {
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
}

func (s *service) FindAroundUsers(userName string, searchRange, limit int64) ([]*User, error) {
	if limit == 0 {
		limit = 5
	}

	if u, err := s.getUser(userName); err != nil {
		return nil, err
	} else if users, err := s.db.MySQL.AroundUser(u.UserName, u.Latitude, u.Hardness, searchRange, limit); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (s *service) getUser(userName string) (*User, error) {
	if u, err := s.db.MySQL.GetUser(userName); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}

func (s *service) UploadFile(userName string, header *multipart.FileHeader, file multipart.File) error {
	fileName := header.Filename
	fileExt := filepath.Ext(fileName)

	if !solveImageExtension(fileExt) {
		return errors.New("Failed To Solve Extension")
	} else {
		path := "./temp"
		filePath := fmt.Sprintf("%s/%s", path, fileName)

		if out, err := os.Create(filePath); err != nil {
			return err
		} else {
			defer out.Close()

			if _, err := io.Copy(out, file); err != nil {
				return err
			} else {
				if err = s.putFileToS3(fileName, userName, strings.TrimPrefix(fileExt, "."), filePath); err != nil {
					return err
				} else {
					return nil
				}
			}
		}
	}
	return nil
}

func (s *service) putFileToS3(fileName, userName, extension, path string) error {
	key := userName + "/" + fileName

	if f, err := os.Open(path); err != nil {
		return err
	} else {
		defer f.Close()

		if err = s.aws.PutFileToS3(key, extension, f); err != nil {
			return err
		} else {
			return nil
		}
	}
}
