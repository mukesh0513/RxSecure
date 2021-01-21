package main

import (
	"github.com/mukesh0513/RxSecure/internal/config"
	"github.com/mukesh0513/RxSecure/internal/database"
	"github.com/mukesh0513/RxSecure/pkg/router"
	"log"
)

func init() {
	config.Setup()
	database.Setup()
}

func main() {
	config := config.GetConfig()

	db := database.GetDB()
	r := router.Setup(db)

	log.Printf("Server is starting at 127.0.0.1:%s", config.Server.Port)
	r.Run("127.0.0.1:" + config.Server.Port)
}
