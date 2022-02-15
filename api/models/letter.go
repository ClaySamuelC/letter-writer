package models

type Letter struct {
	ID       int64
	fromName string
	toName   string
	heading  string
	date     string
	address  string
	body     string
}
