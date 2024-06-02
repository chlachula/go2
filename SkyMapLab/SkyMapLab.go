/*
https://www.astroleague.org/caldwell-program-object-list/
https://www.astroleague.org/messier-program-list/
*/

package SkyMapLab

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type StarRecord struct {
	RA  float64 `json:"RA"`
	De  float64 `json:"De"`
	Mag float64 `json:"Mag"`
}

type MapColors struct {
	ConstLine   string
	OuterCircle string
}

var MapColorsRed = MapColors{ConstLine: "red", OuterCircle: "#ffeee6"}
var MapBlackAndWhite = MapColors{ConstLine: "black", OuterCircle: "silver"}
var MapColorsOrange = MapColors{ConstLine: "orange", OuterCircle: "#f2e1e9"}

type MapStyle struct {
	NorthMap              bool
	Radius                float64
	Rlat                  float64
	RadiusDeclinationZero float64
	RAwidth               float64
	RAhour_length         float64
	RAhalfHour_length     float64
	RAciphersRadius       float64
	Latitude              float64
	LowestStarDecl        float64
	LowestConstDecl       float64
	Axis                  float64
	AxisWidth             float64
	ConstLineWidth        float64
	DateRadius            float64
	DateNRadius           float64
	MonthsRadius          float64
	MagMin                float64
	MagBrightest          float64
	MagMinName            float64
	Colors                MapColors
}

var SliceOfStars []StarRecord

// var magBrightest = -1.5 // Sirius
// var magMin = 5.0
var monthArcR = 27.0 / 31.0 * math.Pi / 6.0
var dayNArcR = 2. * 2. * math.Pi / 365.25
var SliceOfConstellations []ConstellationCoordPoints

var Map MapStyle

type EqCoords struct {
	RA float64 `json:"RA"`
	De float64 `json:"De"`
}
type EqPoints []EqCoords
type ConstellationCoordPoints struct {
	Abbr    string
	Name    string
	Genitiv string
	NameLoc EqCoords
	Lines   []EqPoints
}

const (
	htmlEnd      = "\n<br/></body></html>"
	svgTemplate1 = `
<svg xmlns="http://www.w3.org/2000/svg" 
    xmlns:xlink="http://www.w3.org/1999/xlink" 
	width="{{.WidthMM}}mm" height="{{.HeightMM}}mm" viewBox="-250 -{{.HeightHalf}} 500 {{.Height}}" 
	style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd;background:beige">
    <title>Sky Map Lab</title>
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
		stroke-width: {{.CrossStrokeWidth}};
		fill:none
	 }
    </style>
  %s
 </defs> 
  <rect width="500" height="{{.Height}}" x="-250" y="-{{.HeightHalf}}" stroke="blue" stroke-width="1" fill="azure" />
  <text x="-244" y="-{{.HeightHalf}}" fill="blue" font-size="8"><tspan dy="10">{{.PaperName}} ({{.WidthMM}}mm by {{.HeightMM}}mm)</tspan></text>
  <g id="draw_plots">
    <use xlink:href="#plotConstellations" />
    <use xlink:href="#plotConstellationNames" />
    <use xlink:href="#plotOuterCircle" />
    <use xlink:href="#plotEcliptic" />
    <use xlink:href="#plotHorizon" />
    <use xlink:href="#plotStars" />
    <use xlink:href="#plotDateRoundScale" />
    <use xlink:href="#plotRaHourScale" />
    <use xlink:href="#plotRaCross" />
  </g>

</svg>
`
)

type SvgDataType = struct {
	FontSize         float64
	TopColor         string
	BottomColor      string
	CrossStrokeWidth float64
	PaperName        string
	WidthMM          float64
	HeightMM         float64
	Height           float64
	HeightHalf       float64
}

var (
	TopText    string
	BottomText string
)

