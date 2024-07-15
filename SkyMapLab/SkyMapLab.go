/*
https://www.astroleague.org/caldwell-program-object-list/
https://www.astroleague.org/messier-program-list/

How long is my SVG <text> element? https://www.balisage.net/Proceedings/vol26/html/Birnbaum01/BalisageVol26-Birnbaum01.html
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

type ObjectRecord struct {
	CName string  `json:"CName"` // Common name
	OType string  `json:"OType"` // Object type
	Const string  `json:"Const"` // Constellation
	Mes   int     `json:"Mes"`   // Messier
	Cal   int     `json:"Cal"`   // Caldwell
	NGC   int     `json:"NGC"`   // New General Catalogue
	IC    int     `json:"IC"`    // Index Catalogue
	RA    float64 `json:"RA"`
	De    float64 `json:"De"`
	Mag   float64 `json:"Mag"`
	Size  string  `json:"Size"`
}

type MapColors struct {
	ConstLine   string
	ConstName   string
	OuterCircle string
	Months      string
	Star        string
	Ecliptic    string
	Horizon     string
}

var (
	toRad    = math.Pi / 180.0
	toDeg    = 180.0 / math.Pi
	𝜀Deg2025 = 23.436040

	MapColorsRed     = MapColors{ConstLine: "red", OuterCircle: "#ffeee6", Months: "black", ConstName: "green", Star: "blue", Ecliptic: "orange", Horizon: "green"}
	MapBlackAndWhite = MapColors{ConstLine: "black", OuterCircle: "silver", Months: "black", ConstName: "black", Star: "black", Ecliptic: "black", Horizon: "black"}
	MapColorsOrange  = MapColors{ConstLine: "orange", OuterCircle: "#f2e1e9", Months: "orange", ConstName: "green", Star: "darkblue", Ecliptic: "yellow", Horizon: "darkgreen"}
)

type MapStyle struct {
	NorthMap bool
	//	Radius                float64
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
	DashedEcliptic        bool
	DashedHorizon         bool
}

var SliceOfStars []StarRecord
var SliceOfObjects []ObjectRecord

// var magBrightest = -1.5 // Sirius
// var magMin = 5.0
const dayArcR = 2.0 * math.Pi / 365.25
const twoDaysArcR = 2. * dayArcR
const monthArcR = 27.0 / 31.0 * math.Pi / 6.0

var SliceOfConstellations []ConstellationCoordPoints
var Map MapStyle

type isoLatitudeCircleToEq func(float64, float64, float64) (float64, float64)

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

const (
	htmlEnd      = "\n<br/></body></html>"
	svgTemplate1 = `
<svg xmlns="http://www.w3.org/2000/svg" 
    xmlns:xlink="http://www.w3.org/1999/xlink" 
	width="{{.WidthMM}}mm" height="{{.HeightMM}}mm" viewBox="{{.VBminX}} {{.VBminY}} {{.VBwidth}} {{.VBheight}}" 
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
	 .fontAxis{
		font-size: {{.FontSizeAxis}}px;
		font-family: Franklin Gothic, sans-serif;
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
	<marker 
      id='arrow_head' 
      orient="auto" 
      markerWidth='15' 
      markerHeight='20' 
      refX='0.1' 
      refY='4'
     >
     <path d='M0,0 V8 L4,4 Z' fill="black" />
    </marker>
  %s
	<g id="draw_AZ_grid" transform="rotate(180)">
	  <circle cx="0" cy="0" r="{{.RLat}}" stroke="black" stroke-width="0.5" fill="none" />
      <use xlink:href="#plotHorizon" />
      <use xlink:href="#plotAlmucantarats" />
	  <use xlink:href="#plotMeridians" />	
    </g>
  <g id="draw_map">
    <use xlink:href="#plotConstellations" />
    <use xlink:href="#plotConstellationNames" />
    <use xlink:href="#plotOuterCircle" />
    <use xlink:href="#plotEcliptic" />
    <use xlink:href="#plotStars" />
    <use xlink:href="#plotObjects" />
    <use xlink:href="#plotDateRoundScale" />
    <use xlink:href="#plotRaHourScale" />
    <use xlink:href="#plotRaCross" />
    <use xlink:href="#plotAxisDeclinations" />
    <use xlink:href="#plotDirectionsOfTheApparentRotationOfTheSky" />	
  </g>
  <g id="draw_platonYear_map">
    <use xlink:href="#plotPlatonYear" />
    <use xlink:href="#draw_map" />
  </g>
  <g id="draw_all">
    <use xlink:href="#plotPlatonYear" />
    <use xlink:href="#draw_map" />
    <use xlink:href="#draw_AZ_grid" />
  </g>
 </defs> 
  <rect width="500" height="{{.Height}}" x="-250" y="-{{.HeightHalf}}" stroke="blue" stroke-width="1" fill="azure" />
  <text x="-244" y="-{{.HeightHalf}}" fill="blue" font-size="8"><tspan dy="10">{{.PaperName}} ({{.WidthMM}}mm by {{.HeightMM}}mm) - Latitude: {{.Latitude}}</tspan></text>
  
  <use xlink:href="#%s" />
</svg>
`
)

type SvgDataType = struct {
	FontSize         float64
	FontSizeAxis     float64
	TopColor         string
	BottomColor      string
	CrossStrokeWidth float64
	PaperName        string
	Latitude         string
	RLat             float64
	WidthMM          float64
	HeightMM         float64
	Height           float64
	HeightHalf       float64
	VBminX           float64
	VBminY           float64
	VBwidth          float64
	VBheight         float64
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
cos𝛼*cos𝛿 = cos𝛽*cos𝜆 => cos𝛿 = cos𝛽*cos𝜆 / cos𝛼
sin𝛼*cos𝛿 = cos𝛽*sin𝜆*cos𝜀 − sin𝜀*sin𝛽 = sin𝛼/cos𝛼 * cos𝛽*cos𝜆 = tan𝛼*cos𝛽*cos𝜆
sin𝛿 = sin𝛽*cos𝜀 + sin𝜀*cos𝛽*sin𝜆
*/
func EclipticalToEquatorial(La, Be, 𝜀 float64) (float64, float64) {
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
func viewBox(w float64, h float64, rows, colums, rIndex, cIndex int) (float64, float64, float64, float64) {
	minX, minY := -0.5*w, -0.5*h
	width, height := w, h
	dw, dh := w, h
	if rows != 0 || colums != 0 {
		if rows != 0 {
			dh /= float64(rows)
		}
		if colums != 0 {
			dw /= float64(colums)
		}
	}
	minX += float64(cIndex) * dw
	minY += float64(rIndex) * dh
	width = dw
	height = dh
	return minX, minY, width, height
}
func getSvgData(color bool, i int) SvgDataType {
	factor := Map.Rlat / 150.0

	width := 500.0
	height := width * papers[i].HeightMM / papers[i].WidthMM
	rowsNumber, columnsNumber := 0, 0
	rowIndex, columnIndex := 0, 0
	vMinX, vMinY, vWidth, vHeight := viewBox(width, height, rowsNumber, columnsNumber, rowIndex, columnIndex)
	data := SvgDataType{
		TopColor:         "green",
		BottomColor:      "black",
		FontSize:         8.0 * factor,
		FontSizeAxis:     4.0 * factor,
		CrossStrokeWidth: 0.25 * factor,
		Latitude:         fmt.Sprintf("%.f", Map.Latitude),
		RLat:             Map.Rlat,
		PaperName:        papers[i].Name,
		WidthMM:          papers[i].WidthMM,
		HeightMM:         papers[i].HeightMM,
		Height:           height,
		HeightHalf:       0.5 * height,
		VBminX:           vMinX,
		VBminY:           vMinY,
		VBwidth:          vWidth,
		VBheight:         vHeight,
	}
	if !color {
		data.TopColor = "black"
		data.BottomColor = "darkgray"
	}
	return data
}

/*
https://go.dev/play/p/u2437C2rRFG
Special characters needing encoding are: ':', '/', '?', '#', '[', ']', '@', '!', '$', '&', "'", '(', ')', '*', '+', ',', ';', '=', as well as '%' itself.
URL special characters not needing percent coding: 34"  45-  46.  60<  62>  92\  94^  95_  96`  123{  124|  125}  126~
*/

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
func plotDirectionsOfTheApparentRotationOfTheSky() string {
	form1 := `
	    <path id="dirArrow" d="M%.1f,%.1f A%.1f,%.1f 0 0,0 %.1f,%.1f " 
		      style="fill:none;stroke:black;stroke-width: 0.432"  marker-end="url(#arrow_head)" />
	    <text alignment-baseline="baseline" text-anchor="start" font-size="3.4" font-family="Franklin Gothic, sans-serif" fill="black" dy="-1.0">
	      <textPath xlink:href="#dirArrow">%s</textPath>
        </text>
`
	arcAngle := 14.0
	a1 := (90.0 - arcAngle) * 0.5
	a2 := a1 + arcAngle
	r := Map.RadiusDeclinationZero * 2.1
	x1, y1 := cartesianXY(r, a1*toRad)
	x2, y2 := cartesianXY(r, a2*toRad)
	g := "\n      <g id=\"directionOfTheApparentRotationOfTheSky\">"
	text := "direction of the apparent rotation of the sky"
	g += fmt.Sprintf(form1, x2, y2, r, r, x1, y1, text)
	g += "      </g>\n"

	paths := "      <g id=\"plotDirectionsOfTheApparentRotationOfTheSky\" >\n"
	paths += "        <use xlink:href=\"#directionOfTheApparentRotationOfTheSky\" />\n"
	paths += "        <use xlink:href=\"#directionOfTheApparentRotationOfTheSky\"  transform=\"rotate(090)\" />\n"
	paths += "        <use xlink:href=\"#directionOfTheApparentRotationOfTheSky\"  transform=\"rotate(180)\" />\n"
	paths += "        <use xlink:href=\"#directionOfTheApparentRotationOfTheSky\"  transform=\"rotate(270)\" />\n"
	paths += "      </g>\n"
	return g + paths
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
func plotAxisDeclinations() string {
	stepDegs := 30.0
	a := 90.0 - stepDegs
	dx := 3.5
	dy := Map.RadiusDeclinationZero / 90.0 * stepDegs
	y := dy
	g := "\n    <g id=\"AxisDeclMarks\">"
	texts := ""
	path := fmt.Sprintf("\n       <path d=\"M%.1f,0 ", dx)
	for ; a > -44.0; a = a - stepDegs {
		//exclude marks for zero declination
		if math.Abs(a) < 1.0 {
			path += fmt.Sprintf("m0,%.1f ", dy)
		} else {
			path += fmt.Sprintf("m%.1f,%.1f h%.1f ", -2.*dx, dy, 2.*dx)
		}
		texts += fmt.Sprintf("  	   <text x=\"%.1f\" y=\"%.1f\" class=\"fontAxis\" >%.0f°</text>\n", 0.25*dx, y-0.25*dx, a)
		y += dy
	}
	path += "\" style=\"fill:none;stroke:black;stroke-width: 0.432\" />\n"
	g += path + texts + "    </g>\n"

	paths := "    <g id=\"plotAxisDeclinations\" >\n"
	paths += "        <use xlink:href=\"#AxisDeclMarks\" />\n"
	paths += "        <use xlink:href=\"#AxisDeclMarks\"  transform=\"rotate(090)\" />\n"
	paths += "        <use xlink:href=\"#AxisDeclMarks\"  transform=\"rotate(180)\" />\n"
	paths += "        <use xlink:href=\"#AxisDeclMarks\"  transform=\"rotate(270)\" />\n"
	paths += "    </g>\n"

	return g + paths
}
func plotRaHourRoundScale() string {
	r1 := Map.Rlat                         //150
	r2 := Map.Rlat + Map.RAwidth           //152
	r3 := Map.Rlat + Map.RAhour_length     //155
	r4 := Map.Rlat + Map.RAhalfHour_length //154
	//	r5 := Map.Rlat * 1.08                  //162

	form0 := `
	<g id="plotRaHourScale">
	  <circle cx="0" cy="0" r="%.1f" stroke="white" stroke-width="%.1f" fill="none" />
	  <circle cx="0" cy="0" r="%.1f" stroke="black" stroke-width="0.5" fill="none" />
	  <circle cx="0" cy="0" r="%.1f" stroke="black" stroke-width="0.5" fill="none" />
	  %s
	</g>
`
	s := "\n"
	form1 := "      <line x1=\"%.1f\" y1=\"%.1f\" x2=\"%.1f\" y2=\"%.1f\" class=\"cross\" />\n"
	strokeWidth := 0.3 / 150.0 * Map.Rlat
	form2 := `      <path id="raHour%d" d="M%.1f,%.1f A%.1f,%.1f 0 0,0  %.1f,%.1f " style="fill:none;fill-opacity: 1;stroke:none;stroke-width: %.1f"/>
      <text alignment-baseline="baseline" text-anchor="start" class="font1 downFont">
	    <textPath xlink:href="#raHour%d">%d</textPath>
      </text>

`
	aQuaterHour := math.Pi / 48.0
	aMinute := math.Pi / 720.0
	for ra := 0; ra <= 23; ra++ {
		a := float64(ra*15) * math.Pi / 180.0
		x1, y1 := cartesianXY(r1, a)
		x2, y2 := cartesianXY(r3, a)
		s += fmt.Sprintf(form1, x1, y1, x2, y2) // concentric hour short line

		for min := 5; min <= 55; min = min + 5 {
			a5min := a + float64(min)*aMinute
			x1, y1 = cartesianXY(r1, a5min)
			x2, y2 = cartesianXY(r2, a5min)
			s += fmt.Sprintf(form1, x1, y1, x2, y2) // concentric 5 minutes steps
		}

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
	r0 := 0.5 * (r1 + r2)
	w0 := 0.5 * (r2 - r1)
	return fmt.Sprintf(form0, r0, w0, r1, r2, s)
}
func circleDayN(date time.Time, a float64) string {
	n := date.Format("2")
	a1 := a - twoDaysArcR + twoDaysArcR*0.18*float64(len(n))
	return circleArchText("DAY_"+date.Format("Jan02"), n, Map.DateNRadius, a1, twoDaysArcR, "none", "black", Map.Rlat*0.025)
}
func circleMonthN(date time.Time, a float64) string {
	n := date.Format("January")
	a1 := a - 1.3*dayArcR*float64(len(n))
	return circleArchText("MONTH_"+date.Format("Jan"), n, Map.MonthsRadius, a1, monthArcR, "none", Map.Colors.Months, Map.Rlat*0.06)
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
		if date.Day() == 1 {
			//			s += circleArchText("MONTH_"+date.Format("Jan"), date.Format("January"), Map.MonthsRadius, a, monthArcR, "yellow", "red", Map.Rlat*0.06)
			s += circleMonthN(date, a)
			s += circleDayN(date, a)
			bar = r1 * 0.031978 //5.5
		} else if date.Day()%5 == 0 {
			if date.Day()%10 == 0 {
				s += circleDayN(date, a)
				bar = r1 * 0.023256 //4
			} else {
				s += circleDayN(date, a)
				bar = r1 * 0.008721 //1.5
			}
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
			s += fmt.Sprintf(form1, x, y, rMag, Map.Colors.Star)
		}
	}
	s += "      </g>\n"

	return s
}
func plotObject(obj ObjectRecord) string {
	s := fmt.Sprintf("       <g id=\"obj_%s%d\">\n", "M", obj.Mes)
	x, y := eqToCartesianXY(obj.RA, obj.De)
	rMag := magToRadius(0.0)
	width := 0.1 * rMag
	dash := width
	color := "brown"
	//<circle r="45" cx="350" cy="100" fill="pink" stroke="blue" stroke-width="4" stroke-dasharray="10,5" />
	formGC := "        <circle cx=\"%.1f\" cy=\"%.1f\" r=\"%.1f\" stroke=\"%s\" stroke-width=\"%.1f\" stroke-dasharray=\"%.1f,%.1f\"  fill=\"none\" />\n"
	formGA := "        <ellipse cx=\"%.1f\" cy=\"%.1f\" rx=\"%.1f\" ry=\"%.1f\" stroke=\"%s\" stroke-width=\"%.1f\" stroke-dasharray=\"%.1f,%.1f\"  fill=\"none\" />\n"
	if obj.OType != "whatever" {
		s += fmt.Sprintf(formGA, x, y, rMag, rMag*0.5, color, width, dash, dash)
	} else {
		s += fmt.Sprintf(formGC, x, y, rMag, color, width, dash, dash)
	}
	s += "       </g>\n"

	return s
}
func plotObjects() string {
	s := "      <g id=\"plotObjects\">\n"
	for _, obj := range SliceOfObjects {
		s += plotObject(obj)
	}
	s += "      </g>\n"

	return s
}
func P_lotStarNames() string {
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
			s += tangentText(cId, c.Abbr, declinationToRadius(c.NameLoc.De), raR, Map.Rlat*0.06, "none", Map.Colors.ConstName, Map.Rlat*0.035) // #d5ff80
		}
	}
	s += "      </g>\n"

	return s
}

