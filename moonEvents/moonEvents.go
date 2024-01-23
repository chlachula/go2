package moonEvents

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type LocationInfo = struct {
	Index      int
	N          string
	Name       string
	CB         string
	CF         string
	Address    string
	Directions string
}

const (
	WatsonFields      = 0
	A125Live          = 1
	SalemGlenn        = 2
	OxbowPark         = 3
	RootRiverPark     = 4
	Frontenac         = 5
	Keller            = 6
	EagleBluff        = 7
	Forestville       = 8
	Elgin             = 9
	AustinSolaFideObs = 10
	AustinSolaFide2   = 11
	DodgeCenterPL     = 12
	RedWing           = 13
)

var myLocations []LocationInfo

func locationByIndex(index int) LocationInfo {
	location := myLocations[index]
	return location
}
func locationByN(N string) LocationInfo {
	var unknownLocation LocationInfo
	for _, loc := range myLocations {
		if loc.N == N {
			return loc
		}
	}
	fmt.Printf("Unknown location ID N: %s\n", N)
	unknownLocation.Address = "Unknown Address"
	unknownLocation.Name = "Unknown '" + N + "'"
	unknownLocation.CF = "red"
	unknownLocation.CB = "lightgray"
	unknownLocation.N = N
	return unknownLocation
}

var timeLocation *time.Location = time.Local

const tpl1 = `
Public Sky Observing of the {{.DarkMoon}} {{.Planets}} bright stars, double stars, and more by telescope with guidance members of the Rochester Astronomy Club.
<p>What to expect at Public Sky Observing – <a href="https://rochesterskies.org/what-to-expect-at-rac-night-sky-viewing-events/" target="_blank" rel="noopener">click here</a>.</p>
<p style="color:blue">Reload this page before you leave for the event. Event can be cancelled in case of cloudy weather or another reason.</p>

{{.DriveDirections}}

{{.MoonPicture}}
<br/>
<p><strong>Resources:</strong>
<a href="https://in-the-sky.org/newscal.php?month={{.MM}}&year={{.YYYY}}&maxdiff=7#datesel" target="_blank" rel="noopener noreferrer">In the Sky {{.MONTH}}/{{.YYYY}}</a>,
<a href="http://us.cal3.net/moon-phase-calendar-{{.YYYY}}" target="_blank" rel="noopener noreferrer">Phases of the Moon {{.YYYY}}</a>,
<a href="https://www.almanac.com/astronomy/planets-rise-and-set/MN/Rochester/{{.YYYY}}-{{.MM}}-{{.DD}}" target="_blank" rel="noopener noreferrer">Planets rise and set {{.MM}}/{{.DD}}</a>, 
<a href="https://www.cleardarksky.com/cgi-bin/sunmoondata.py?id=RchstrMN&year={{.YYYY}}&month={{.MM}}&day={{.DD}}&&tz=-6.0&lat=None&long=None" target="_blank" rel="noopener noreferrer">Sun and Moon {{.MM}}/{{.DD}}</a>, 
<a href="https://www.suntoday.org/" target="_blank" rel="noopener">Sun today</a>, 
<a href="https://www.timeanddate.com/calendar/?year={{.YYYY}}&country=1" target="_blank" rel="noopener noreferrer">Holidays {{.YYYY}}</a>
<br/>
<a href="https://www.planetary.org/articles/moon-features-you-can-see-from-earth" target="_blank" rel="noopener">Moon Features You Can See From Earth</a>,
<br/>
Messier <a href="https://www.astroleague.org/al/obsclubs/messier/messlist.html" target="_blank" rel="noopener">list</a>, 
Caldwell <a href="https://www.astroleague.org/al/obsclubs/caldwell/cldwlist.html" target="_blank" rel="noopener">list</a>, 
Double stars <a href="https://rochesterskies.org/astronomy-league-double-stars-list/" target="_blank" rel="noopener">list</a>,
Carbon stars <a href="https://rochesterskies.org/astronomy-league-carbon-stars-list/" target="_blank" rel="noopener">list</a> </p>
`

const tpl2head = `
<a name="{{.YYYY}}plan"></a>
<table bgcolor="gray">
<caption>{{.YYYY}} Public Sky Observing Plan </caption>
    <tbody>
    <tr bgcolor="white">
        <th align="left">Month</th>
        <th align="left">Day</th>
        <th align="left">DoW</th>
        <th align="left">Location</th>
        <th align="left">Sunset</th>
        <th align="left">Start</th>
        <th align="left">End</th>
        <th align="left">☽ age</th>
        <th align="left">Planets</th>
        <th align="left">Title</th>
    </tr>
    </tbody>`
