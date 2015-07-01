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
	return t.AddDate(0, 0, 1).Truncate(24 * time.Hour)
}

// Week returns a time.Time referencing the beginning of the next week Sun->Sat
func Week(t time.Time) time.Time {
	weekRemainder := 7 - int(t.Weekday())
	return t.AddDate(0, 0, weekRemainder).Truncate(24 * time.Hour)
}