// great circle or orthodrome is the circular intersection of a sphere and a plane passing through the sphere's center point
func plotGreatCircle(color string, dashed bool, fixAngleDeg float64, convertToEq isoLatitudeCircleToEq) string {
	return plotIsoLatitudeCircle(color, dashed, fixAngleDeg, 0.0, convertToEq)
}

func plotIsoLatitudeCircle(strokeColor string, dashed bool, fixAngleDeg float64, lat float64, convertToEq isoLatitudeCircleToEq) string {
	s := ""
	form1 := "        <path d=\"%s\" stroke=\"%s\" stroke-width=\"%.2f\" fill=\"none\" />\n"
	fixAngleR := fixAngleDeg * toRad
	latR := lat * toRad
	ra, de := convertToEq(0.0, latR, fixAngleR)
	x, y := eqToCartesianXY(ra*toDeg, de*toDeg)
	d := fmt.Sprintf("M%.1f,%.1f ", x, y)
	formContinual := "%.1f,%.1f "
	form0 := formContinual
	L := false
	c := ""
	if !dashed {
		d += "L"
	}
	for la := 1.0; la < 360.1; la = la + 1.0 {
		ra, de := convertToEq(la*toRad, latR, fixAngleR)
		x, y := eqToCartesianXY(ra*toDeg, de*toDeg)
		if dashed {
			if L {
				c = "M"
			} else {
				c = "L"
			}
			L = !L
		}
		d += fmt.Sprintf(c+form0, x, y)
	}
	strokeWidth := 0.25
	s += fmt.Sprintf(form1, d, strokeColor, strokeWidth)
	return s
}

