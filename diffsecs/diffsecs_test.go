package diffsecs

import (
	"fmt"
	"testing"
	"time"
)

func diffDatesSecs(y1 int, m1 time.Month, d1, h1, mi1, s1 int, y2 int, m2 time.Month, d2, h2, mi2, s2 int) (float64, error) {
	date1 := time.Date(y1, m1, d1, h1, mi1, s1, 0, time.UTC)
	date2 := time.Date(y2, m2, d2, h2, mi2, s2, 0, time.UTC)
	sec, err := DatesDiffInSeconds(date1, date2)
	return sec, err
}

func TestInc1sec(t *testing.T) {
	expectedSec := 86401.0

	/*
		d1 := time.Date(1972, time.January, 1, 0, 0, 0, 0, time.UTC)
		d2 := time.Date(1972, time.January, 2, 0, 0, 0, 0, time.UTC)
		sec, err := DatesDiffInSeconds(d1, d2)
	*/
	sec, err := diffDatesSecs(1972, time.January, 1, 0, 0, 0, 1972, time.January, 2, 0, 0, 0)
	if err != nil {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, 86400.0, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}
	if sec != 86400.0 {
		fmt.Println("seconds", sec)
	}

	d1 = time.Date(1973, time.December, 31, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(1974, time.January, 1, 12, 0, 0, 0, time.UTC)
	sec, err = DatesDiffInSeconds(d1, d2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != expectedSec {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, expectedSec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}

	d1 = time.Date(1981, time.June, 30, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(1981, time.July, 1, 12, 0, 0, 0, time.UTC)
	sec, err = DatesDiffInSeconds(d1, d2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != expectedSec {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, expectedSec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}

	d1 = time.Date(2016, time.December, 31, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(2017, time.January, 1, 12, 0, 0, 0, time.UTC)
	sec, err = DatesDiffInSeconds(d1, d2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != expectedSec {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, expectedSec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}

	d1 = time.Date(2016, time.December, 24, 23, 59, 55, 0, time.UTC)
	d2 = time.Date(2016, time.December, 25, 0, 0, 5, 0, time.UTC)
	sec, err = DatesDiffInSeconds(d1, d2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedSec = float64(10)
	if sec != expectedSec {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, expectedSec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}

	verbose = true
	d1 = time.Date(2016, time.December, 31, 23, 59, 55, 0, time.UTC)
	d2 = time.Date(2017, time.January, 1, 0, 0, 5, 0, time.UTC)
	sec, err = DatesDiffInSeconds(d1, d2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedSec = float64(11)
	if sec != expectedSec {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, expectedSec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}
	verbose = false

	d1 = time.Date(1971, time.January, 1, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(2021, time.January, 1, 12, 0, 0, 0, time.UTC)
	sec, err = DatesDiffInSeconds(d1, d2)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectedSec = float64(((2021-1971)*365+13)*86400 + 27)
	if sec != expectedSec {
		t.Errorf("error:unexpected seconds %f instead of %f between %s..%s", sec, expectedSec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}

}
