package controllers

import (
	"encoding/csv"

	"github.com/ClaySamuelC/letter-writer/api/models"
	"github.com/gin-gonic/gin"
)

type LetterController struct {
	db *csv.Writer
}

func (c *LetterController) Init(db *csv.Writer) {
	c.db = db
}

func (c *LetterController) CreateLetter(ctx *gin.Context) {
	var letter models.Letter
	ctx.BindJSON(&letter)

	if letter.Heading == "" {
		ctx.JSON(400, gin.H{
			"error": "heading should not be empty",
		})
		return
	}
	if err := c.db.Write(letter.ToSlice()); err != nil {
		ctx.JSON(500, gin.H{
			"error": "error writing letter to db",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"letter": letter,
	})
}