func SetMapStyle(r, lat float64, c MapColors) {
	var m MapStyle
	m.NorthMap = false
	if lat > 0.0 {
		m.NorthMap = true
	}
	//	r2 := Map.Rlat * 1.12 //170
	//	w := Map.Rlat * 0.23  //40 1.12+0.23/2=1.235

	m.Rlat = r * 0.80971659919028340080971659919028
	m.Latitude = lat
	m.Colors = c
	m.Axis = r * 0.82995951417004048582995951417004
	m.AxisWidth = r * 0.00202429149797570850202429149798
	m.RAwidth = r * 0.01052631578947368421052631578947
	m.RAhour_length = r * 0.02672064777327935222672064777328
	m.RAhalfHour_length = r * 0.02186234817813765182186234817814
	m.RAciphersRadius = r * 0.87449392712550607287449392712551
	m.ConstLineWidth = r * 0.00161943319838056680161943319838
	m.DateRadius = r * 0.92874493927125506072874493927126
	m.DateNRadius = r * 0.945
	m.MonthsRadius = r * 0.98137651821862348178137651821862
	m.MagBrightest = -1.5 // Sirius
	m.MagMin = 5.0
	m.MagMinName = 1.0

	m.RadiusDeclinationZero = 90.0 * m.Rlat / (180.0 + lat)
	m.LowestConstDecl = 60.0      //Southern sky map
	m.LowestStarDecl = lat + 90.0 //Southern sky map
	if m.NorthMap {
		m.RadiusDeclinationZero = 90.0 * m.Rlat / (180.0 - lat)
		m.LowestConstDecl *= -1.0     // Northern sky map
		m.LowestStarDecl = lat - 90.0 // Northern sky map -45
	}
	Map = m
}

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
sinTcos𝛿 = cosHsinA
cosTcos𝛿 = cosFsinH + sinFcosHcosA
sin𝛿 = sinFsinH - cosFcosHcosA
*/
func AzimutalToEquatoreal_I(A, h, fi float64) (float64, float64) {
	sinT := math.Cos(h) * math.Sin(A)
	cosT := math.Cos(fi)*math.Sin(h) + math.Sin(fi)*math.Cos(h)*math.Cos(A)
	sinD := math.Sin(fi)*math.Sin(h) - math.Cos(fi)*math.Cos(h)*math.Cos(A)
	t := math.Atan2(sinT, cosT)
	if t < 0.0 {
		t += 2.0 * math.Pi
	}
	de := math.Asin(sinD)
	return t, de
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
	type paperType = struct {
		Name     string
		WidthMM  float64
		HeightMM float64
		Height   float64
	}
	var papers = []paperType{
		{Name: "A4", WidthMM: 210.0, HeightMM: 297.0},
		{Name: "A3", WidthMM: 297.0, HeightMM: 420.0},
		{Name: "Letter 8.5\"x11\"", WidthMM: 215.9, HeightMM: 279.4},
		{Name: "Legal 8.5\"x14\"", WidthMM: 215.9, HeightMM: 355.6},
		{Name: "Ledger 11\"x17\"", WidthMM: 279.4, HeightMM: 431.8},
	}
	factor := Map.Rlat / 150.0

	i := 4
	data := SvgDataType{
		TopColor:         "green",
		BottomColor:      "red",
		FontSize:         8.0 * factor,
		CrossStrokeWidth: 0.25 * factor,
		PaperName:        papers[i].Name,
		WidthMM:          papers[i].WidthMM,
		HeightMM:         papers[i].HeightMM,
		Height:           500.0 * papers[i].HeightMM / papers[i].WidthMM,
		HeightHalf:       250.0 * papers[i].HeightMM / papers[i].WidthMM,
	}
	if !color {
		data.TopColor = "black"
		data.BottomColor = "darkgray"
	}
	return data
}

