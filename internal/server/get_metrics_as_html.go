package server

import (
	"github.com/MlDenis/dm-go-musthave-metrics/internal/templates"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

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
	s.Router.SetHTMLTemplate(templates.DataTemplate)

	err = templates.DataTemplate.Execute(ctx.Writer, data)

	if err != nil {
		log.Println(err)
		ctx.HTML(http.StatusInternalServerError, "error", err)
	}

	ctx.HTML(http.StatusOK, "data", data)
}
