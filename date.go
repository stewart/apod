package apod

import "time"

// the APOD date format
const dateFormat = "2006-01-02"

// APOD uses a YYYY-MM-DD format for dates
type Date struct {
	time.Time
}

// Creates a new Date from the provided string, in the format YYYY-MM-DD
func NewDate(s string) (Date, error) {
	date := Date{}
	if err := date.UnmarshalJSON([]byte(s)); err != nil {
		return date, err
	}
	return date, nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	d.Time, err = time.Parse(dateFormat, string(b))
	return
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

// Returns the Date in string format - YYYY-MM-DD
func (d *Date) String() string {
	return d.Time.Format(dateFormat)
}
