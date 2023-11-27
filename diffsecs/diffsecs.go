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
	{YYYY: 1986, Jun30: 0, Dec31: 0},
	{YYYY: 1987, Jun30: 0, Dec31: 1},
	{YYYY: 1988, Jun30: 0, Dec31: 0},
	{YYYY: 1989, Jun30: 0, Dec31: 1},
	{YYYY: 1990, Jun30: 0, Dec31: 1},
	{YYYY: 1991, Jun30: 0, Dec31: 0},
	{YYYY: 1992, Jun30: 1, Dec31: 0},
	{YYYY: 1993, Jun30: 1, Dec31: 0},
	{YYYY: 1994, Jun30: 1, Dec31: 0},
	{YYYY: 1995, Jun30: 0, Dec31: 1},
	{YYYY: 1996, Jun30: 0, Dec31: 0},
	{YYYY: 1997, Jun30: 1, Dec31: 0},
	{YYYY: 1998, Jun30: 0, Dec31: 1},
	{YYYY: 1999, Jun30: 0, Dec31: 0},
	{YYYY: 2000, Jun30: 0, Dec31: 0},
	{YYYY: 2001, Jun30: 0, Dec31: 0},
	{YYYY: 2002, Jun30: 0, Dec31: 0},
	{YYYY: 2003, Jun30: 0, Dec31: 0},
	{YYYY: 2004, Jun30: 0, Dec31: 0},
	{YYYY: 2005, Jun30: 0, Dec31: 1},
	{YYYY: 2006, Jun30: 0, Dec31: 0},
	{YYYY: 2007, Jun30: 0, Dec31: 0},
	{YYYY: 2008, Jun30: 0, Dec31: 1},
	{YYYY: 2009, Jun30: 0, Dec31: 0},
	{YYYY: 2010, Jun30: 0, Dec31: 0},
	{YYYY: 2011, Jun30: 0, Dec31: 0},
	{YYYY: 2012, Jun30: 1, Dec31: 0},
	{YYYY: 2013, Jun30: 0, Dec31: 0},
	{YYYY: 2014, Jun30: 0, Dec31: 0},
	{YYYY: 2015, Jun30: 1, Dec31: 0},
	{YYYY: 2016, Jun30: 0, Dec31: 1},
	{YYYY: 2017, Jun30: 0, Dec31: 0},
	{YYYY: 2018, Jun30: 0, Dec31: 0},
	{YYYY: 2019, Jun30: 0, Dec31: 0},
	{YYYY: 2020, Jun30: 0, Dec31: 0},
	{YYYY: 2021, Jun30: 0, Dec31: 0},
	{YYYY: 2022, Jun30: 0, Dec31: 0},
	{YYYY: 2023, Jun30: 0, Dec31: 0},
}

func ShowLeapSeconds() {
	for _, r := range secs {
		fmt.Printf("%d: %3d, %3d\n", r.YYYY, r.Jun30, r.Dec31)

	}
}
func inc1sec(d1, d2, leap time.Time) (d1a, d2a time.Time) {
	d1a = d1
	if d1.After(leap) {
		d1a = d1.Add(time.Second)
		fmt.Println("DEBUG +1s ", d1.String(), leap.String())
	}
	d2a = d2
	if d2.After(leap) {
		d2a = d2.Add(time.Second)
		fmt.Println("DEBUG +1s ", d2.String(), leap.String())
	}
	return d1a, d2a
}
func DatesDiffInSeconds(d1, d2 time.Time) (float64, error) {
	if d1.After(d2) {
		return 0.0, fmt.Errorf("the first date is after second date")
	}
	fmt.Println("DEBUG DatesDiffInSeconds ")

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
