package starTime

import (
	"fmt"
	"time"
)

const WikiDefinition = `Earth's rotation period 
relative to the International Celestial Reference Frame, 
called its stellar day by the International Earth Rotation 
and Reference Systems Service (IERS), 
is 86 164.098 903 691 seconds of mean solar time (UT1) 
(23h 56m 4.098903691s, 0.99726966323716 mean solar days).
Earth's rotation period relative to the precessing mean vernal equinox, 
named sidereal day, is 86164.09053083288 seconds of mean solar time (UT1) 
(23h 56m 4.09053083288s, 0.99726956632908 mean solar days).
Thus, the sidereal day is shorter than the stellar day 
by about 8.4 ms (jch: i.e. 3.066s per year).

`
const StellarDayNs time.Duration = 86164098903691

//const Stellar_DayS = 86164.09890369100
//const SideReaDayNs = 86164090530833
//const SideRealDayS = 86164.09053083288

func Compute(t1str, t2str string) {
	t1, _ := time.Parse("2006-01-02_15:04:05", t1str)
	t1int := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	t2, _ := time.Parse("2006-01-02", t2str)
	diff := t2.Sub(t1int)
	d := diff.Hours() / 24
	plus := time.Duration(d) * StellarDayNs * time.Nanosecond
	t2s := t1.Add(plus)
	fmt.Println("days diff ", d)
	fmt.Println("New time ", t2s)
}
