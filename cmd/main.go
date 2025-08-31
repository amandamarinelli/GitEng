package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const gitClientIdName = "GITHUB_CLIENT_ID"
const gitSecretName = "GITHUB_CLIENT_ID"

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gitClientId := os.Getenv(gitClientIdName)
	gitSecret := os.Getenv(gitSecretName)

	if gitClientId == "" || gitSecret == "" {
		panic("GitHub client ID or secret not set in environment variables")
	}

	// Initialize Router
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	err = router.Run()
	if err != nil {
		panic(err)
	}

}
