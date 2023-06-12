package helper

import (
	"time"
)

// LocalTime is a type alias for time.Time
type LocalTime time.Time

// Although the actual type of the data type is time.Time, it does not inherit the built-in method of time.Time,
// so we need to rewrite part of built-in method, and MarshalJSON is exactly what we need to reimplement.
// MarshalJSON serializes the LocalTime to JSON
func (lt LocalTime) MarshalJSON() ([]byte, error) {
	t := time.Time(lt)
	if t.IsZero() {
		return []byte(`""`), nil
	}

	date := t.Format("Monday, 02 January 2006 15:04 MST")
	return []byte(`"` + date + `"`), nil
}
