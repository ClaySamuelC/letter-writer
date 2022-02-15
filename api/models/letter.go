package models

type Letter struct {
	FromName string
	ToName   string
	Heading  string
	Date     string
	Address  string
	Body     string
}

func (l Letter) ToSlice() []string {
	return []string{l.FromName, l.ToName, l.Heading, l.Date, l.Address, l.Body}
}
