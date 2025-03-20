package network

import (
	"geo-query-aws-cache/config"
	"geo-query-aws-cache/module/API/service"
	"github.com/gin-gonic/gin"
)

type Network struct {
	s    service.ServiceImpl
	e    *gin.Engine
	port string
	cfg  *config.Config
}

func NewNetwork(
	cfg *config.Config,
	s service.ServiceImpl,
) *Network {
	n := &Network{cfg: cfg, s: s, e: gin.New(), port: cfg.Info.Port}
	setGin(n.e)
	userRouter(n)
	return n
}

func (n *Network) Start() error {
	return n.e.Run(n.port)
}
