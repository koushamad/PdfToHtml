package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func (m *Middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("request_time", t)
		c.Next()

		latency := time.Since(t)

		log.Print(latency)
		log.Println(c.Writer.Status())
	}
}
