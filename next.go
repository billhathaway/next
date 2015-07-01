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
	t = t.AddDate(0, 0, 1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Week returns a time.Time referencing the beginning of the next week Sun->Sat
func Week(t time.Time) time.Time {
	weekRemainder := 7 - int(t.Weekday())
	nextWeek := t.AddDate(0, 0, weekRemainder)
	return time.Date(nextWeek.Year(), nextWeek.Month(), nextWeek.Day(), 0, 0, 0, 0, t.Location())
}

// Month returns a time.Time referencing the beginning of the next month
func Month(t time.Time) time.Time {
	// truncate starting time back to beginning of the month
	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return t.AddDate(0, 1, 0)
}
