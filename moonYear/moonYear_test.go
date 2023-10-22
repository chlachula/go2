package moonYear

import (
	"testing"
	"time"
)

func TestMonthColor(t *testing.T) {
	want := "lightgray"
	date := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	got := monthColor(date)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	want = "lightblue"
	date = time.Date(2009, time.December, 10, 23, 0, 0, 0, time.UTC)
	got = monthColor(date)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// isFirstQuaterFriday(date time.Time, moonAngle float64) bool {
func TestIsFirstQuaterFriday(t *testing.T) {
	date := time.Date(2023, time.October, 22, 12, 0, 0, 0, time.UTC)
	want := false
	got := isFirstQuaterFriday(date, 0.0)
	if got != want {
		t.Errorf("#1 got %t, want %t", got, want)
	}

	date = time.Date(2023, time.October, 20, 12, 0, 0, 0, time.UTC)
	want = true
	got = isFirstQuaterFriday(date, 90.0)
	if got != want {
		t.Errorf("#2 got %t, want %t", got, want)
	}
}
