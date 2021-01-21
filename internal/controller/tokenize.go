package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/service"
)

func (base *Controller) Get(c *gin.Context) {
	var args model.GetApiParams

	args.Token = c.DefaultQuery("token", "null")

	// Fetch results from database
	post, err := service.GetTokenizeValue(c, base.DB, args)
	if err != nil {
		c.AbortWithStatus(500)
	}

	// Fill return data struct
	data := model.GetApiMessageData{
		Data:         post,
	}

	c.JSON(200, data)
}

func (base *Controller) Create(c *gin.Context) {
	post := new(model.Keys)

	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	post, err = service.CreateToken(c, base.DB, post)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, post)
}