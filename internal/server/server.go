package server

import (
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
}

func NewServerConfig(address string) ServerConfig {
	return ServerConfig{addr: address}
}

type MSServer struct {
	MS     storage.MemStorage
	Config ServerConfig
	Router *gin.Engine
	//tmpl   *template.Template
	tmpl string
}

func MakeNewMSServer(adress string) MSServer {
	sc := NewServerConfig(adress)
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
	//
	//tmpl := `
	//<!DOCTYPE html>
	//<html>
	//	<head>
	//		<meta charset="UTF-8">
	//		<title>{{.Title}}</title>
	//	</head>
	//	<body>
	//		<div>
	//			{{.Items}}
	//		</div>
	//	</body>
	//</html>`

	//t, err := template.New("data").Parse(tmpl)
	//if err != nil {
	//	return fmt.Errorf("template parsing problem  %w", err)
	//}
	//s.Router.SetHTMLTemplate(t)
	//
	//s.tmpl = t
	//s.tmpl = tmpl

	return s.Router.Run(s.Config.addr)
}
