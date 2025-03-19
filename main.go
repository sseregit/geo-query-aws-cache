package main

import (
	"flag"
	"fmt"
	"geo-query-aws-cache/config"
)

var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()

	config.NewConfig(*cfgPath)
	fmt.Println("start")
}
