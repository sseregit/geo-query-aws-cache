package aws

import (
	"geo-query-aws-cache/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Aws struct {
	session *session.Session
	S3      *s3.S3
	Bucket  string
	cfg     *config.Config
}

func NewAws(cfg *config.Config) *Aws {
	a := &Aws{cfg: cfg}
	var err error

	if a.session, err = session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Aws.Region),
		Credentials: credentials.NewStaticCredentials(cfg.Aws.Key, cfg.Aws.SecretKey, ""),
	}); err != nil {
		panic(err)
	} else {
		a.Bucket = cfg.Aws.Bucket
		a.S3 = s3.New(a.session)
	}

	return a
}
