package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/PdfToHtml/src/config"
	"github.com/koushamad/PdfToHtml/src/middleware"
	"github.com/koushamad/PdfToHtml/src/route"
	"log"
)

func main() {
	engine := gin.New()

	middleware.New(engine).Invoke()
	route.New(engine).API()

	err := engine.Run(":" + config.Get("PORT"))
	log.Panicln(err)
}