func cartesianXY(r, a float64) (float64, float64) {
	x := -r * math.Sin(a)
	y := r * math.Cos(a)
	return x, y
}
func declinationToRadius(decl float64) float64 {
	var r1 float64
	if Map.NorthMap {
		r1 = Map.RadiusDeclinationZero * (1.0 - decl/90.0)
	} else {
		r1 = Map.RadiusDeclinationZero * (1.0 + decl/90.0)
	}
	return r1
}
func eqToCartesianXY(RA, De float64) (float64, float64) {
	a := RA * math.Pi / 180.0
	r1 := declinationToRadius(De)
	return cartesianXY(r1, a)
}

func plotRaCross() string {
	r2 := Map.Axis //154
	w := Map.AxisWidth
	formCross := `
	<g id="plotRaCross">	  
      <line x1="-%.1f" y1="0" x2="%.1f" y2="0" class="cross" />
	  <line x1="0" y1="-%.1f" x2="0" y2="%.1f" class="cross" />
	  <circle cx="0" cy="0" r="%.1f" stroke="black" stroke-width="%.1f" fill="none" />
	</g>
`
	return fmt.Sprintf(formCross, r2, r2, r2, r2, Map.RadiusDeclinationZero, w)
}

func plotRaHourRoundScale() string {
	r1 := Map.Rlat                         //150
	r2 := Map.Rlat + Map.RAwidth           //152
	r3 := Map.Rlat + Map.RAhour_length     //155
	r4 := Map.Rlat + Map.RAhalfHour_length //154
	//	r5 := Map.Rlat * 1.08                  //162

	form0 := `
	<g id="plotRaHourScale">
	  <circle cx="0" cy="0" r="%.1f" stroke="black" stroke-width="0.5" fill="none" />
	  <circle cx="0" cy="0" r="%.1f" stroke="black" stroke-width="0.5" fill="none" />
	  %s
	</g>
`
	s := "\n"
	form1 := "      <line x1=\"%.1f\" y1=\"%.1f\" x2=\"%.1f\" y2=\"%.1f\" class=\"cross\" />\n"
	strokeWidth := 0.3 / 150.0 * Map.Rlat
	form2 := `      <path id="raHour%d" d="M%.1f,%.1f A%.1f,%.1f 0 0,0  %.1f,%.1f " style="fill:none;fill-opacity: 1;stroke:green;stroke-width: %.1f"/>
      <text alignment-baseline="baseline" text-anchor="start" class="font1 downFont">
	    <textPath xlink:href="#raHour%d">%d</textPath>
      </text>

`
	aQuaterHour := math.Pi / 48.0
	for ra := 0; ra <= 23; ra++ {
		a := float64(ra*15) * math.Pi / 180.0
		x1, y1 := cartesianXY(r1, a)
		x2, y2 := cartesianXY(r3, a)
		s += fmt.Sprintf(form1, x1, y1, x2, y2) // concentric hour short line

		x1, y1 = cartesianXY(r1, a+2.0*aQuaterHour)
		x2, y2 = cartesianXY(r4, a+2.0*aQuaterHour)
		s += fmt.Sprintf(form1, x1, y1, x2, y2) // concentric hour and half short line

		//improvement needed: to center an hour digit to middle of the arc
		ah := 0.3 * aQuaterHour
		if ra > 9 {
			ah *= 2.0
		}
		x1, y1 = cartesianXY(Map.RAciphersRadius, a-ah)
		x2, y2 = cartesianXY(Map.RAciphersRadius, a+ah)
		r := Map.RAciphersRadius
		s += fmt.Sprintf(form2, ra, x2, y2, r, r, x1, y1, strokeWidth, ra, ra) // circle arch for an hour number text
	}
	return fmt.Sprintf(form0, r1, r2, s)
}