func ecliplicalLongitudeToCartesianXY(ecLongitudeR float64, ecLatitudeR float64) (float64, float64) {
	𝜀Deg2025R := 𝜀Deg2025 * toRad
	ra, de := EclipticalToEquatorial(ecLongitudeR, ecLatitudeR, 𝜀Deg2025R)
	x, y := eqToCartesianXY(ra*toDeg, de*toDeg)
	return x, y
}
func plotPlatonYearDescription() string {
	speedByYear := 50.3 // angle seconds
	speedByYearDeg := speedByYear / 3600.0
	strokeColor := "brown"
	strokeWidth := 0.25
	form1 := "        <path d=\"M%.1f,%.1f L%.1f,%.1f \" stroke=\"%s\" stroke-width=\"%.2f\" fill=\"none\" />\n"
	form2 := `
	    <path id="%s" d="M%.1f,%.1f L%.1f,%.1f" stroke="none" fill="none" />
	    <text font-size="%.1f" font-family="Franklin Gothic, sans-serif" fill="%s" >
	      <textPath xlink:href="#%s" alignment-baseline="middle" text-anchor="start"  startOffset="1"> %d</textPath>
        </text>
`
	angle := 90.0
	sign := 1.0
	if Map.Latitude < 0 {
		sign = -1.0
		angle = -90.0
	}
	ecLat90R := sign * (90.0 - 𝜀Deg2025) * toRad
	ecLat89R := sign * (89.0 - 𝜀Deg2025) * toRad
	ecLat45R := sign * (45.0 - 𝜀Deg2025) * toRad

	s := ""
	for y := 1500; y > -23100; y = y - 500 {
		eclipticalLongitudeR := (angle + float64(2000-y)*speedByYearDeg) * toRad
		x1, y1 := ecliplicalLongitudeToCartesianXY(eclipticalLongitudeR, ecLat90R)
		x2, y2 := ecliplicalLongitudeToCartesianXY(eclipticalLongitudeR, ecLat89R)
		x3, y3 := ecliplicalLongitudeToCartesianXY(eclipticalLongitudeR, ecLat45R)
		id := fmt.Sprintf("PlatonY%d", y)
		s += fmt.Sprintf(form1, x1, y1, x2, y2, strokeColor, strokeWidth)
		if y%1000 == 0 {
			s += fmt.Sprintf(form2, id, x2, y2, x3, y3, 3.0, strokeColor, id, y)
		}
	}

	return s
}
func plotEcliptic() string {
	s := "      <g id=\"plotEcliptic\">\n"
	s += plotGreatCircle(Map.Colors.Ecliptic, Map.DashedEcliptic, 𝜀Deg2025, EclipticalToEquatorial)
	s += "      </g>\n"
	return s
}

