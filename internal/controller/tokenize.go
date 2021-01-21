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
		Data: post,
	}

	c.JSON(200, data)
}

func (base *Controller) Create(c *gin.Context) {
	payload := new(model.Payload)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	token, err := service.CreateToken(c, base.DB, payload)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, token)
}

func (base *Controller) Fetch(c *gin.Context) {
	data := new(model.Data)

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	token, err := service.GetToken(c, base.DB, data)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, token)
}