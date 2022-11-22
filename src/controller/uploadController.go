package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) Upload(c *gin.Context) {
	// validate pdf file type

	if hash, err := controller.service.Upload(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "uploaded",
			"hash":    hash,
		})
	}
}
