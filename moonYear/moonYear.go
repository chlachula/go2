package moonYear

import (
	"fmt"
	"math"
	"os"
	"time"
)

var SynodicMoon = 29.530589
var satColor = "#eec0c0"
var sunColor = "#ef6c6c"

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
	s := "<tr><td>Month</td>"
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
	if moonAngle < 2.4*d || moonAngle > 9.5*d {
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
			if isFirstQuaterFriday(date, moonAngle) {
				class = "highlight"
				day = date.Format("<a href=\"https://www.cleardarksky.com/cgi-bin/sunmoondata.py?id=RchstrMN&year=2006&month=1&day=2&&tz=-6.0&lat=None&long=None\" target=\"_blank\">2</a>")
			} else if isSecondTuesdayMonth(date) {
				class = "secondTue"
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
func CreateWebpageWithTable(y int, moonAgeDaysJanuary1st float64) {
	pageFormat := `<html>
<head>
 <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
 <title>%s</title>
 <style>
  td {border:2px none solid;}
  .highlight {border:2px blue solid;}
  .secondTue {border:2px purple solid;}
  .example_background {background-image: url(PATH/TO/IMAGE.JPG); }
</style>
</head>
<body>
 %s
<body>
</html>
`
	title := fmt.Sprintf("%d moon phases 4 weeks calendar", y)
	table := createTable(y, moonAgeDaysJanuary1st)
	s := fmt.Sprintf(pageFormat, title, table)
	createFile(fmt.Sprintf("moonYear%d.htm", y), s)
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
