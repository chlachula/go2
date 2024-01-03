package moonView

import (
	"fmt"
	"net/http"
	"time"
)

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
func getImgUrl(t time.Time) string {
	nn00, nnnn := svsMagicNumbers(t.Year())
	//example := "https://svs.gsfc.nasa.gov/vis/a000000/a005000/a005048/frames/730x730_1x1_30p/moon.8456.jpg"
	frames := fmt.Sprintf("https://svs.gsfc.nasa.gov/vis/a000000/a00%d/a00%d/frames/", nn00, nnnn)
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
func EventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, part1)
	fmt.Fprint(w, part2)
	// ?date=2023-12-25&utc_hour=4&grid=on&showinfo=on
	getParams := fmt.Sprintf("GET params were: %s", r.URL.Query())
	//t := time.Date(2023, time.December, 22, 20, 0, 0, 0, time.UTC)
	t := getTime(r)
	imgURL := getImgUrl(t)
	println(imgURL)
	// 2023-12-18 03:00 UTC, JD:2460296.625, 8428 hours since 2023-1-1
	fmt.Fprintf(w, part3, getParams, imgURL, imgURL)
	fmt.Fprint(w, part4)
}