func circleArchText(id, text string, r, a, deltaA float64, strokeColor string, fillColor string, fontSize float64) string {
	form1 := `       <path id="%s" d="M%.1f,%.1f A%.1f,%.1f 0 0,0  %.1f,%.1f " style="fill:none;fill-opacity: 1;stroke:%s;stroke-width: 0.7"/>
       <text font-size="%.1f" font-family="Franklin Gothic, sans-serif" fill="%s" >
	     <textPath xlink:href="#%s" text-anchor="start">%s</textPath>
       </text>

`
	s := ""
	x1, y1 := cartesianXY(r, a)
	x2, y2 := cartesianXY(r, a+deltaA)
	s += fmt.Sprintf(form1, id, x2, y2, r, r, x1, y1, strokeColor, fontSize, fillColor, id, text) // circle arch for an hour number text
	return s
}
func tangetDirective(a float64) float64 {
	k := math.Cos(a)
	return k
}
func tangentText(id, text string, r, a, length float64, strokeColor string, fillColor string, fontSize float64) string {
	form1 := `       <path id="%s" d="M%.1f,%.1f l%.1f,%.1f " style="fill:none;fill-opacity: 1;stroke:%s;stroke-width: 0.7"/>
       <text font-size="%.1f" font-family="Franklin Gothic, sans-serif" fill="%s" >
	     <textPath xlink:href="#%s" text-anchor="start">%s</textPath>
       </text>

`
	s := ""
	x1, y1 := cartesianXY(r, a)
	k := tangetDirective(a)
	dx := k * length
	dy := math.Sqrt(length*length - dx*dx)
	if a > math.Pi {
		dy = -dy
	}
	s += fmt.Sprintf(form1, id, x1, y1, dx, dy, strokeColor, fontSize, fillColor, id, text) // circle arch for an hour number text
	return s
}
func plotDateRoundScale() string {
	s := "     <g id=\"plotDateRoundScale\">\n"
	//r1 := Map.Rlat * 1.147 // 172.0
	r1 := Map.DateRadius //172
	form1 := "       <line x1=\"%.1f\" y1=\"%.1f\" x2=\"%.1f\" y2=\"%.1f\" class=\"cross\" />\n"
	aDelta := 2.0 * math.Pi / 365.0
	date := time.Date(2000, time.March, 21, 0, 0, 0, 0, time.UTC)
	for d := 0; d < 365; d++ {
		a := float64(d) * aDelta
		x1, y1 := cartesianXY(r1, a)
		//r := 0.7
		bar := r1 * 0.004067 //0.7
		if date.Day()%5 == 0 {
			//r = 1.5
			bar = r1 * 0.008721 //1.5
		}
		if date.Day()%10 == 0 {
			//r = 4.0
			bar = r1 * 0.023256 //4
		}
		if date.Day() == 1 {
			s += circleArchText("DAY_"+date.Format("Jan02"), date.Format("2"), Map.DateNRadius, a, dayNArcR, "pink", "black", Map.Rlat*0.025)
			//r = 5.5
			bar = r1 * 0.031978 //5.5
			s += circleArchText("MONTH_"+date.Format("Jan"), date.Format("January"), Map.MonthsRadius, a, monthArcR, "yellow", "red", Map.Rlat*0.06)
		}
		//s += fmt.Sprintf(f1, x1, y1, r, "black")
		x2, y2 := cartesianXY(r1-bar, a)
		s += fmt.Sprintf(form1, x1, y1, x2, y2) // concentric day(1,5,10) short line
		date = date.Add(24 * time.Hour)
	}
	s += "      </g>\n"

	return s
}
func magToRadius(mag float64) float64 {
	r0 := Map.Rlat / 500.0 // 0.3
	r1 := Map.Rlat / 58.0  //2.6
	if mag < Map.MagBrightest {
		mag = Map.MagBrightest
	}
	magRange := Map.MagMin - Map.MagBrightest
	rMag := r0 + r1*(Map.MagMin-mag)/magRange
	return rMag
}
func constellationCanBeVisible(m MapStyle, c ConstellationCoordPoints) bool {
	if m.NorthMap && c.NameLoc.De > m.LowestConstDecl {
		return true
	}
	if !m.NorthMap && c.NameLoc.De < m.LowestConstDecl {
		return true
	}
	return false
}
func starCanBeVisible(m MapStyle, star StarRecord) bool {
	if m.NorthMap && star.De > m.LowestStarDecl {
		return true
	}
	if !m.NorthMap && star.De < m.LowestStarDecl {
		return true
	}
	return false
}
func plotStars() string {
	s := "      <g id=\"plotStars\">\n"

	form1 := "        <circle cx=\"%.1f\" cy=\"%.1f\" r=\"%.1f\" stroke=\"white\" stroke-width=\"0.05\" fill=\"%s\" />\n"
	sort.SliceStable(SliceOfStars, func(i, j int) bool { return SliceOfStars[i].Mag < SliceOfStars[j].Mag })
	for _, star := range SliceOfStars {
		if star.Mag < Map.MagMin && starCanBeVisible(Map, star) {
			x, y := eqToCartesianXY(star.RA, star.De)
			rMag := magToRadius(star.Mag)
			s += fmt.Sprintf(form1, x, y, rMag, "blue")
		}
	}
	s += "      </g>\n"

	return s
}
func plotStarNames() string {
	s := "      <g id=\"plotStarNames\">\n"

	form1 := "        <circle cx=\"%.1f\" cy=\"%.1f\" r=\"%.1f\" stroke=\"white\" stroke-width=\"0.05\" fill=\"%s\" />\n"
	sort.SliceStable(SliceOfStars, func(i, j int) bool { return SliceOfStars[i].Mag < SliceOfStars[j].Mag })
	for _, star := range SliceOfStars {
		if star.Mag < Map.MagMinName && starCanBeVisible(Map, star) {
			x, y := eqToCartesianXY(star.RA, star.De)
			rMag := magToRadius(star.Mag)
			s += fmt.Sprintf(form1, x, y, rMag, "blue")
		}
	}
	s += "      </g>\n"

	return s
}
func plotOuterCircle() string {
	r2 := Map.Rlat * 1.12 //170
	w := Map.Rlat * 0.23  //40
	s := "      <g id=\"plotOuterCircle\">\n"
	form1 := "        <circle cx=\"0\" cy=\"0\" r=\"%.1f\" stroke-width=\"%.1f\" stroke=\"%s\" fill=\"none\" />\n"
	s += fmt.Sprintf(form1, r2, w, Map.Colors.OuterCircle)
	s += "      </g>\n"
	return s
}

