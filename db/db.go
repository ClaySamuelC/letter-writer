package db

import (
	"encoding/csv"
	"log"
	"os"
)

func NewConnection() *csv.Writer {
	recordFile, err := os.Create("./letters.csv")
	if err != nil {
		log.Printf("Error opening db")
	}

	return csv.NewWriter(recordFile)
}
