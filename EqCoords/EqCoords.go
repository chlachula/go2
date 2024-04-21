package EqCoords

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type (
	SphericalCoordsRad struct {
		RA float64
		DE float64
	}
)

var (
	inputJD  = 2451545.0     // J2000.0 = January 1, 2000 at 12:00 TT
	outputJD = 2460676.50000 // 2025-01-01 0:00 UTC
	reHDMS   = regexp.MustCompile(`(?m)(?P<sign>\+|\-|)(?P<d>\d{1,2})[^\d](?P<m>\d{1,2})[^\d](?P<s>(\d+(?:\.\d+)?))[^\d]`)
)

func regexParamsMap(compRegEx *regexp.Regexp, str string) (paramsMap map[string]string) {
	//var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(str)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func HDMSvalues(hours bool, dmsString string) (float64, error) {
	sign, hd, mm, ss := 1.0, 0.0, 0.0, 0.0

	m := regexParamsMap(reHDMS, dmsString)
	if m["sign"] == "-" {
		sign = -1.0
	}
	if a, err := strconv.ParseInt(m["d"], 10, 32); err == nil {
		hd = float64(a)
	} else {
		return 0.0, fmt.Errorf("error parsing deg or hour '%s': %s", dmsString, err.Error())
	}
	if a, err := strconv.ParseInt(m["m"], 10, 32); err == nil {
		mm = float64(a)
	} else {
		return 0.0, fmt.Errorf("error parsing minutes '%s': %s", dmsString, err.Error())
	}
	if a, err := strconv.ParseFloat(m["s"], 64); err == nil {
		ss = a
	} else {
		return 0.0, fmt.Errorf("error parsing seconds '%s': %s", dmsString, err.Error())
	}
	f := 1.0
	if hours {
		f = 15.0
	}
	return sign * f * (hd + (mm/60.0 + ss/3600.0)) * math.Pi / 180.0, nil
}
func DegreesStr2rad(degStr string) (float64, error) {
	if value, err := strconv.ParseFloat(degStr, 64); err == nil {
		return value * math.Pi / 180.0, nil
	} else {
		return value, err
	}
}
func DegMinSecStr2rad(degStr string) (float64, error) {
	return HDMSvalues(false, degStr)
}
func HourMinSecStr2rad(degStr string) (float64, error) {
	return HDMSvalues(true, degStr)
}

/*
Julian 5 October 1582 = Gregorian 15 October 1582 = JDN 2299161
JDN at 12:00 UT YYYY-MM-DD
*/
func gregorianNoonToJulianDayNumber(Y int, M int, D int) int {
	JDN := (1461*(Y+4800+(M-14)/12))/4 + (367*(M-2-12*((M-14)/12)))/12 - (3*((Y+4900+(M-14)/12)/100))/4 + D - 32075
	return JDN
}

func timeToJulianDay(t time.Time) float64 {
	t1 := t.UTC()
	t2 := t1.Add(-12 * time.Hour)
	y := t2.Year()
	m := int(t2.Month())
	d := t2.Day()
	jdn := gregorianNoonToJulianDayNumber(y, m, d)
	sec := t2.Hour()*3600 + t2.Minute()*60 + t2.Second()
	frac := (float64(sec) + float64(t2.Nanosecond())/1000000000.0) / 86400.0
	return float64(jdn) + frac
}
func SetOutputTime(str string) {
	if t, err := time.Parse("2006-01-02", str); err == nil {
		outputJD = timeToJulianDay(t)
	} else {
		fmt.Printf("Error at output date %s\n", str)
	}
}

// 18h36m56.33635s,+38°47'01.2802" //J2000.0 Vega
func ConvertCoords(str string) {
	c1 := ConvertCoordsStr2Rad(str)
	c2 := ConvertCoordToJD(inputJD, outputJD, c1)
	fmt.Printf("Coordinates: input %v, output %v", c1, c2)
}
func ConvertCoordToJD(j1, j2 float64, c1 SphericalCoordsRad) SphericalCoordsRad {
	var c2 SphericalCoordsRad
	return c2
}
func ConvertCoordsStr2Rad(str string) SphericalCoordsRad {
	var c SphericalCoordsRad
	a := strings.Split(str, ",")
	if i := len(a); i != 2 {
		fmt.Printf("error: string '%s' split into %d parts instead of 2.", str, i)
		return c
	}
	strings.Split(str, ",")
	delta := outputJD - inputJD
	fmt.Printf("Days since J2000.0: %.3f \n", delta)

	var err error
	if c.RA, err = HourMinSecStr2rad(a[0]); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	if c.DE, err = DegMinSecStr2rad(a[1]); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("radians: RA %f, DE %f \n", c.RA, c.DE)
	return c
}
