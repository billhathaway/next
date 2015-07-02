package next

import (
	"time"
)

// Minute returns a time.Time referencing the beginning of the next minute
func Minute(t time.Time) time.Time {
	return t.Add(time.Minute).Truncate(time.Minute)
}

// Hour returns a time.Time referencing the beginning of the next hour
func Hour(t time.Time) time.Time {
	return t.Add(time.Hour).Truncate(time.Hour)
}

// Day returns a time.Time referencing the beginning of the next day
func Day(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
}

// Week returns a time.Time referencing the beginning of the next week Sun->Sat
func Week(t time.Time) time.Time {
	weekRemainder := 7 - int(t.Weekday())
	return time.Date(t.Year(), t.Month(), t.Day()+weekRemainder, 0, 0, 0, 0, t.Location())
}

// Month returns a time.Time referencing the beginning of the next month
func Month(t time.Time) time.Time {
	// truncate starting time back to beginning of the month
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
}
