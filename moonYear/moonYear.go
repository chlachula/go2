package moonYear

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var SynodicMoon = 29.530589
var satColor = "#eec0c0"
var sunColor = "#ef6c6c"
var events = make([]EventRecord, 0)

type EventRecord struct {
	Color        string
	URL_prefix   string
	Date_time    time.Time
	Duration_min int
	Name         string
	Location     string
}

func monthColor(date time.Time) string {
	s := "lightgray"
	if int(date.Month())%2 == 0 {
		s = "lightblue"
	}
	return s
}
func getMoonIcon(moonAngle float64, date time.Time) string {
	moonAngle += 90.0 / (2.0 * 7.0)
	moonAngle = to0_360(moonAngle)
	n := int(moonAngle / (90.0 / 7.0))
	moonAngle -= 90.0 / (2.0 * 7.0)
	moonAngle = to0_360(moonAngle)
	age := moonAngle / 360. * SynodicMoon
	imgFormat := `<img src="../moonPhases/moon28f%02d.svg" width="30" title="%.1fd"/>`
	moonSunData := date.Format("<a href=\"https://www.cleardarksky.com/cgi-bin/sunmoondata.py?id=RchstrMN&year=2006&month=1&day=2&&tz=-6.0&lat=None&long=None\" target=\"_blank\">")

	icon := fmt.Sprintf(imgFormat, n, age)
	return moonSunData + icon + "</a>"
}
func rowMonthNameCell(date time.Time) string {
	d2 := date.Add(13 * 24 * time.Hour)
	bgcolor := monthColor(d2)
	monthName := d2.Format("January")
	return fmt.Sprintf(`<td bgcolor="%s">%s</td>`, bgcolor, monthName)
}
func nameRow() string {
	s := "<tr><td>Month 1st</td>"
	week := "<td>Mon</td><td>Tue</td><td>Wed</td><td>Thu</td><td>Fri</td><td>Sat</td><td>Sun</td>"
	s += week + week + week + week
	return s + "</tr>\n"
}
func to0_360(x float64) float64 {
	x360 := math.Remainder(x, 360.0)
	if x360 < 0.0 {
		x360 += 360.0
	}
	return x360
}
func isFirstQuaterFriday(date time.Time, moonAngle float64) bool {
	if date.Weekday() != time.Friday {
		return false
	}
	moonAngle = to0_360(moonAngle)
	d := 360.0 / SynodicMoon
	min := 2.4 * d //  29.26 degree
	max := 9.5 * d // 115.81 degree
	if moonAngle < min || moonAngle > max {
		return false
	}
	return true
}
func isSecondTuesdayMonth(date time.Time) bool {
	if date.Weekday() != time.Tuesday {
		return false
	}
	if date.Day() < 8 || date.Day() > 14 {
		return false
	}
	return true
}
func isCSVfileEventDay(date time.Time) bool {
	for _, e := range events {
		if e.Date_time.Equal(date) {
			return true
		}
	}
	return false
}
func CSVfileEventDay(date time.Time) (string, string) {
	dd := date.Format("02")
	for _, e := range events {
		if e.Date_time.Equal(date) {
			ymd := date.Format("-2006-01-02")
			return fmt.Sprintf("<a href=\"%s\" target=\"_blank\" title=\"%dmin @%s:%s\">%s</a>", e.URL_prefix+ymd, e.Duration_min, e.Location, e.Name, dd),
				"high_" + e.Color
		}
	}
	return dd, ""
}
func column(index int, cols []string, prev string) string {
	if index < len(cols) {
		col := strings.Trim(cols[index], " ")
		if len(col) > 0 {
			return col
		} else {
			return prev
		}
	} else {
		return prev
	}
}
func columsToEventRecord(cols []string, prevEvent EventRecord) EventRecord {
	e := prevEvent
	e.Color = column(0, cols, prevEvent.Color)
	e.URL_prefix = column(1, cols, prevEvent.URL_prefix)
	s := cols[2]
	t1, err := time.Parse("2006-01-02", s)
	if err == nil {
		e.Date_time = t1
	} else {
		e.Date_time = prevEvent.Date_time
	}
	s = cols[3]
	var i1 int
	i1, err = strconv.Atoi(s)
	if err == nil {
		e.Duration_min = i1
	} else {
		e.Duration_min = prevEvent.Duration_min
	}
	e.Name = column(4, cols, prevEvent.Name)
	e.Location = column(5, cols, prevEvent.Location)
	return e
}
func CsvFileToEvents(fname string) []EventRecord {
	events := make([]EventRecord, 0)
	bytes, err := os.ReadFile(fname) //Read entire file content. No need to close the file
	if err != nil {
		fmt.Println(err.Error())
		return events
	}
	text := string(bytes)
	lines := strings.Split(text, "\n")
	var prevEvent EventRecord
	// skip the first line with columns names
	for i := 1; i < len(lines); i++ {
		cols := strings.Split(lines[i], ",")
		if len(lines[i]) > 4 { // could be 0, it's unlikely to have 4 invisible chars on line
			event := columsToEventRecord(cols, prevEvent)
			events = append(events, event)
			prevEvent = event
			///fmt.Printf("--- %s\n%q\n\n", lines[i], event)
		}
	}
	return events
}
func createEventsList() string {
	s := ""
	if len(events) > 0 {
		for _, e := range events {
			dow := e.Date_time.Format("Mon")
			ymd := e.Date_time.Format("2006-01-02")
			color := "black"
			switch e.Color {
			case "r":
				color = "red"
			case "b":
				color = "darkblue"
			case "m":
				color = "magenta"
			}
			s += fmt.Sprintf("<code>%s %s</code>: <a href=\"%s-%s\">%s</a>, <span style=\"color:%s;\">%s</span><br/>\n", dow, ymd, e.URL_prefix, ymd, e.Name, color, e.Location)
		}
		return s + "\n<hr/>\n"
	}
	return s
}