func plotConstellations() string {
	s := "      <g id=\"plotConstellations\">\n"
	form1 := "        <path d=\"%s\" stroke=\"%s\" stroke-width=\"%.2f\" fill=\"none\" />\n"
	d := ""
	for _, c := range SliceOfConstellations {
		if constellationCanBeVisible(Map, c) {
			for _, line := range c.Lines {
				x, y := eqToCartesianXY(line[0].RA, line[0].De)
				d += fmt.Sprintf("M%.1f,%.1f ", x, y)
				for i := 1; i < len(line); i++ {
					x, y = eqToCartesianXY(line[i].RA, line[i].De)
					d += fmt.Sprintf("L%.1f,%.1f ", x, y)
				}
			}
		}
	}
	s += fmt.Sprintf(form1, d, Map.Colors.ConstLine, Map.ConstLineWidth)
	s += "      </g>\n"

	return s
}
func plotConstellationNames() string {
	s := "      <g id=\"plotConstellationNames\">\n"
	for _, c := range SliceOfConstellations {
		if constellationCanBeVisible(Map, c) {
			cId := fmt.Sprintf("CONST_%s", c.Abbr)
			raR := c.NameLoc.RA * math.Pi / 180.0
			s += tangentText(cId, c.Abbr, declinationToRadius(c.NameLoc.De), raR, Map.Rlat*0.06, "none", "green", Map.Rlat*0.035) // #d5ff80
		}
	}
	s += "      </g>\n"

	return s
}
func plotEcliptic() string {
	s := "      <g id=\"plotEcliptic\">\n"
	form1 := "        <path d=\"%s\" stroke=\"orange\" stroke-width=\"0.25\" fill=\"none\" />\n"
	toRad := math.Pi / 180.0
	toDeg := 180.0 / math.Pi
	x, y := eqToCartesianXY(0.0, 0.0)
	d := fmt.Sprintf("M%.1f,%.1f L", x, y)
	for la := 1.0; la < 360.1; la = la + 1.0 {
		ra, de := EclipticalToEquatorial(la*toRad, 0.0)
		x, y := eqToCartesianXY(ra*toDeg, de*toDeg)
		d += fmt.Sprintf("%.1f,%.1f ", x, y)
	}
	s += fmt.Sprintf(form1, d)
	s += "      </g>\n"

	return s
}

