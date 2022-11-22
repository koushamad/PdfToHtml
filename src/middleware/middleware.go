package middleware

import "github.com/gin-gonic/gin"

type Middleware struct {
	engine *gin.Engine
}

func New(engine *gin.Engine) *Middleware {
	return &Middleware{engine: engine}
}

func (m *Middleware) Invoke() {
	m.engine.Use(m.Logger())
}
