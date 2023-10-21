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
