package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/service"
)

func (base *Controller) Create(c *gin.Context) {
	payload := new(model.Payload)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	token, err := service.CreateToken(c, payload)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, token)
}

func (base *Controller) Fetch(c *gin.Context) {
	//var args model.GetApiParams
	//
	//args.Token = c.DefaultQuery("token", "null")

	token, _ := c.Params.Get("id")

	token, err := service.GetToken(c, token)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, token)
}

func (base *Controller) Delete(c *gin.Context) {
	c.JSON(200, nil)
}
