package main

import (
	"mhlv/cmd"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/coordinate-agent"),
		gin.Recovery(),
	)

	cmd.Boot(router)
	router.Use(cors.New(cors.Config{AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"}}))
	router.Run("localhost:8080")
}
