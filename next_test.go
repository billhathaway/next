package next

import (
	"testing"
	"time"
)

var (
	loc = time.Now().Location()
	t1  = time.Date(2015, 7, 1, 15, 4, 30, 0, loc)
)

func TestMinute(t *testing.T) {
	expected := time.Date(2015, 7, 1, 15, 5, 0, 0, loc)
	result := Minute(t1)
	if !result.Equal(expected) {
		t.Errorf("unexpected time - wanted %v got %v\n", expected, result)
	}
}
func TestHour(t *testing.T) {
	expected := time.Date(2015, 7, 1, 16, 0, 0, 0, loc)
	result := Hour(t1)
	if !result.Equal(expected) {
		t.Errorf("unexpected time - wanted %v got %v\n", expected, result)
	}
}

func TestDay(t *testing.T) {
	expected := time.Date(2015, 7, 2, 0, 0, 0, 0, loc)
	result := Day(t1)
	if !result.Equal(expected) {
		t.Errorf("unexpected time - wanted %v got %v\n", expected, result)
	}
	t2 := time.Date(2015, 7, 1, 11, 12, 0, 0, loc)
	expected = time.Date(2015, 7, 2, 0, 0, 0, 0, loc)
	result = Day(t2)
	if !result.Equal(expected) {
		t.Errorf("unexpected time - wanted %v got %v\n", expected, result)
	}
}

func TestWeek(t *testing.T) {
	expected := time.Date(2015, 7, 5, 0, 0, 0, 0, loc)
	result := Week(t1)
	if !result.Equal(expected) {
		t.Errorf("unexpected time - wanted %v got %v\n", expected, result)
	}
}

func TestMonth(t *testing.T) {
	nextMonth := time.Date(2015, 8, 1, 0, 0, 0, 0, loc)
	result := Month(t1)
	if !result.Equal(nextMonth) {
		t.Errorf("unexpected time - wanted %v got %v\n", nextMonth, result)
	}
}
