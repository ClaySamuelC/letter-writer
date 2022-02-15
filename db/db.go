package db

import (
	"encoding/csv"
	"os"
)

type Connection struct {
	f *os.File
}

func (c *Connection) Init() error {
	file, err := os.Create("./db/letters.csv")
	c.f = file

	return err
}

func (c *Connection) WriteHeader(header []string) error {
	writer := csv.NewWriter(c.f)
	defer writer.Flush()

	return writer.Write(header)
}

func (c *Connection) AddEntry(entry []string) error {
	writer := csv.NewWriter(c.f)
	defer writer.Flush()

	return writer.Write(entry)
}

func (c *Connection) CloseConnection() error {
	return c.f.Close()
}