func plotHorizon() string {
	s := "      <g id=\"plotHorizon\" >\n"
	geographicLatitude := Map.Latitude
	s += plotGreatCircle(Map.Colors.Horizon, Map.DashedHorizon, geographicLatitude, AzimutalToEquatoreal_I)
	s += "      </g>\n"
	return s
}

func plotAlmucantarats() string {
	s := "      <g id=\"plotAlmucantarats\" >\n"
	hInc := 10.0
	for h := hInc; h < 90.0; h = h + hInc {
		s += plotIsoLatitudeCircle(Map.Colors.Horizon, Map.DashedHorizon, Map.Latitude, h, AzimutalToEquatoreal_I) // plot Almucantarat
	}
	s += "      </g>\n"
	return s
}

// Platon ecliptic move 50" per year ~ 25920 years
func plotPlatonYear() string {
	s := "      <g id=\"plotPlatonYear\"  >\n"
	eclLatitude := 90.0 - 𝜀Deg2025
	if Map.Latitude < 0.0 {
		eclLatitude *= -1.0
	}
	s += plotIsoLatitudeCircle(Map.Colors.Ecliptic, Map.DashedEcliptic, 𝜀Deg2025, eclLatitude, EclipticalToEquatorial)
	s += plotPlatonYearDescription()
	s += "      </g>\n"
	return s
}

