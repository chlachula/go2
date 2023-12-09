package diffsecs

import (
	"testing"
	"time"
)

func expectedDiffDatesSecs(t *testing.T, expectedSec float64, expectedLeap int, y1 int, m1 time.Month, d1, h1, mi1, s1 int, y2 int, m2 time.Month, d2, h2, mi2, s2 int) {
	date1 := leapDate(y1, m1, d1, h1, mi1, s1, 0, time.UTC)
	date2 := leapDate(y2, m2, d2, h2, mi2, s2, 0, time.UTC)
	sec, leap, err := datesDiffInSeconds(date1, date2)
	if err != nil {
		t.Errorf(err.Error())
	}
	f := "2006.01.02 15h"
	if sec != expectedSec {
		t.Errorf("unexpected seconds %f instead of %f between %s and %s", sec, expectedSec, date1.time.Format(f), date2.time.Format(f))
	}
	if leap != expectedLeap {
		t.Errorf("unexpected leap seconds %d instead of %d between %s and %s", leap, expectedLeap, date1.time.Format(f), date2.time.Format(f))
	}
}

func TestInc1sec(t *testing.T) {
	Verbose = false
	expectedSec := 86400.0
	expectedDiffDatesSecs(t, expectedSec, 0, 1972, time.January, 1, 0, 0, 0, 1972, time.January, 2, 0, 0, 0)

	expectedSec = 86401.0
	expectedDiffDatesSecs(t, expectedSec, 1, 1973, time.December, 31, 12, 0, 0, 1974, time.January, 1, 12, 0, 0)
	expectedDiffDatesSecs(t, expectedSec, 1, 1981, time.June, 30, 12, 0, 0, 1981, time.July, 1, 12, 0, 0)
	expectedDiffDatesSecs(t, expectedSec, 1, 2016, time.December, 31, 12, 0, 0, 2017, time.January, 1, 12, 0, 0)

	expectedSec = 86400.0*(1983-1981)*365 + 1 + 3
	expectedDiffDatesSecs(t, expectedSec, 3, 1981, time.June, 30, 23, 59, 59, 1983, time.July, 1, 0, 0, 0)

	expectedSec = float64(((2021-1971)*365+13)*86400 + 27)
	expectedDiffDatesSecs(t, expectedSec, 27, 1971, time.January, 1, 12, 0, 0, 2021, time.January, 1, 12, 0, 0)
}
