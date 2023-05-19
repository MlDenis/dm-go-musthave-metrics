package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tpl = `
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

func (s *MSServer) GetMSDataHowHTML(ctx *gin.Context) {

	pageInfo, err := s.MS.GetHTMLPageInfo()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", err)
	}

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