func plotHorizon() string {
	s := "      <g id=\"plotHorizon\" transform=\"rotate(180)\" >\n"
	form1 := "        <path d=\"%s\" stroke=\"green\" stroke-width=\"0.25\" fill=\"none\" />\n"
	toRad := math.Pi / 180.0
	toDeg := 180.0 / math.Pi
	hR := 0.0
	fi := Map.Latitude
	fiR := fi * toRad
	t, de := AzimutalToEquatoreal_I(0.0, hR, fiR)
	x, y := eqToCartesianXY(t*toDeg, de*toDeg)
	d := fmt.Sprintf("M%.1f,%.1f L", x, y)
	for az := 1.0; az < 360.1; az = az + 1.0 {
		t, de = AzimutalToEquatoreal_I(az*toRad, hR, fiR)
		x, y = eqToCartesianXY(t*toDeg, de*toDeg)
		d += fmt.Sprintf("%.1f,%.1f ", x, y)
	}
	s += fmt.Sprintf(form1, d)
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
 	  <title>Redirect to SkyMapLab</title>
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

	fmt.Fprint(w, "<html><head><title>SkyMap</title></head>\n")
	fmt.Fprint(w, "<body style=\"text-align: center;\">\n")
	fmt.Fprint(w, "<h1>SkyMap Lab</h1>\n")
	fmt.Fprint(w, `<table border="1" style="margin: 0px auto;">
	<tr><td></td><td>N</td><td>S</td></tr>
	<tr><td>color</td><td><a href="/img/svg/skymap/co/n44">+44</a></td><td><a href="/img/svg/skymap/co/s44">-44</a></td></tr>
	<tr><td>b&amp;w</td><td><a href="/img/svg/skymap/bw/n44">+44</a></td><td><a href="/img/svg/skymap/bw/s44">-44</a></td></tr>
	</table>`)
}
func getLatitude(str string) float64 {
	sign := 1
	if strings.HasPrefix(str, "s") || strings.HasPrefix(str, "S") {
		sign = -1
	}
	if f, err := strconv.ParseFloat(str[1:], 64); err != nil {
		return 50.0
	} else {
		return float64(sign) * f
	}
}
func HandlerSkyMapGeneral(w http.ResponseWriter, r *http.Request) {
	colorId := r.PathValue("colorId")
	lat := getLatitude(r.PathValue("latId"))
	if strings.HasPrefix(colorId, "co") {
		SetMapStyle(249.0, lat, MapColorsRed)
	} else {
		SetMapStyle(249.0, lat, MapBlackAndWhite)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	defs := plotRaCross()
	defs += plotRaHourRoundScale()
	defs += plotDateRoundScale()
	defs += plotConstellations()
	defs += plotConstellationNames()
	defs += plotOuterCircle()
	defs += plotEcliptic()
	defs += plotHorizon()
	defs += plotStars()

	svgTemplate2 := fmt.Sprintf(svgTemplate1, defs)
	if t, err := template.New("SkyMap").Parse(svgTemplate2); err == nil {
		data := getSvgData(strings.HasPrefix(colorId, "co"))
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
	//w.WriteHeader(http.StatusOK)
}
