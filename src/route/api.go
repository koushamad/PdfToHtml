package route

import (
	"github.com/koushamad/PdfToHtml/src/controller"
)

func (r Route) API() {
	c := controller.New(r.engine)
	r.engine.GET("/ping", c.Ping)
	r.engine.POST("/upload", c.Upload)
}
