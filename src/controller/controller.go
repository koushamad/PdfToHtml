package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/PdfToHtml/src/service"
)

type Controller struct {
	engine  *gin.Engine
	service *service.Service
}

func New(engine *gin.Engine) *Controller {
	return &Controller{
		engine:  engine,
		service: service.New(),
	}
}
