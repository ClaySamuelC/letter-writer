package db

import (
	"encoding/csv"
	"os"
)

type Connection struct {
	db *csv.Writer
}

func (c *Connection) Init() error {
	recordFile, err := os.Create("./letters.csv")

	c.db = csv.NewWriter(recordFile)

	return err
}

func (c *Connection) WriteHeader(headerModel []string) error {
	return c.db.Write(headerModel)
}

func (c *Connection) AddEntry(entry []string) error {
	return c.db.Write(entry)
}
