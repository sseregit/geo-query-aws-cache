package main

import (
	"flag"
	"geo-query-aws-cache/aws"
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/db"
	"geo-query-aws-cache/module/API"
	"go.uber.org/fx"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*cfgPath)

	fx.New(
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(func() *db.DBRoot { return db.RootDB(cfg) }),
		fx.Provide(func() *aws.Aws { return aws.NewAws(cfg) }),
		fx.Provide(API.NewAPI),

		fx.Invoke(func(root *API.API) {}),
	).Run()
}
