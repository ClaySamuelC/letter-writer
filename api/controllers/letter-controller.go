package controllers

import (
	"log"
	"time"

	"github.com/ClaySamuelC/letter-writer/api/models"
	"github.com/ClaySamuelC/letter-writer/db"
	"github.com/gin-gonic/gin"
)

type LetterController struct {
	db *db.Connection
}

func (c *LetterController) Init(db *db.Connection) {
	c.db = db
}

func (c *LetterController) CreateLetter(ctx *gin.Context) {
	log.Printf("Attempting to write a letter")

	var letter models.Letter
	ctx.BindJSON(&letter)

	if letter.Heading == "" {
		ctx.JSON(400, gin.H{
			"error": "heading should not be empty",
		})
		return
	}

	letter.Date = time.Now().Format("Jan-01-2000")
	letter.Address = "123 Roosavelt Str" // Placeholder for when the user passes their address data

	if err := c.db.AddEntry(letter.ToSlice()); err != nil {
		ctx.JSON(500, gin.H{
			"error": "error writing letter to db",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"letter": letter,
	})
}
