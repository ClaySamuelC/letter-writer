package letter

type Letter struct {
	fromName string
	toName   string
	heading  string
	date     string
	address  string
	body     string
}

func CreateLetter(fromName, toName, heading, date, address, body string) Letter {
	var newLetter Letter

	newLetter.fromName = fromName
	newLetter.toName = toName
	newLetter.heading = heading
	newLetter.date = date
	newLetter.address = address
	newLetter.body = body

	return newLetter
}
