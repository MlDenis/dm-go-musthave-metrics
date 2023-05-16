package server

import (
	"fmt"
	"github.com/MlDenis/dm-go-musthave-metrics/internal/storage"
	"github.com/gin-gonic/gin"
)

const (
	TypeS  = "type"
	NameS  = "name"
	ValueS = "value"
)

type ServerConfig struct {
	addr string
	prt  string
}

func NewServerConfig(address string, port string) ServerConfig {
	return ServerConfig{addr: address, prt: port}
}

type MSServer struct {
	MS     storage.MemStorage
	Config ServerConfig
	Router *gin.Engine
}

func MakeNewMSServer(adress string, port string) MSServer {
	sc := NewServerConfig(adress, port)
	ms := storage.NewMetricsStorage()
	s := MSServer{MS: ms, Config: sc}

	r := gin.Default()
	r.RedirectTrailingSlash = false

	r.POST("/update/:type/:name/:value", s.PostSingleValue)
	r.GET("/value/:type/:name", s.GetSingleValue)
	r.GET("/", s.GetMSDataHowHTML)

	s.Router = r

	return s
}

func (s *MSServer) ServerStart() error {
	return s.Router.Run(fmt.Sprintf("%s:%s", s.Config.addr, s.Config.prt))
}
