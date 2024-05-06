package SkyMapLab

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"
)

type StarRecord struct {
	RA  float64 `json:"RA"`
	De  float64 `json:"De"`
	Mag float64 `json:"Mag"`
}

var SliceOfStars []StarRecord
var magBrightest = -1.5 // Sirius
var magMin = 5.0
var SliceOfConstellations []ConstellationCoordPoints

type EqCoords struct {
	RA float64 `json:"RA"`
	De float64 `json:"De"`
}
type EqPoints []EqCoords
type ConstellationCoordPoints struct {
	Abbr  string
	Lines []EqPoints
}

const (
	htmlEnd      = "\n<br/></body></html>"
	svgTemplate1 = `
<svg xmlns="http://www.w3.org/2000/svg" 
    xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="-250 -250 500 500">
    <title>Sky Map</title>
 <defs>
    <style>
	 .font1 { 
		font-size: {{.FontSize}}px;
		font-family: Franklin Gothic, sans-serif;
		font-weight: 90; 		
		letter-spacing: 2px;
	 }
	 .upFont { 
		fill: {{.TopColor}};
	 }
	 .downFont { 
		fill: {{.BottomColor}};
	 }
	 .board {
		stroke:orange;
		stroke-width:0.5
		fill:pink;
	 }
	 .cross {
		stroke:black;
		stroke-width:0.25
		fill:none
	 }
    </style>
  %s
 </defs> 

  <g id="draw">
   <use xlink:href="#plotConstellations" />
   <use xlink:href="#plotStars" />
   <use xlink:href="#dateRoundScale" />
   <use xlink:href="#raHourScale" />
   <use xlink:href="#raCross" />
   </g>

</svg>
`
)

type SvgDataType = struct {
	FontSize    float64
	TopColor    string
	BottomColor string
}

var (
	TopText    string
	BottomText string
)

/*
obliquity = 23.43929111 - 46.8150"t - 0.00059"t^2 + 0.001813*t^3
T = (JD-2451545)/36525 ... centuries since J2000.0
Y:2000 T:0.00 𝜀 = 23.439291
Y:2025 T:0.25 𝜀 = 23.436040
Y:2050 T:0.50 𝜀 = 23.432789
Y:2075 T:0.75 𝜀 = 23.429538
Y:2100 T:1.00 𝜀 = 23.426287
*/
func EclipticObliquity(T float64) float64 {
	𝜀 := 23.43929111 - (1.300416666666666666666666666667e-2+(1.6388888888888888888888888888889e-7-5.0361111111111111111111111111111e-7*T)*T)*T
	return 𝜀
}

/*
cos𝛿*cos𝛼 = cos𝛽*cos𝜆 => cos𝛿 = cos𝛽*cos𝜆 / cos𝛼
cos𝛿*sin𝛼 = cos𝛽*sin𝜆*cos𝜀 − sin𝜀*sin𝛽 = sin𝛼/cos𝛼 * cos𝛽*cos𝜆 = tan𝛼*cos𝛽*cos𝜆

	sin𝛿 = sin𝛽*cos𝜀 + sin𝜀*cos𝛽*sin𝜆
*/
func EclipticalToEquatorial(La, Be float64) (float64, float64) {
	𝜀 := 23.436040 * math.Pi / 180.0 //for year 2025
	sinRAcosDe := math.Cos(Be)*math.Sin(La)*math.Cos(𝜀) - math.Sin(𝜀)*math.Sin(Be)

	RA := math.Atan2(sinRAcosDe, (math.Cos(Be) * math.Cos(La)))
	if RA < 0.0 {
		RA += 2.0 * math.Pi
	}

	sinDe := math.Sin(Be)*math.Cos(𝜀) + math.Sin(𝜀)*math.Cos(Be)*math.Sin(La)
	De := math.Asin(sinDe)

	return RA, De
}

func SetVariables(top, bottom string) {
	TopText = top
	BottomText = bottom
	fmt.Printf("TOP:    %s\nBOTTOM: %s\n", TopText, BottomText)
}

func getSvgData(color bool) SvgDataType {
	data := SvgDataType{
		TopColor:    "green",
		BottomColor: "red",
		FontSize:    8,
	}
	if !color {
		data.TopColor = "black"
		data.BottomColor = "darkgray"
	}
	return data
}

