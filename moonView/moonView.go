package moonView

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type (
	EquatorialCoordinates struct {
		RA  float32 `json:"ra"`
		Dec float32 `json:"dec"`
	}
	GeographicCoordinates struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	}
	TypeMoonInfo struct {
		Time     string                `json:"time"`
		Phase    float32               `json:"phase"`
		Age      float32               `json:"age"`
		Diameter float32               `json:"diameter"`
		Distance float32               `json:"distance"`
		J2000    EquatorialCoordinates `json:"j2000"`
		SubSolar GeographicCoordinates `json:"subsolar"`
		SubEarth GeographicCoordinates `json:"subearth"`
		Posangle float32               `json:"posangle"`
	}
)

/*
	{
	 "time":"01 Jan 2024 00:00 UT", "phase":78.03, "age":19.019, "diameter":1771.3, "distance":404634,
	 "j2000":{"ra":10.5867, "dec":12.7508},
	 "subsolar":{"lon":-55.867, "lat":-1.554},
	 "subearth":{"lon":0.041, "lat":-4.685},
	 "posangle":20.699
	},
*/
func svsMagicNumbers(y int) (int, int) {
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

	*/
	svsMagic := []int{
		3810, //2011
		3894, //2012
		4000, //2013
		4118, //2014
		4236, //2015
		4404, //2016
		4537, //2017
		4604, //2018
		4442, //2019
		4768, //2020
		4874, //2021
		4955, //2022
		5048, //2023
		5187, //2024
	}
	y1 := 2011
	i := len(svsMagic) - 1
	y2 := y1 + i

	if y >= y1 && y <= y2 {
		i = y - y1
	}
	nnnn := svsMagic[i]
	nn := nnnn / 100
	nn00 := nn * 100
	return nn00, nnnn
}
func wholeHoursSinceJanuary1(t time.Time) int {
	jan1 := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	tdiff := t.Sub(jan1)
	return int(tdiff.Hours())
}
func svsFrames(t time.Time) string {
	nn00, nnnn := svsMagicNumbers(t.Year())
	return fmt.Sprintf("https://svs.gsfc.nasa.gov/vis/a000000/a00%d/a00%d/frames/", nn00, nnnn)
}
func getImgUrl(t time.Time) string {
	frames := svsFrames(t)
	h := wholeHoursSinceJanuary1(t)
	hhhh := fmt.Sprintf("%04d", h)
	return fmt.Sprintf("%s730x730_1x1_30p/moon.%s.jpg", frames, hhhh)
}
func getTime(r *http.Request) time.Time {
	timeForm := "2006-01-02T15 MST"
	dateStr := r.URL.Query().Get("date")
	hourStr := r.URL.Query().Get("utc_hour")
	if len(hourStr) < 2 {
		hourStr = "0" + hourStr
	}
	timeStr := dateStr + "T" + hourStr + " UTC"
	t, err := time.Parse(timeForm, timeStr)
	if err != nil {
		t = time.Now()
	}
	return t
}

/*
Julian 5 October 1582 = Gregorian 15 October 1582 = JDN 2299161
JDN at 12:00 UT YYYY-MM-DD
*/
func gregorianNoonToJulianDayNumber(Y int, M int, D int) int {
	JDN := (1461*(Y+4800+(M-14)/12))/4 + (367*(M-2-12*((M-14)/12)))/12 - (3*((Y+4900+(M-14)/12)/100))/4 + D - 32075
	return JDN
}
func julianNoonToJulianDayNumber(Y int, M int, D int) int {
	//DN = 367 × Y − (7 × (Y + 5001 + (M − 9)/7))/4 + (275 × M)/9 + D + 1729777
	JDN := 367*Y - (7*(Y+5001+(M-9)/7))/4 + (275*M)/9 + D + 1729777
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
func timeInfo(t time.Time) string {
	// 2023-12-18 03:00 UTC, JD:2460296.625, 8428 hours since 2023-1-1
	t1 := t.Format("2006-01-02 15:04 MST")
	t1UTC := t.UTC().Format("15:04 MST")
	j1 := timeToJulianDay(t)
	h1 := wholeHoursSinceJanuary1(t)
	return fmt.Sprintf("%s, %s, JD:%.4f, %d hours since %d-1-1", t1, t1UTC, j1, h1, t.Year())
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	// ?date=2023-12-25&utc_hour=4&grid=on&showinfo=on
	getParams := fmt.Sprintf("GET params were: %s", r.URL.Query())
	t := getTime(r)
	dHHHHd := fmt.Sprintf(".%04d.", wholeHoursSinceJanuary1(t))
	template1 := part1 + part2 + part3 + part_moon_hour_resources + part4

	type TypeData = struct{ YYYY, SVSframes, Hours, TimeInfo, Radius, GetParams string }
	data := TypeData{t.Format("2006"), svsFrames(t), dHHHHd, timeInfo(t), "330", getParams}

	webpage := "webpage1"
	if html1, err := template.New(webpage).Parse(template1); err != nil {
		fmt.Fprint(w, "Error parsing template "+webpage+": "+err.Error())
	} else {
		if err = html1.Execute(w, data); err != nil {
			fmt.Fprint(w, "Error executing html "+webpage+": "+err.Error())
		}
	}

}