func plotMeridians() string {
	s := "      <g id=\"plotMeridians\">\n"
	aInc := 10.0
	for a := aInc; a < 360.1; a = a + aInc {
		s += plotMeridian(Map.Colors.Horizon, Map.DashedHorizon, Map.Latitude, a)
	}
	s += "      </g>\n"
	return s
}

func plotMeridian(color string, dashed bool, fixAngleDeg float64, a float64) string {
	form1 := "        <path id=\"MerA%.f\" d=\"%s\" stroke=\"%s\" stroke-width=\"0.25\" fill=\"none\" />\n"
	fixAngleR := fixAngleDeg * toRad
	aR := a * toRad
	ra, de := AzimutalToEquatoreal_I(aR, 0.0, fixAngleR)
	x, y := eqToCartesianXY(ra*toDeg, de*toDeg)
	d := fmt.Sprintf("M%.1f,%.1f ", x, y)
	formContinual := "%.1f,%.1f "
	form0 := formContinual
	L := false
	c := ""
	if !dashed {
		d += "L"
	}
	hMax := 80.1
	if int(a)%90 == 0 {
		hMax = 90.1
	}
	for h := 0.0; h < hMax; h = h + 1.0 {
		hR := h * toRad
		ra, de := AzimutalToEquatoreal_I(aR, hR, fixAngleR)
		x, y := eqToCartesianXY(ra*toDeg, de*toDeg)
		if dashed {
			if L {
				c = "M"
			} else {
				c = "L"
			}
			L = !L
		}
		d += fmt.Sprintf(c+form0, x, y)
	}
	s := fmt.Sprintf(form1, a, d, color)

	return s
}