func createTable(y int, moonAgeDaysJanuary1st float64) string {
	moonAngle := 360.0 * moonAgeDaysJanuary1st / SynodicMoon
	date := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
	for date.Weekday() != time.Monday {
		date = date.Add(-24 * time.Hour)
		moonAngle -= 360.0 / SynodicMoon
	}
	moonAngle = to0_360(moonAngle)
	s := fmt.Sprintf("<table><caption><a href=\"moonYear%d.htm\">&lt;</a> - %d - <a href=\"moonYear%d.htm\">&gt;</a></caption>\n", y-1, y, y+1)
	s += nameRow()
	for row := 0; row <= 13; row++ {
		s += fmt.Sprintf("\n<tr>" + rowMonthNameCell(date))
		for c := 0; c < 28; c++ {
			bgcolor := monthColor(date)
			if date.Weekday() == time.Saturday {
				bgcolor = satColor
			}
			if date.Weekday() == time.Sunday {
				bgcolor = sunColor
			}
			day := date.Format("2")
			class := ""
			if len(events) == 0 {
				if isFirstQuaterFriday(date, moonAngle) {
					class = "highlight"
					day = date.Format("<a href=\"https://www.cleardarksky.com/cgi-bin/sunmoondata.py?id=RchstrMN&year=2006&month=1&day=2&&tz=-6.0&lat=None&long=None\" target=\"_blank\">2</a>")
				} else if isSecondTuesdayMonth(date) {
					class = "secondTue"
				}

			} else {
				if isCSVfileEventDay(date) {
					day, class = CSVfileEventDay(date)
				}
			}
			s += fmt.Sprintf("\n<td align=\"center\" bgcolor=\"%s\" class=\"%s\">%s<br/>%s</td>", bgcolor, class, getMoonIcon(moonAngle, date), day)
			date = date.Add(24 * time.Hour)
			moonAngle += 360.0 / SynodicMoon
		}
		s += "</tr>"
	}
	s += "</table>\n\n"
	return s
}
func createTime() string {
	now := time.Now()
	return now.Format("Mon 2006-01-02_15:04:05")
}
func CreateWebpageWithTable(y int, moonAgeDaysJanuary1st float64, csvFileName string, webPageDir string) {
	pageFormat := `<html>
<head>
 <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
 <title>%s</title>
 <style>
  td {border:2px none solid;}
  .highlight {border:2px blue solid;}
  .high_r {background-image: url("bg_image_r.svg");}
  .high_m {background-image: url("bg_image_m.svg");}
  .high_b {background-image: url("bg_image_b.svg");}
  .secondTue {border:2px purple solid;}
  .bg_image {background-image: url("bg_image1.svg"); }
</style>
</head>
<body>
 %s
<body>
</html>
`
	events = CsvFileToEvents(csvFileName)
	title := fmt.Sprintf("%d moon phases 4 weeks calendar", y)
	list := createEventsList()
	table := createTable(y, moonAgeDaysJanuary1st)
	created := "\n<br/><br/> Created on " + createTime()
	s := fmt.Sprintf(pageFormat, title, list+table+created)
	filename := filepath.Join(webPageDir, "moonYear%d.htm")
	createFile(fmt.Sprintf(filename, y), s)
}
func createFile(fname, ftext string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes := []byte(ftext)
	if _, err := f.Write(bytes); err != nil {
		return err
	}
	return nil
}
