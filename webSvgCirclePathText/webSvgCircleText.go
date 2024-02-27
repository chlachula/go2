package webSvgCirclePathText

import (
	"fmt"
	"math"
	"net/http"
	"text/template"
)

const (
	htmlHead0 = `<html><head><title>SVG round logo: %s</title>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<link rel="icon" type="image/ico" href="favicon.ico">
</head>
<body style="text-align: center;">

`
	htmlEnd      = "\n<br/></body></html>"
	svgTemplate1 = `
<svg xmlns="http://www.w3.org/2000/svg" 
    xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="-250 -250 500 500">
    <title>Red Hot Chilli Peppers Logo http://thenewcode.com/482/Placing-Text-on-a-Circle-with-SVG </title>
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
</style>
</defs> 
<!--
<path d="M0,0 m-100,0 a100,100 0 1,1  200,0 " style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 1"/>
<path d="M0,0 m-100,0 a100,100 0 0,0  200,0 " style="fill:none;fill-opacity: 1;stroke:green;stroke-width: 1"/>
<path id="linT"    d="M0,0 m-{{.Tx}},{{.Ty}} l{{.Tx}},-{{.Ty}} {{.Tx}},{{.Ty}}" style="fill:none;fill-opacity: 1;stroke:black;stroke-width: 0.5"/>
<path id="linB" d="M0,0 m-{{.Bx}},{{.By}} l{{.Bx}},-{{.By}} {{.Bx}},{{.By}}" style="fill:none;fill-opacity: 1;stroke:black;stroke-width: 0.5"/>
-->
 <path id="relT" d="M0,0 m-{{.Tx}},{{.Ty}} a{{.R1}},{{.R1}} 0 1,1  {{.Tx2}},0 " style="fill:none;fill-opacity: 1;stroke:pink;stroke-width: 10.5"/>
 <path id="relB" d="M0,0 m-{{.Bx}},{{.By}} a{{.R1}},{{.R1}} 0 0,0  {{.Bx2}},0 " style="fill:none;fill-opacity: 1;stroke:yellow;stroke-width: 10.5"/>
 <!--text 
  text-anchor="start|middle|end" 
  lengthAdjust="spacing|spacingAndGlyphs"
  alignment-baseline="auto|baseline|before-edge|text-before-edge|middle|central|after-edge|text-after-edge|ideographic|alphabetic|hanging|mathematical|top|center|bottom"
  dominant-baseline="auto | text-bottom | alphabetic | ideographic | middle | central | mathematical | hanging | text-top"
  ></text-->

  <circle cx="0" cy="0" r="{{.RingRadius}}" stroke="{{.RingColor}}" stroke-width="{{.RingWidth}}" fill="none" />

  <text dy="{{.Dy1}}" dx="{{.Dx1}}" textLength="{{.Tlen}}" dominant-baseline="hanging" class="font1 upFont">
      <textPath xlink:href="#relT" >{{.UpperText}}</textPath>
  </text>    
  <text dy="{{.Dy2}}" dx="{{.Dx2}}" textLength="{{.Blen}}"  class="font1 downFont">
      <textPath xlink:href="#relB">{{.BottomText}}</textPath>
  </text>
  <circle cx="{{.Qx}}" cy="{{.Qy}}" r="{{.Qr}}" stroke="none" stroke-width="0" fill="black" />
  <circle cx="-{{.Qx}}" cy="{{.Qy}}" r="{{.Qr}}" stroke="none" stroke-width="0" fill="black" />
  
    <circle cx="0" cy="0" r="{{.R0}}" stroke="black" stroke-width="0.5" fill="none" />
	<circle cx="0" cy="0" r="{{.RupperDown}}" stroke="black" stroke-width="0.5" fill="none" />
	<circle cx="0" cy="0" r="{{.RbottomTop}}" stroke="black" stroke-width="0.5" fill="none" />
	<circle cx="0" cy="0" r="{{.R1}}" stroke="black" stroke-width="0.5" fill="none" />

</svg>
`
)

