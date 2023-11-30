package diffsecs

import (
	"fmt"
	"time"
)

/*
Year	30 Jun	31 Dec
1972	+1	+1
1973	0	+1
1974	0	+1
1975	0	+1
1976	0	+1
1977	0	+1
1978	0	+1
1979	0	+1
1980	0	0
1981	+1	0
1982	+1	0
1983	+1	0
1984	0	0
1985	+1	0
1986	0	0
1987	0	+1
1988	0	0
1989	0	+1
1990	0	+1
1991	0	0
1992	+1	0
1993	+1	0
1994	+1	0
1995	0	+1
1996	0	0
1997	+1	0
1998	0	+1
1999	0	0
2000	0	0
2001	0	0
2002	0	0
2003	0	0
2004	0	0
2005	0	+1
2006	0	0
2007	0	0
2008	0	+1
2009	0	0
2010	0	0
2011	0	0
2012	+1	0
2013	0	0
2014	0	0
2015	+1	0
2016	0	+1
2017	0	0
2018	0	0
2019	0	0
2020	0	0
2021	0	0
2022	0	0
2023	0	0
Year	30 Jun	31 Dec
Total	11	16
27
Current TAI − UTC
37
*/
type YYYYLeapSeconds struct {
	YYYY  int
	Jun30 int
	Dec31 int
}

type LeapSecondsTime struct {
	leapSeconds int
	time        time.Time
}
type LeapSecondsDate struct {
	year  int
	month time.Month
	day   int
}

func LeapDate(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) LeapSecondsTime {
	var t LeapSecondsTime
	t.time = time.Date(year, month, day, hour, min, sec, nsec, loc)
	t.leapSeconds = 10
	i1 := 0
	time1 := time.Date(leapDates[i1].year, leapDates[i1].month, leapDates[i1].day, 23, 59, 59, 999999999, time.UTC)
	if t.time.Before(time1) {
		return t
	}
	i2 := len(leapDates) - 1
	time2 := time.Date(leapDates[i2].year, leapDates[i2].month, leapDates[i2].day, 23, 59, 59, 999999999, time.UTC)
	if t.time.After(time2) {
		t.leapSeconds += len(leapDates)
		return t
	}
	i := i2 / 2
	for i1 < i2 {
		time := time.Date(leapDates[i].year, leapDates[i].month, leapDates[i].day, 23, 59, 59, 999999999, time.UTC)
		if t.time.After(time) {
			i1 = i
			time1 = time
		} else {
			i2 = i
			time2 = time
		}
		i = i1 + (i2-i1)/2
	}
	t.leapSeconds += i + 1
	return t
}

var verbose = false
var iDEB1 int
var iDEB2 int

var secs = []YYYYLeapSeconds{
	{YYYY: 1972, Jun30: 1, Dec31: 1},
	{YYYY: 1973, Jun30: 0, Dec31: 1},
	{YYYY: 1974, Jun30: 0, Dec31: 1},
	{YYYY: 1975, Jun30: 0, Dec31: 1},
	{YYYY: 1976, Jun30: 0, Dec31: 1},
	{YYYY: 1977, Jun30: 0, Dec31: 1},
	{YYYY: 1978, Jun30: 0, Dec31: 1},
	{YYYY: 1979, Jun30: 0, Dec31: 1},
	{YYYY: 1980, Jun30: 0, Dec31: 0},
	{YYYY: 1981, Jun30: 1, Dec31: 0},
	{YYYY: 1982, Jun30: 1, Dec31: 0},
	{YYYY: 1983, Jun30: 1, Dec31: 0},
	{YYYY: 1984, Jun30: 0, Dec31: 0},
	{YYYY: 1985, Jun30: 1, Dec31: 0},
	//	{YYYY: 1986, Jun30: 0, Dec31: 0},
	{YYYY: 1987, Jun30: 0, Dec31: 1},
	//	{YYYY: 1988, Jun30: 0, Dec31: 0},
	{YYYY: 1989, Jun30: 0, Dec31: 1},
	{YYYY: 1990, Jun30: 0, Dec31: 1},
	//	{YYYY: 1991, Jun30: 0, Dec31: 0},
	{YYYY: 1992, Jun30: 1, Dec31: 0},
	{YYYY: 1993, Jun30: 1, Dec31: 0},
	{YYYY: 1994, Jun30: 1, Dec31: 0},
	{YYYY: 1995, Jun30: 0, Dec31: 1},
	//	{YYYY: 1996, Jun30: 0, Dec31: 0},
	{YYYY: 1997, Jun30: 1, Dec31: 0},
	{YYYY: 1998, Jun30: 0, Dec31: 1},
	//	{YYYY: 1999, Jun30: 0, Dec31: 0},
	//	{YYYY: 2000, Jun30: 0, Dec31: 0},
	//	{YYYY: 2001, Jun30: 0, Dec31: 0},
	//	{YYYY: 2002, Jun30: 0, Dec31: 0},
	//	{YYYY: 2003, Jun30: 0, Dec31: 0},
	//	{YYYY: 2004, Jun30: 0, Dec31: 0},
	{YYYY: 2005, Jun30: 0, Dec31: 1},
	//	{YYYY: 2006, Jun30: 0, Dec31: 0},
	//	{YYYY: 2007, Jun30: 0, Dec31: 0},
	{YYYY: 2008, Jun30: 0, Dec31: 1},
	//	{YYYY: 2009, Jun30: 0, Dec31: 0},
	//	{YYYY: 2010, Jun30: 0, Dec31: 0},
	//	{YYYY: 2011, Jun30: 0, Dec31: 0},
	{YYYY: 2012, Jun30: 1, Dec31: 0},
	//	{YYYY: 2013, Jun30: 0, Dec31: 0},
	//	{YYYY: 2014, Jun30: 0, Dec31: 0},
	{YYYY: 2015, Jun30: 1, Dec31: 0},
	{YYYY: 2016, Jun30: 0, Dec31: 1},
	//	{YYYY: 2017, Jun30: 0, Dec31: 0},
	//	{YYYY: 2018, Jun30: 0, Dec31: 0},
	//	{YYYY: 2019, Jun30: 0, Dec31: 0},
	//	{YYYY: 2020, Jun30: 0, Dec31: 0},
	//	{YYYY: 2021, Jun30: 0, Dec31: 0},
	//	{YYYY: 2022, Jun30: 0, Dec31: 0},
	//	{YYYY: 2023, Jun30: 0, Dec31: 0},
}
var leapDates = []LeapSecondsDate{
	/* 01 */ {year: 1972, month: time.June, day: 30},
	/* 02 */ {year: 1972, month: time.December, day: 31},
	/* 03 */ {year: 1973, month: time.December, day: 31},
	/* 04 */ {year: 1974, month: time.December, day: 31},
	/* 05 */ {year: 1975, month: time.December, day: 31},
	/* 06 */ {year: 1976, month: time.December, day: 31},
	/* 07 */ {year: 1977, month: time.December, day: 31},
	/* 08 */ {year: 1978, month: time.December, day: 31},
	/* 09 */ {year: 1979, month: time.December, day: 31},
	/* 10 */ {year: 1981, month: time.June, day: 30},
	/* 11 */ {year: 1982, month: time.June, day: 30},
	/* 12 */ {year: 1983, month: time.June, day: 30},
	/* 13 */ {year: 1985, month: time.June, day: 30},
	/* 14 */ {year: 1987, month: time.December, day: 31},
	/* 15 */ {year: 1989, month: time.December, day: 31},
	/* 16 */ {year: 1990, month: time.December, day: 31},
	/* 17 */ {year: 1992, month: time.June, day: 30},
	/* 18 */ {year: 1993, month: time.June, day: 30},
	/* 19 */ {year: 1994, month: time.June, day: 30},
	/* 20 */ {year: 1995, month: time.December, day: 31},
	/* 21 */ {year: 1997, month: time.June, day: 30},
	/* 22 */ {year: 1998, month: time.December, day: 31},
	/* 23 */ {year: 2005, month: time.December, day: 31},
	/* 24 */ {year: 2008, month: time.December, day: 31},
	/* 25 */ {year: 2012, month: time.June, day: 30},
	/* 26 */ {year: 2015, month: time.June, day: 30},
	/* 27 */ {year: 2016, month: time.December, day: 31},
}

