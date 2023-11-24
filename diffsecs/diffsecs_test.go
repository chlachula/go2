package diffsecs

import (
	"testing"
	"time"
)

func TestInc1sec(t *testing.T) {
	//	want := "lightgray"
	//	got := monthColor(date)
	d1 := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2015, time.January, 2, 0, 0, 0, 0, time.UTC)
	leap := time.Date(2015, time.June, 30, 23, 59, 59, 999999999, time.UTC)
	d1a, d2a := inc1sec(d1, d2, leap)
	//to be finished
}
