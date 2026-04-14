package main

import (
	"log"
	"os"

	"github.com/boybenson/pay-server/internals/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	router := gin.Default()

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&auth.User{})

	if err != nil {
		panic("failed to migrate database")
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authService := auth.NewService(db)
	authHandler := auth.NewHandler(authService)

	router.POST("/signup", authHandler.CreateUserHandler)
	router.POST("/signin", authHandler.SignInHandler)

	router.Run()
}
