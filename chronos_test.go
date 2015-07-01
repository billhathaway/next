package next

import (
	"testing"
	"time"
)

var (
	loc = &time.Location{}
	t1  = time.Date(2006, 1, 2, 15, 4, 5, 0, loc)
)

func TestMinute(t *testing.T) {
	nextMin := time.Date(2006, 1, 2, 15, 5, 0, 0, loc)
	tmp := Minute(t1)
	if !tmp.Equal(nextMin) {
		t.Errorf("unexpected time - wanted %v got %v\n", nextMin, tmp)
	}
}
func TestHour(t *testing.T) {
	nextHour := time.Date(2006, 1, 2, 16, 0, 0, 0, loc)
	tmp := Hour(t1)
	if !tmp.Equal(nextHour) {
		t.Errorf("unexpected time - wanted %v got %v\n", nextHour, tmp)
	}
}

func TestDay(t *testing.T) {
	nextDay := time.Date(2006, 1, 3, 0, 0, 0, 0, loc)
	tmp := Day(t1)
	if !tmp.Equal(nextDay) {
		t.Errorf("unexpected time - wanted %v got %v\n", nextDay, tmp)
	}
}

func TestWeek(t *testing.T) {
	nextWeek := time.Date(2006, 1, 8, 0, 0, 0, 0, loc)
	tmp := Week(t1)
	if !tmp.Equal(nextWeek) {
		t.Errorf("unexpected time - wanted %v got %v\n", nextWeek, tmp)
	}
}
