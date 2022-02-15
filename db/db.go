package db

import (
	"encoding/csv"
	"log"
	"os"
)

type Connection struct {
	db *csv.Writer
}

func (c *Connection) Init() {
	recordFile, err := os.Create("./letters.csv")
	if err != nil {
		log.Printf("Error creating db file")
	}

	c.db = csv.NewWriter(recordFile)
}

func (c *Connection) WriteHeader(headerModel []string) error {
	return c.db.Write(headerModel)
}

func (c *Connection) AddEntry(entry []string) error {
	return c.db.Write(entry)
}