func raCross() string {
	str := `
	<g id="raCross">	  
      <line x1="-154" y1="0" x2="154" y2="0" class="cross" />
	  <line x1="0" y1="-154" x2="0" y2="154" class="cross" />
	  <circle cx="0" cy="0" r="100" stroke="black" stroke-width="0.5" fill="none" />
	</g>
`
	return str
}
func cartesianXY(r, a float64) (float64, float64) {
	x := -r * math.Sin(a)
	y := r * math.Cos(a)
	return x, y
}
func raHourRoundScale() string {
	r1 := 150.0
	r2 := 155.0
	r3 := 162.0

	f0 := `
	<g id="raHourScale">
	  <circle cx="0" cy="0" r="150" stroke="black" stroke-width="0.5" fill="none" />
	  <circle cx="0" cy="0" r="152" stroke="black" stroke-width="0.5" fill="none" />
	  %s
	</g>
`
	s := "\n"
	f1 := "      <line x1=\"%.1f\" y1=\"%.1f\" x2=\"%.1f\" y2=\"%.1f\" class=\"cross\" />\n"
	f2 := ` 
	 <path id="relB" d="M0,0 m-{{.Bx}},{{.By}} a{{.R1}},{{.R1}} 0 0,0  {{.Bx2}},0 " style="fill:none;fill-opacity: 1;stroke:yellow;stroke-width: 10.5"/>
	 <text dy="{{.Dy2}}" dx="{{.Dx2}}" textLength="{{.Blen}}"  class="font1 downFont">
	     <textPath xlink:href="#relB">{{.BottomText}}</textPath>
     </text>
`
	f2 = `      <path id="raHour%d" d="M%.1f,%.1f A162.0,162.0 0 0,0  %.1f,%.1f " style="fill:none;fill-opacity: 1;stroke:green;stroke-width: 0.7"/>
      <text alignment-baseline="baseline" text-anchor="start" class="font1 downFont">
	    <textPath xlink:href="#raHour%d">%d</textPath>
      </text>

`
	aQuaterHour := math.Pi / 48.0
	for ra := 0; ra <= 23; ra++ {
		a := float64(ra*15) * math.Pi / 180.0
		x1, y1 := cartesianXY(r1, a)
		x2, y2 := cartesianXY(r2, a)
		s += fmt.Sprintf(f1, x1, y1, x2, y2) // concentric hour short line

		x1, y1 = cartesianXY(r1, a+2.0*aQuaterHour)
		x2, y2 = cartesianXY(r2-0.9, a+2.0*aQuaterHour)
		s += fmt.Sprintf(f1, x1, y1, x2, y2) // concentric hour and half short line

		//improvement needed: to center an hour digit to middle of the arc
		ah := 0.3 * aQuaterHour
		if ra > 9 {
			ah *= 2.0
		}
		x1, y1 = cartesianXY(r3, a-ah)
		x2, y2 = cartesianXY(r3, a+ah)
		s += fmt.Sprintf(f2, ra, x2, y2, x1, y1, ra, ra) // circle arch for an hour number text
	}
	return fmt.Sprintf(f0, s)
}
func circleArchText(id, text string, r, a, deltaA float64) string {
	f1 := `      <path id="raHour%s" d="M%.1f,%.1f A%.1f,%.1f 0 0,0  %.1f,%.1f " style="fill:none;fill-opacity: 1;stroke:pink;stroke-width: 0.7"/>
      <text class="font1 downFont">
	    <textPath xlink:href="#raHour%s" text-anchor="start">%s</textPath>
      </text>

`
	s := ""
	x1, y1 := cartesianXY(r, a)
	x2, y2 := cartesianXY(r, a+deltaA)
	s += fmt.Sprintf(f1, id, x2, y2, r, r, x1, y1, id, text) // circle arch for an hour number text
	return s
}
func dateRoundScale() string {
	s := "      <g id=\"dateRoundScale\">\n"
	r1 := 172.0
	f1 := "        <circle cx=\"%.1f\" cy=\"%.1f\" r=\"%.1f\" stroke=\"black\" stroke-width=\"0.09\" fill=\"%s\" />\n"
	f1 = "<line x1=\"%.1f\" y1=\"%.1f\" x2=\"%.1f\" y2=\"%.1f\" class=\"cross\" />\n"
	aDelta := 2.0 * math.Pi / 365.0
	date := time.Date(2000, time.March, 21, 0, 0, 0, 0, time.UTC)
	for d := 0; d < 365; d++ {
		a := float64(d) * aDelta
		x1, y1 := cartesianXY(r1, a)
		r := 0.7
		if date.Day()%5 == 0 {
			r = 1.5
		}
		if date.Day()%10 == 0 {
			r = 4.0
		}
		if date.Day() == 1 {
			r = 5.5
			s += circleArchText("MONTH_"+date.Format("Jan"), date.Format("January"), r1+10.0, a, 27.0/31.0*math.Pi/6.0)
		}
		//s += fmt.Sprintf(f1, x1, y1, r, "black")
		x2, y2 := cartesianXY(r1-r, a)
		s += fmt.Sprintf(f1, x1, y1, x2, y2) // concentric day(1,5,10) short line
		date = date.Add(24 * time.Hour)
	}
	s += "      </g>\n"

	return s
}
func magToRadius(mag float64) float64 {
	if mag < magBrightest {
		mag = magBrightest
	}
	magRange := magMin - magBrightest
	rMag := 0.3 + 2.6*(magMin-mag)/magRange
	return rMag
}
func eqToCartesianXY(RA, De float64) (float64, float64) {
	r0 := 100.0
	a := RA * math.Pi / 180.0
	r := r0 * (1.0 - De/90.0)
	return cartesianXY(r, a)
}
func plotStars() string {
	s := "      <g id=\"plotStars\">\n"

	f1 := "        <circle cx=\"%.1f\" cy=\"%.1f\" r=\"%.1f\" stroke=\"white\" stroke-width=\"0.05\" fill=\"%s\" />\n"
	lowestDeclination := -45.0
	sort.SliceStable(SliceOfStars, func(i, j int) bool { return SliceOfStars[i].Mag < SliceOfStars[j].Mag })
	for _, star := range SliceOfStars {
		if star.Mag < magMin && star.De > lowestDeclination {
			x, y := eqToCartesianXY(star.RA, star.De)
			rMag := magToRadius(star.Mag)
			s += fmt.Sprintf(f1, x, y, rMag, "blue")
		}
	}
	s += "      </g>\n"

	return s
}

