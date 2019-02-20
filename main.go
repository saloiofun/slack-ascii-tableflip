package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saloiofun/ascii/handlers"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	r.POST("/tableflip", handlers.Tableflip)
	r.GET("/slack/authorization", handlers.Authorization)
	r.GET("/slack/callback/:code/:state", handlers.Callback)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
	// router.Run(":3000") for a hard coded port
}