type SvgDataType = struct {
	RingColor   string
	TopColor    string
	BottomColor string
	UpperText   string
	BottomText  string

	RingRadius float64
	RingWidth  float64
	RbottomTop float64
	RupperDown float64
	R0         float64
	R1         float64
	Dy1        float64
	Dx1        float64
	Tlen       float64
	Dy2        float64
	Dx2        float64
	Blen       float64
	Tx         float64
	Tx2        float64
	Ty         float64
	Bx         float64
	Bx2        float64
	By         float64
	Qx         float64
	Qy         float64
	Qr         float64
	FontSize   float64
}

var (
	TopText    string
	BottomText string
)

func SetVariables(top, bottom string) {
	TopText = top
	BottomText = bottom
	fmt.Printf("TOP:    %s\nBOTTOM: %s\n", TopText, BottomText)
}

func getSvgData(color bool) SvgDataType {
	topAngle := 260.0
	botAngle := 76.0
	tA := topAngle * math.Pi / 360.0
	bA := botAngle * math.Pi / 360.0
	qA := ((360-topAngle-botAngle)*0.5 + botAngle) * math.Pi / 360.0
	r1 := 120.0
	ringRadius := r1 * 170.0 / 200.0
	data := SvgDataType{
		RingColor:   "lightblue",
		TopColor:    "green",
		BottomColor: "red",
		UpperText:   TopText,
		BottomText:  BottomText,
		RingRadius:  ringRadius,
		RingWidth:   70.0 / 200.0 * r1,
		R0:          100.0,
		RupperDown:  141.0,
		RbottomTop:  161.0,
		R1:          r1,
		Dy1:         0,
		Dx1:         0,
		Dy2:         0,
		Dx2:         0,
		Tlen:        r1 * 2.0 * tA,
		Tx:          r1 * math.Sin(tA),
		Tx2:         2.0 * r1 * math.Sin(tA),
		Ty:          -r1 * math.Cos(tA),
		Blen:        r1 * 2.0 * bA,
		Bx:          r1 * math.Sin(bA),
		Bx2:         2.0 * r1 * math.Sin(bA),
		By:          r1 * math.Cos(bA),
		Qx:          ringRadius * math.Sin(qA),
		Qy:          ringRadius * math.Cos(qA),
		Qr:          r1 * 0.03,
		FontSize:    59.0 / 200.0 * r1,
	}
	if !color {
		data.RingColor = "lightgray"
		data.TopColor = "black"
		data.BottomColor = "darkgray"
	}
	return data
}
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	writeHtmlHeadAndMenu(w, "/", "Home")
	fmt.Fprint(w, `<div class="container70percentFlex">
	<div>
	<a name="about"></a><h2>About</h2>
	<p>SVG round logo color and black &amp; white
	was developed by 
	<a href="https://github.com/chlachula">Josef Chlachula</a> 
	under <a href="https://github.com/chlachula/go2/tree/main?tab=MIT-1-ov-file#readme">MIT license</a>.
	</p>
    <a name=\"contact\"></a>
	<h2>Contact</h2>
	<p>... will be provided here in future ...</p>
	</div>
</div>
	`)

	fmt.Fprint(w, htmlEnd)
}
func HandlerSvgRoundLogoColor(w http.ResponseWriter, r *http.Request) {
	writeHtmlHeadAndMenu(w, "/svg-roundlogo-color", "Color")

	if t, err := template.New("SvgRoundLogoColor").Parse(svgTemplate1); err == nil {
		data := getSvgData(true)
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}

	fmt.Fprint(w, htmlEnd)
}
func HandlerSvgRoundLogoBlackWhite(w http.ResponseWriter, r *http.Request) {
	writeHtmlHeadAndMenu(w, "/svg-roundlogo-bw", "B&amp;W")

	if t, err := template.New("SvgRoundLogoBlackWhite").Parse(svgTemplate1); err == nil {
		data := getSvgData(false)
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}

	fmt.Fprint(w, htmlEnd)
}