func plotConstellations() string {
	s := "      <g id=\"plotConstellations\">\n"
	f1 := "        <path d=\"%s\" stroke=\"red\" stroke-width=\"0.25\" fill=\"none\" />\n"
	d := ""
	for _, c := range SliceOfConstellations {
		for _, line := range c.Lines {
			x, y := eqToCartesianXY(line[0].RA, line[0].De)
			d += fmt.Sprintf("M%.1f,%.1f ", x, y)
			for i := 1; i < len(line); i++ {
				x, y = eqToCartesianXY(line[i].RA, line[i].De)
				d += fmt.Sprintf("L%.1f,%.1f ", x, y)
			}
		}
	}
	s += fmt.Sprintf(f1, d)
	s += "      </g>\n"

	return s
}

func LoadECSV(filename string) ([][]string, error) {
	rows := make([][]string, 0)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(bytes)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		cols := strings.Split(line, ",")
		if len(line) > 0 && !strings.HasPrefix(line, "#") && len(cols) > 1 {
			for i := 0; i < len(cols); i++ {
				cols[i] = strings.TrimSpace(cols[i])
			}
			rows = append(rows, cols)
		}
	}
	return rows, nil
}

func LoadStars(filename string) {
	if bytes, err := os.ReadFile(filename); err != nil {
		fmt.Printf("Error loading file %s: %s\n", filename, err.Error())
		return
	} else {
		if err1 := json.Unmarshal([]byte(bytes), &SliceOfStars); err1 != nil {
			fmt.Printf("Error unmarshaling content of the json file %s: %s\n", filename, err1.Error())
		}
	}
}
func LoadConstellations(filename string) {
	if bytes, err := os.ReadFile(filename); err != nil {
		fmt.Printf("Error loading file %s: %s\n", filename, err.Error())
		return
	} else {
		if err1 := json.Unmarshal([]byte(bytes), &SliceOfConstellations); err1 != nil {
			fmt.Printf("Error unmarshaling content of the json file %s: %s\n", filename, err1.Error())
		}
	}
}
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	//writeHtmlHeadAndMenu(w, "/", "Home")
	fmt.Fprint(w, `<html>
 <head>
 <meta http-equiv="refresh" content="0; url=/SkyMapLab">
 	  <title>redirect to SkyMapLab</title>
 </head>
 <body>
  <h1>Click to: <a href="/SkyMapLab">SkyMapLab</a></h1>
 </body>
</html>
	`)

	fmt.Fprint(w, htmlEnd)
}
func HandlerSkyMapLab(w http.ResponseWriter, r *http.Request) {
	//writeHtmlHeadAndMenu(w, "/svg-roundlogo-color", "Color")

	fmt.Fprint(w, "<h1>SkyMap Lab</h1>")
	fmt.Fprint(w, "<h1>SkyMap <a href=\"/img/svg-skymap-color\">Color</a></h1>")
	fmt.Fprint(w, "<h1>SkyMap <a href=\"/img/svg-skymap-bw\">Black and White</a></h1>")
}
func HandlerImageSkymapColor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	defs := raCross()
	defs += raHourRoundScale()
	defs += dateRoundScale()
	defs += plotConstellations()
	defs += plotStars()

	svgTemplate2 := fmt.Sprintf(svgTemplate1, defs)
	if t, err := template.New("SkyMap").Parse(svgTemplate2); err == nil {
		data := getSvgData(false)
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
	// Send the response
	//w.WriteHeader(http.StatusOK)
}
func HandlerImageSkymapBW(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	// Optional: Set additional headers if needed
	// w.Header().Set("Last-Modified", "...")
	if t, err := template.New("SvgRoundLogoBlackWhite").Parse(svgTemplate1); err == nil {
		data := getSvgData(false)
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
	// Send the response
	//w.WriteHeader(http.StatusOK)
}