// Enhanced Character Separated Values table format https://docs.astropy.org/en/stable/io/ascii/ecsv.html
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
func LoadObjects(filename string) {
	if bytes, err := os.ReadFile(filename); err != nil {
		fmt.Printf("Error loading file %s: %s\n", filename, err.Error())
		return
	} else {
		if err1 := json.Unmarshal([]byte(bytes), &SliceOfObjects); err1 != nil {
			fmt.Printf("Error unmarshaling content of the json file %s: %s\n", filename, err1.Error())
		}
	}
}
func LoadJsonFileSlice(filename string, SliceOfAnyStructs *any) {
	if bytes, err := os.ReadFile(filename); err != nil {
		fmt.Printf("Error loading file %s: %s\n", filename, err.Error())
		return
	} else {
		if err1 := json.Unmarshal([]byte(bytes), SliceOfAnyStructs); err1 != nil {
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

	fmt.Fprint(w, `<html>
  <head>
    <title>SkyMap</title>
	<script>
	  function openSkyMap(f){
		// f == this form
		// http://localhost:8080/img/svg/skymap/co/n44/2/x
		let u = '/img/svg/skymap/';
		u +=  f.color_style.value + '/'
		u +=  f.hemisphere.value + f.latitude.value + '/'
		u +=  f.paper.value + '/'
		u +=  f.draw.value 
		
		//alert('Hello from SkyMap submit! url='+u);		
		//window.location.href = u; // in the same tab

		const anchor = document.createElement("a");
		anchor.href = u;
		anchor.target = "_blank"; // Open in a new tab
		anchor.click();
	  }
	</script>	
  </head>
	 `)
	fmt.Fprint(w, "<body style=\"text-align: center;\">\n")
	fmt.Fprint(w, "<h1>SkyMap Lab select page</h1>\n")
	fmt.Fprint(w, `	
	<form action="javascript:;" onsubmit=" openSkyMap( this ) ">  
     <select name="hemisphere" id="hemisphere">
        <option value="s" >S</option>
        <option value="n" selected="selected">N</option>
     </select>

     <label for="latitude">Latitude:</label>
     <input type="number" id="latitude" name="latitude" value="44" step="1"  min="0" max="90" size="2">
	 
	 <br/>
	 <label for="color">Color:</label>
	 <input type="radio" id="co" name="color_style" value="co" checked="checked">
	 <label for="bw">Black &amp; White</label>
	 
	 <input type="radio" id="bw" name="color_style" value="bw">
     
    <br/>
     <select name="paper" id="paper" title="paper">
        <option value="0" title="297x210">A4</option>
        <option value="1" title="420x297" >A3</option>
        <option value="2" title="215.9x279.4" selected="selected">Letter 8.5x11</option>
        <option value="3" title="215.9x355.6"  >Legal 8.5x14</option>
        <option value="4" title="279.4x431.8" >Ledger 11x17</option>
     </select>
     <select name="draw" id="draw" title="draw">
        <option value="0" selected="selected">Map + Platon Year</option>
        <option value="1">AZ grid</option>
        <option value="2">Map only</option>
        <option value="3">All</option>
     </select>
	 <br/>
	 <br/>
	 <input type="reset" value="RESET">
	 <input type="submit" value="SUBMIT">

    </form>
	<!-- 
	https://github.com/kpawlik/svg2pdf/
	https://pkg.go.dev/github.com/nicholasblaskey/svg-rasterizer#section-readme
	https://helpx.adobe.com/acrobat/kb/print-posters-banners-acrobat-reader.html
	Poster: TileScale, OverLap 0.005 in, Cut marks
	Orientation: Portrate Landscape 
  -->

	`)
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
func intId0(idStr string) int {
	var idInt int64
	var err error
	if idInt, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		idInt = 0
	}
	return int(idInt)
}
func HandlerSkyMapGeneral(w http.ResponseWriter, r *http.Request) {
	colorId := r.PathValue("colorId")
	colorfullMap := strings.HasPrefix(colorId, "co")
	lat := getLatitude(r.PathValue("latId"))
	if colorfullMap {
		SetMapStyle(249.0, lat, MapColorsRed)
	} else {
		SetMapStyle(249.0, lat, MapBlackAndWhite)
	}
	paperIdInt := intId0(r.PathValue("paperId"))
	drawIdInt := intId0(r.PathValue("drawId"))

	w.Header().Set("Content-Type", "image/svg+xml")
	defs := plotRaCross()
	defs += plotAxisDeclinations()
	defs += plotRaHourRoundScale()
	defs += plotDateRoundScale()
	defs += plotConstellations()
	defs += plotConstellationNames()
	defs += plotOuterCircle()
	defs += plotEcliptic()
	defs += plotHorizon()
	defs += plotAlmucantarats()
	defs += plotMeridians()
	defs += plotPlatonYear()
	defs += plotDirectionsOfTheApparentRotationOfTheSky()
	defs += plotStars()
	defs += plotObjects()

	var draws = []string{"draw_platonYear_map", "draw_AZ_grid", "draw_map", "draw_all"}
	draw := draws[drawIdInt]

	svgTemplate2 := fmt.Sprintf(svgTemplate1, defs, draw)
	if t, err := template.New("SkyMap").Parse(svgTemplate2); err == nil {
		data := getSvgData(colorfullMap, paperIdInt)
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
	//w.WriteHeader(http.StatusOK)
}
