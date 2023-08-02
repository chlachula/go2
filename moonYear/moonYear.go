package moonYear

import (
	"fmt"
	"os"
	"time"
)

var SynodicMoon = 29.530589
var satColor = "#eec0c0"
var sunColor = "#ef6c6c"

func CreateTable1(y int) string {
	for d := 1; d < 9; d++ {
		date := time.Date(y, time.January, d, 0, 0, 0, 0, time.UTC)
		fmt.Println(date.Weekday(), int(date.Weekday()))
	}
	return ""
}
func monthColor(date time.Time) string {
	s := "lightgray"
	if int(date.Month())%2 == 0 {
		s = "lightblue"
	}
	return s
}
func getMoonIcon(moonAngle float64) string {
	moonAngle += 90.0 / (2.0 * 7.0)
	for moonAngle > 360.0 {
		moonAngle -= 360.0
	}
	n := int(moonAngle / (90.0 / 7.0))
	moonAngle -= 90.0 / (2.0 * 7.0)
	for moonAngle < 0.0 {
		moonAngle += 360.0
	}
	age := moonAngle / 360. * SynodicMoon
	imgFormat := `<img src="../moonPhases/moon28f%02d.svg" width="30" title="%.1fd"/>`
	return fmt.Sprintf(imgFormat, n, age)
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
func CreateTable(y int, moonAgeDaysJanuary1st float64) {
	moonAngle := 360.0 * moonAgeDaysJanuary1st / SynodicMoon
	date := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)
	for date.Weekday() != time.Monday {
		date = date.Add(-24 * time.Hour)
		moonAngle -= 360.0 / SynodicMoon
	}
	if moonAngle < 0.0 {
		moonAngle += 360.0
	}
	s := fmt.Sprintf("<table><caption>%d</caption>\n", y)
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
			s += fmt.Sprintf("\n<td align=\"center\" bgcolor=\"%s\">%s<br/>%s</td>", bgcolor, getMoonIcon(moonAngle), date.Format("2"))
			date = date.Add(24 * time.Hour)
			moonAngle += 360.0 / SynodicMoon
		}
		s += "</tr>"
	}
	s += "</table>"
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
