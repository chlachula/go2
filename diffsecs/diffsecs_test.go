package diffsecs

import (
	"fmt"
	"testing"
	"time"
)

func TestInc1sec(t *testing.T) {
	//	want := "lightgray"
	//	got := monthColor(date)
	leap := time.Date(2015, time.June, 30, 23, 59, 59, 999999999, time.UTC)

	d1 := time.Date(1972, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(1972, time.January, 2, 0, 0, 0, 0, time.UTC)
	d1a, d2a := inc1sec(d1, d2, leap)
	sec, err := DatesDiffInSeconds(d1a, d2a)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != 86400.0 {
		fmt.Println("seconds", sec)
	}

	d1 = time.Date(1973, time.December, 31, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(1974, time.January, 1, 12, 0, 0, 0, time.UTC)
	d1a, d2a = inc1sec(d1, d2, leap)
	sec, err = DatesDiffInSeconds(d1a, d2a)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != 86401.0 {
		fmt.Println("seconds", sec)
	}

	d1 = time.Date(1981, time.June, 30, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(1981, time.July, 1, 12, 0, 0, 0, time.UTC)
	d1a, d2a = inc1sec(d1, d2, leap)
	sec, err = DatesDiffInSeconds(d1a, d2a)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != 86401.0 {
		fmt.Println("seconds", sec)
	}

	d1 = time.Date(2016, time.December, 31, 12, 0, 0, 0, time.UTC)
	d2 = time.Date(2017, time.January, 1, 12, 0, 0, 0, time.UTC)
	d1a, d2a = inc1sec(d1, d2, leap)
	sec, err = DatesDiffInSeconds(d1a, d2a)
	if err != nil {
		t.Errorf(err.Error())
	}
	if sec != 86401.0 {
		t.Errorf("error:unexpected seconds %f between %s..%s", sec, d1.Format("2006.01.02 15h"), d2.Format("2006.01.02 15h"))
	}

}