func Tmp() {
	i := 0
	lineFormat := "/* %02d */ {year:%d, month:time.%s, day:%02d}, \n"
	for _, x := range secs {
		if x.Jun30 != 0 {
			i += 1
			fmt.Printf(lineFormat, i, x.YYYY, "June", 30)
		}
		if x.Dec31 != 0 {
			i += 1
			fmt.Printf(lineFormat, i, x.YYYY, "December", 31)
		}
	}
}
func ShowLeapSeconds() {
	total := 0
	for _, r := range secs {
		fmt.Printf("%d: %3d, %3d\n", r.YYYY, r.Jun30, r.Dec31)
		total += r.Jun30 + r.Dec31
	}
	fmt.Printf("Total leap seconds: %d\n", total)
}
func inc1sec(d1, d2, leap time.Time) (d1a, d2a time.Time) {
	d1a = d1
	if d1.After(leap) {
		d1a = d1.Add(time.Second)
		if verbose {
			iDEB1 += 1
			fmt.Printf("%2d.DEBUG1 +1s %s %s\n", iDEB1, d1.String(), leap.String())
		}
	}
	d2a = d2
	if d2.After(leap) {
		d2a = d2.Add(time.Second)
		if verbose {
			iDEB2 += 1
			fmt.Printf("%2d.DEBUG2 +1s %s %s\n", iDEB2, d2.String(), leap.String())
		}
	}
	return d1a, d2a
}
func DatesDiffInSeconds(d1, d2 time.Time) (float64, error) {
	if d1.After(d2) {
		return 0.0, fmt.Errorf("the first date is after second date")
	}
	if verbose {
		iDEB1 = 0
		iDEB2 = 0
		fmt.Println("DEBUG DatesDiffInSeconds ", d1.Format("2006.01.02 15h"), "..", d2.Format("2006.01.02 15h"))
	}
	t1972_01_01 := time.Date(1972, 1, 1, 0, 0, 0, 0, time.UTC)
	if d2.After(t1972_01_01) {
		for _, r := range secs {
			if r.Jun30 != 0 {
				tJun30 := time.Date(r.YYYY, time.June, 30, 23, 59, 59, 999999999, time.UTC)
				d1, d2 = inc1sec(d1, d2, tJun30)
			}
			if r.Dec31 != 0 {
				tDec31 := time.Date(r.YYYY, time.December, 31, 23, 59, 59, 999999999, time.UTC)
				d1, d2 = inc1sec(d1, d2, tDec31)
			}
		}
	}

	diff := d2.Sub(d1)
	return diff.Seconds(), nil
}
func SecondsDiff(d1, d2 time.Time) {
	diffSec, err := DatesDiffInSeconds(d1, d2)
	if err != nil {
		fmt.Println("diff", diffSec, "seconds")
	} else {
		fmt.Println("Error: ", err.Error())
	}
}