const tpl2row = `
    <tr bgcolor="%s">
        <td align="left">%s</td>
        <td align="left"><a href="/event/%s-%s/">%02d</a></td>
        <td >%s</td>
        <td bgcolor="%s"><font color="%s" title="%s" >%s</font></td>
        <td align="left">%s</td>
        <td align="left">%s</td>
        <td align="left">%s</td>
        <td align="left">%s</td>
        <td align="left">%s</td>
        <td align="left">%s</td>
    </tr>`
const tpl3 = `
<p>
  ☿ - Mercury, ♀ - Venus, ♂ - Mars, ♃ - Jupiter, ♄ - Saturn, ⛢ - Uranus, ♆ - Neptune
</p>
<table bgcolor="gray">
 <caption>Color code</caption>
 <tr><td 
bgcolor="brown"><font color="white">Oxbow Park</font></td><td 
bgcolor="black"><font color="white">Root River Park</font></td><td 
bgcolor="lightgray">Soccer Fields</td><td 
bgcolor="lightgreen">Out of Rochester</td><td 
bgcolor="#FF8">Moon visible</td></tr>
</table>
<p><a href="https://www.facebook.com/rochesterastronomy" target="_blank" rel="noopener">Facebook page</a></p>

`

type EventType = struct {
	MM      int     `json:"mm"`
	DD      int     `json:"dd"`
	T1h     int     `json:"t1h"`
	T1m     int     `json:"t1m"`
	T1len   int     `json:"t1len"`
	Loc     int     `json:"loc"`
	LocN    string  `json:"N"`
	SunSet  string  `json:"sunset"`
	MDage   float32 `json:"age"`
	Planets string  `json:"planets"`
	Title   string  `json:"title"`
}
type EventDetailedDataType = struct {
	YYYY            string
	MM              string
	MONTH           string
	DD              string
	HH_mm1          string
	HH_mm2          string
	DarkMoon        string
	Planets         string
	DriveDirections string
	MoonPicture     string
}
type EventsYearType = struct {
	YYYY     int `json:"yyyy"`
	MMprev   int `json:"mmPrev"`
	DDprev   int `json:"ddPrev"`
	MMnext   int `json:"mmNext"`
	DDnext   int `json:"ddNext"`
	MagicSVS int `json:"magicSVS"`
	Event    []EventType
}
type ByDate []EventType

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Less(i, j int) bool { return a[i].MM*100+a[i].DD < a[j].MM*100+a[j].DD }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// https://regex101.com/r/cjvtzS/1
var regex1 = regexp.MustCompile(`(?m)\/event\/(public-sky-observing|members-night|star-bq)-(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)

const (
	pre_public                = "public-sky-observing"
	pre_members               = "members-night"         // @
	pre_starbq                = "members-night-star-bq" // !
	ft                        = "1/2 15:04 MST"
	DriveDirections2OxbowPark = ""
	saturdayColor1            = "#ff8"
)

var EventsData EventsYearType

func LoadData(locationsJson string, eventsJson string) {
	var err error
	myLocations, err = LoadLocations(locationsJson)
	if err != nil {
		panic(err)
	}
	EventsData, err = LoadEventsYear(eventsJson)
	if err != nil {
		panic(err)
	}
}
func stellariumObject(b string, t *time.Time, text string) string {
	//f := "https://stellarium-web.org/skysource/Jupiter?fov=0.3&date=2024-01-20T00%3A30%3A00Z&lat=44&lng=-92.5&elev=0"
	f1 := "https://stellarium-web.org/skysource/%s?fov=%s&date=%s&lat=44&lng=-92.5&elev=0"
	f2 := `<a href="%s" title="%s" target="_blank">%s</a>`
	fov := "0.3"
	body := b
	switch b {
	case "☿":
		body = "Mercury"
	case "♀":
		body = "Venus"
	case "♂":
		body = "Mars"
	case "♃":
		body = "Jupiter"
	case "♄":
		body = "Saturn"
	case "⛢":
		body = "Uranus"
	case "♆":
		body = "Neptune"
	}
	if body == "Moon" {
		fov = "1"
	} else if body == "Sun" {
		fov = "30"
	}
	link := fmt.Sprintf(f1, body, fov, t.In(time.UTC).Format(time.RFC3339))
	title := body + " @ " + t.Format(ft) + " = " + t.In(time.UTC).Format(ft)
	return fmt.Sprintf(f2, link, title, text)
}
func makeFirstDataUrl() string {
	return fmt.Sprintf("/event/public-sky-observing-%d-%02d-%02d", EventsData.YYYY, EventsData.Event[0].MM, EventsData.Event[0].DD)
}

func firstDataLink() string {
	s := makeFirstDataUrl()
	return fmt.Sprintf("<a href=\"%s\">%s</a>", s, s)
}

func getUrlParamsMap(compRegEx *regexp.Regexp, url string) (paramsMap map[string]string) {
	//var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
func planetLink(name string, fov float32, t time.Time) string {
	dateTimeStr := t.Format("2006-01-02T15:04:05Z")
	sFormat := "https://stellarium-web.org/skysource/%s?fov=%.1f&date=%s&lat=44&lng=-92.5&elev=0"
	// https://stellarium-web.org/skysource/Jupiter?fov=0.2&date=2023-08-23T03Z&lat=44&lng=-92.5&elev=0
	url := fmt.Sprintf(sFormat, name, fov, dateTimeStr)
	return "<a href=\"" + url + "\" target=\"_blank\">" + name + "</a>"
}
func planetsFullNames(p string, t time.Time) string {
	s := strings.Replace(p, "☿", planetLink("Mercury", 0.1, t), 1)
	s = strings.Replace(s, "♀", planetLink("Venus", 0.1, t), 1)
	s = strings.Replace(s, "♂", planetLink("Mars", 0.1, t), 1)
	s = strings.Replace(s, "♃", planetLink("Jupiter", 0.3, t), 1)
	s = strings.Replace(s, "♄", planetLink("Saturn", 0.2, t), 1)
	s = strings.Replace(s, "⛢", planetLink("Uranus", 0.05, t), 1)
	s = strings.Replace(s, "♆", planetLink("Neptune", 0.05, t), 1)
	return "planets " + s
}
func svsMagicNumbers(y int) (int, int) {
	type ab struct {
		a int
		b int
	}
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
	aa := []ab{
		{a: 3800, b: 3810}, //2011
		{a: 3800, b: 3894}, //2012
		{a: 4000, b: 4000}, //2013
		{a: 4100, b: 4118}, //2014
		{a: 4200, b: 4236}, //2015
		{a: 4400, b: 4404}, //2016
		{a: 4500, b: 4537}, //2017
		{a: 4600, b: 4604}, //2018
		{a: 4400, b: 4442}, //2019
		{a: 4700, b: 4768}, //2020
		{a: 4800, b: 4874}, //2021
		{a: 4900, b: 4955}, //2022
		{a: 5000, b: 5048}, //2023
		{a: 5100, b: 5187}, //2024
	}
	y1 := 2011
	i := len(aa) - 1
	y2 := y1 + i

	if y >= y1 && y <= y2 {
		i = y - y1
	}
	return aa[i].a, aa[i].b
}
func makeMoonPicture(t time.Time) string {
	nn00, nnnn := svsMagicNumbers(t.Year())
	yrday := t.YearDay()
	jpg := (yrday-1)*24 + t.Hour() + 6 + 1
	title := t.Format(ft) + " = " + t.In(time.UTC).Format(ft)
	moonPic := `Moon on %s evening <br/>
<img src="https://svs.gsfc.nasa.gov/vis/a000000/a00%d/a00%d/frames/730x730_1x1_30p/moon.%04d.jpg" title="moon.%04d.jpg %s"/>
<br/>
Credit: NASA <a href="https://svs.gsfc.nasa.gov/cgi-bin/details.cgi?aid=%d" rel="noopener noreferrer" target="_blank">Scientific Visualization Studio</a>
`
	return fmt.Sprintf(moonPic, t.Format("Monday January 2"), nn00, nnnn, jpg, jpg, title, nnnn)
}
func getTimePrevAndNext(iEv int, events EventsYearType) (time.Time, time.Time) {
	var ev EventType
	yyyyPrev := events.YYYY
	if iEv > 0 {
		ev = events.Event[iEv-1]
	} else {
		yyyyPrev -= 1
		ev = EventType{MM: events.MMprev, DD: events.DDprev, T1h: 18, T1m: 00, T1len: 90, Loc: OxbowPark, LocN: "Oxb", SunSet: "16:00", MDage: 0.0, Planets: ""}
	}
	tPrev := time.Date(yyyyPrev, time.Month(ev.MM), ev.DD, ev.T1h, ev.T1m, 0, 0, timeLocation)

	yyyyNext := events.YYYY
	if iEv < len(events.Event)-1 {
		ev = events.Event[iEv+1]
	} else {
		yyyyNext += 1
		ev = EventType{MM: events.MMnext, DD: events.DDnext, T1h: 18, T1m: 00, T1len: 90, Loc: OxbowPark, LocN: "Oxb", SunSet: "16:00", MDage: 0.0, Planets: ""}
	}
	tNext := time.Date(yyyyNext, time.Month(ev.MM), ev.DD, ev.T1h, ev.T1m, 0, 0, timeLocation)

	return tPrev, tNext
}
func createSunLink(YYYY int, ev EventType) string {
	hh, _ := strconv.Atoi(ev.SunSet[:2])
	mm, _ := strconv.Atoi(ev.SunSet[3:])
	t := time.Date(YYYY, time.Month(ev.MM), ev.DD, hh, mm, 0, 0, timeLocation)
	return stellariumObject("Sun", &t, ev.SunSet)
}
func createPlanetLinks(planetsString string, t time.Time) string {
	s := ""
	for _, c := range planetsString {
		cStr := string(c)
		if c > 255 {
			s += stellariumObject(cStr, &t, cStr)
		} else {
			s += cStr
		}
	}
	return s
}
func makeEventTable(iEv int, events EventsYearType) string {
	s := tpl2head
	for i, ev := range events.Event {
		//Mon Jan 2 15:04:05 -0700 MST 2006
		t1 := time.Date(events.YYYY, time.Month(ev.MM), ev.DD, ev.T1h, ev.T1m, 0, 0, timeLocation)
		rowBgColor := "white"
		if ev.MDage < 14.0 {
			rowBgColor = saturdayColor1
		}
		if strings.HasPrefix(ev.Title, "@") || strings.HasPrefix(ev.Title, "!") {
			rowBgColor = "lightblue"
		}
		if strings.HasPrefix(ev.Title, "@") || strings.HasPrefix(ev.Title, "!") {
			rowBgColor = "lightblue"
		}
		if i == iEv {
			rowBgColor = "pink"
		}
		location := locationByN(ev.LocN)
		locBgColor := location.CB
		fontColor := location.CF
		nameAddr := location.Name
		addr := location.Address
		eventDate := t1.Format("2006-01-02")
		eventPrefix := pre_public
		if strings.HasPrefix(ev.Title, "@") {
			eventPrefix = pre_members
		}
		if strings.HasPrefix(ev.Title, "!") {
			eventPrefix = pre_starbq
			eventDate = t1.Format("2006")
		}
		sunLink := createSunLink(events.YYYY, ev)
		moonLink := stellariumObject("Moon", &t1, fmt.Sprintf("%.1f", ev.MDage))
		planetLinks := createPlanetLinks(ev.Planets, t1)
		s += fmt.Sprintf(tpl2row, rowBgColor, t1.Format("Jan"), eventPrefix, eventDate, ev.DD, t1.Format("Mon"),
			locBgColor, fontColor, ev.LocN+": "+addr, nameAddr, sunLink, t1.Format("3:04 PM"), t1.Add(time.Minute*time.Duration(ev.T1len)).Format("3:04 PM"),
			moonLink, planetLinks, ev.Title)
	}
	tPrev, tNext := getTimePrevAndNext(iEv, events)
	s += fmt.Sprintf(`<tr bgcolor="white">
	<td>event</td>
	<td colspan="4" align="right"    ><a href="/event/public-sky-observing-%s/">%s</a> previous</td>
	<td colspan="5" align="left">next <a href="/event/public-sky-observing-%s/">%s</a> </td>
	</tr>`, tPrev.Format("2006-01-02"), tPrev.Format("January 02"), tNext.Format("2006-01-02"), tNext.Format("January 02"))

	s += "\n</table>\n"
	return s
}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func GenerateGoogleCalendarListCSV(events EventsYearType, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	w.WriteString("Subject,Start Date,Start Time,End Date,End Time,All day event,Description,Location\n")
	for _, ev := range events.Event {
		//Mon Jan 2 15:04:05 -0700 MST 2006
		t1 := time.Date(events.YYYY, time.Month(ev.MM), ev.DD, ev.T1h, ev.T1m, 0, 0, timeLocation)
		subject := "jch:Public Sky Observing"
		startDate := t1.Format("01/02/2006")
		startTime := t1.Format("3:04 PM")
		endDate := startDate
		endTime := t1.Add(time.Minute * time.Duration(ev.T1len)).Format("3:04 PM")
		allDayEvent := "FALSE"
		description := "Public Sky Observing celestial wonders with guidance members of the Rochester Astronomy Club."
		location := locationByN(ev.LocN)
		address := location.Address
		line := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,\n", subject, startDate, startTime, endDate, endTime, allDayEvent, description, address)
		w.WriteString(line)
	}
	w.Flush()
	f.Close()
}
func makeDarkMoon(MDage float32) string {
	s := ""
	if MDage < 15.0 {
		s = fmt.Sprintf("young Moon – %.1f days after New Moon,", MDage)
	}
	return s
}
func EventHandler(w http.ResponseWriter, r *http.Request) {
	m := getUrlParamsMap(regex1, r.URL.Path)
	yyyyStr := m["Year"]
	mmStr := m["Month"]
	ddStr := m["Day"]
	iEv := -1
	if yyyyStr == "" || mmStr == "" || ddStr == "" {
		iEv = -2
	}

	nYYYY, _ := strconv.Atoi(yyyyStr)
	nMM, _ := strconv.Atoi(mmStr)
	nDD, _ := strconv.Atoi(ddStr)

	if iEv > -2 {
		for i, evI := range EventsData.Event {
			if evI.MM == nMM && evI.DD == nDD {
				iEv = i
				break
			}
		}
	}
	title := "PSO-RAC"
	urlOK := "URL OK "
	if iEv < 0 || EventsData.YYYY != nYYYY {
		title = "Error"
		urlOK = "No data for URL "
	}

	fmt.Fprintf(w, `<html><head><title>%s</title>
	<meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\">
	<link rel="icon" type="image/ico" href="fsvicon.ico">
	</head>
	<body>
	<a href="/">home</a> <a href="/api">/api</a>
	<span>%s`+r.URL.Path+"</span>", title, urlOK)
	if iEv < 0 {
		//		iEv = 0
		fmt.Fprint(w, firstDataLink())
		return
	}

	ev := EventsData.Event[iEv]
	t1 := time.Date(EventsData.YYYY, time.Month(ev.MM), ev.DD, ev.T1h, ev.T1m, 0, 0, timeLocation)
	fTitle := "Public Sky Observing 2006-01/02 3:04PM "
	if strings.HasPrefix(ev.Title, "@") {
		fTitle = "Member's Night 2006-01/02 3:04PM "
	}
	evTitle := t1.Format(fTitle)
	loc := locationByN(ev.LocN)
	fmt.Fprintf(w, "\n%s at %s\n<hr/>\n\n\n", evTitle, loc.Name+": "+loc.Address)

	moonPic := ""
	if ev.MDage < 15.0 {
		moonPic = makeMoonPicture(t1)
	}

	data := EventDetailedDataType{
		YYYY:            t1.Format("2006"),
		MM:              t1.Format("01"),
		MONTH:           t1.Format("January"),
		DD:              t1.Format("02"),
		HH_mm1:          t1.Format("3:04PM"),
		HH_mm2:          t1.Add(time.Minute * time.Duration(ev.T1len)).Format("3:04PM"),
		DarkMoon:        makeDarkMoon(ev.MDage),
		Planets:         planetsFullNames(ev.Planets, t1.Add(6*time.Hour)),
		DriveDirections: "",
		MoonPicture:     moonPic,
	}
	/*	if ev.Loc == OxbowPark {
		data.DriveDirections = DriveDirections2OxbowPark
	} */
	tpl := strings.ReplaceAll(tpl1, "\n", " ") + makeEventTable(iEv, EventsData) + tpl3
	t, err := template.New("webpage1").Parse(tpl)
	check(err)
	err = t.Execute(w, data)
	check(err)
}
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	data := &EventsData
	b, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}
