package main

import (
	"log"
	"os"

	"github.com/ClaySamuelC/letter-writer/api/controllers"
	"github.com/ClaySamuelC/letter-writer/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connection{}
	err := db.Init()
	if err != nil {
		log.Printf("Error creating db")
	}

	err = db.WriteHeader([]string{"FromName", "ToName", "Heading", "Date", "Address", "Body"})
	if err != nil {
		log.Printf("Error initializing db")
	}

	letterController := controllers.LetterController{}
	letterController.Init(&db)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/letters", func(c *gin.Context) {
		log.Printf("Getting letters")
		c.JSON(200, gin.H{
			"message": "Here are all the Letters on file!",
		})
	})

	router.POST("/api/letters", letterController.CreateLetter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
