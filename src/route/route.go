package route

import "github.com/gin-gonic/gin"

type Route struct {
	engine *gin.Engine
}

func New(engine *gin.Engine) *Route {
	return &Route{engine: engine}
}
