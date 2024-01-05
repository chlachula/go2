package moonView

import (
	"encoding/json"
	"math"
	"testing"
	"time"
)

/*
"2011" => "a003800/a003810",
"2012" => "a003800/a003894",
"2013" => "a004000/a004000",
"2014" => "a004100/a004118",
"2015" => "a004200/a004236",
"2016" => "a004400/a004404",
"2017" => "a004500/a004537",
"2018" => "a004600/a004604",
"2019" => "a004400/a004442",
"2020" => "a004700/a004768",
"2021" => "a004800/a004874",
"2022" => "a004900/a004955",
"2023" => "a005000/a005048"
"2024" => "a005100/a005187"

https://svs.gsfc.nasa.gov/vis/a000000/a005100/a005187/frames/730x730_1x1_30p/moon.0001.jpg
*/
func TestSvsMagicNumbers(t *testing.T) {
	year := 2011
	wantNN00 := 3800
	wantNNNN := 3810
	gotNN00, gotNNNN := svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}

	year = 2017
	wantNN00 = 4500
	wantNNNN = 4537
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}

	year = 2024 //last available year
	wantNN00 = 5100
	wantNNNN = 5187
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}
	year = 1955
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}
	year = 2055
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}
}
func TestGregorianNoonToJulianDayNumber(t *testing.T) {
	wantJdn := 2459861
	y := 2022
	m := 10
	d := 8
	gotJdn := gregorianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}

	wantJdn = 2415020
	y = 1899
	m = 12
	d = 31
	gotJdn = gregorianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}

	wantJdn = 2451545
	y = 2000
	m = 1
	d = 1
	gotJdn = gregorianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}

	wantJdn = 2299161
	y = 1582
	m = 10
	d = 15
	gotJdn = gregorianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}
}
func TestJulianNoonToJulianDayNumber(t *testing.T) {
	wantJdn := 2299161
	y := 1582
	m := 10
	d := 5
	gotJdn := julianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}

	//	JD 0 corresponds to 1 January 4713 BC in the Julian calendar, or 24 November 4714 BC in the Gregorian calendar.
	wantJdn = 0
	y = -4713 + 1
	m = 1
	d = 1
	gotJdn = julianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For jul. date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}
	y = -4714 + 1
	m = 11
	d = 24
	gotJdn = gregorianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For greg. date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}

}
func TestFirstGregorianDayWithJulian(t *testing.T) {
	wantJdn := 2299161
	y := 1582
	m := 10
	d := 15
	gotJdn := gregorianNoonToJulianDayNumber(y, m, d)
	if wantJdn != gotJdn {
		t.Errorf("For date %d-%02d-%02d WANT: %d; GOT: %d julian date number", y, m, d, wantJdn, gotJdn)
	}
	d5 := 5
	julJdn := julianNoonToJulianDayNumber(y, m, d5)
	if julJdn != gotJdn {
		t.Errorf("JDN for greg %d-%02d-%02d is not equal for jul  %d-%02d-%02d", y, m, d, y, m, d5)
	}
}
func jTest(t *testing.T, d time.Time, wantJ float64) {
	j := timeToJulianDay(d)
	diffJ := j - wantJ
	if math.Abs(diffJ) > 0.0001 {
		t.Errorf("unexpected diff=%.7f j=%.4f want: j=%.4f %s = %s\n", diffJ, j, wantJ, d, d.UTC())
	}
}
func TestTimeToJulianDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Chicago")
	d := time.Date(2023, 12, 31, 20, 0, 0, 0, loc)
	jTest(t, d, 2460310.583333)
	d = time.Date(2024, 1, 1, 6, 0, 0, 0, loc)
	jTest(t, d, 2460311.0)
	d = time.Date(2024, 1, 1, 17, 59, 59, 999999999, loc)
	jTest(t, d, 2460311.5)
	d = time.Date(2024, 1, 2, 5, 59, 59, 999999999, loc)
	jTest(t, d, 2460312.0)
}
func TestWholeHoursSinceJanuary1(t *testing.T) {
	timeLayout := "2006-01-02 15:04 MST"
	want := 1
	jan1, _ := time.Parse(timeLayout, "2024-01-01 00:00 UTC")
	got := wholeHoursSinceJanuary1(jan1)
	if want != got {
		t.Errorf("unexpected hour after New Year 00:00..00:59 want:%d, got:%d\n", want, got)
	}
	jan1, _ = time.Parse(timeLayout, "2024-01-01 00:59 UTC")
	got = wholeHoursSinceJanuary1(jan1)
	if want != got {
		t.Errorf("unexpected hour after New Year 00:00..00:59 want:%d, got:%d\n", want, got)
	}
	want = 2
	jan1, _ = time.Parse(timeLayout, "2024-01-01 01:00 UTC")
	got = wholeHoursSinceJanuary1(jan1)
	if want != got {
		t.Errorf("unexpected hour after New Year 00:00..00:59 want:%d, got:%d\n", want, got)
	}
}
func TestMoonInfos(t *testing.T) {
	jsonData := ` [
        {
		"time":"01 Jan 2024 00:00 UT", "phase":78.03, "age":19.019, "diameter":1771.3, "distance":404634,
		"j2000":{"ra":10.5867, "dec":12.7508},
		"subsolar":{"lon":-55.867, "lat":-1.554},
		"subearth":{"lon":0.041, "lat":-4.685},
		"posangle":20.699}
	   ]
	   `
	var arr TypeMoonInfos
	err := json.Unmarshal([]byte(jsonData), &arr)
	want := 1
	got := len(arr)
	if want != got {
		t.Errorf("unexpected length TestMoonInfos - want:%d, got:%d\n%s\n", want, got, err.Error())
	}
	wantF := float32(20.699)
	gotF := arr[0].Posangle
	if wantF != gotF {
		t.Errorf("unexpected TestMoonInfo.Posangle - want:%.4f, got:%.4f\n", wantF, gotF)
	}
}
