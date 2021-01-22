package router

import (
	"github.com/mukesh0513/RxSecure/internal/controller"
	"github.com/mukesh0513/RxSecure/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	api := controller.Controller{}

	routes := r.Group("/", gin.BasicAuth(gin.Accounts{
		"username1": "password1",
		"username2": "password2",
		"username3": "password3",
	}))
	{
		routes.POST("/", api.Create)
		routes.GET("/:id", api.Fetch)
		routes.DELETE("/", api.Delete)
	}

	encryptionKeyGenerateroute := r.Group("/generate_key", gin.BasicAuth(gin.Accounts{
		"who_are_you": "batman",
	}))
	{
		encryptionKeyGenerateroute.POST("/", api.GenerateEncryptionKeys)
	}

	return r
}
