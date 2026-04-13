package main

import (
	"github.com/boybenson/pay-server/internals/auth"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authHandler := auth.NewHandler()
	
	router.POST("/signup", authHandler.CreateUserHandler)

	
	router.Run() 
}
