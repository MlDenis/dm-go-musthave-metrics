package server

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func (s *MSServer) GetMSDataHowHTML(ctx *gin.Context) {

	pageInfo := s.MS.GetHTMLPageInfo()

	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			<div>
				{{.Items}}
			</div>
		</body>
	</html>`

	data := struct {
		Title template.HTML
		Items string
	}{
		Title: "<em>MS Data</em>",
		Items: pageInfo,
	}

	t, err := template.New("data").Parse(tpl)
	if err != nil {
		return
	}

	s.Router.SetHTMLTemplate(t)

	ctx.HTML(http.StatusOK, "data", data)
}
