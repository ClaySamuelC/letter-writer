package controllers

import (
	"log"

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
