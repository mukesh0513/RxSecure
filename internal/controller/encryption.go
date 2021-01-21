package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/service"
)

func (base *Controller) GenerateEncryptionKeys(c *gin.Context) {
	token, err := service.GenerateEncryptionKeys(c)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, token)
}
