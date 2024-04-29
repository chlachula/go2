package SkyMapLab

import (
	"fmt"
	"math"
	"net/http"
	"text/template"
)

const (
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

  <path id="relT" d="M0,0 m-{{.Tx}},{{.Ty}} a{{.R1}},{{.R1}} 0 1,1  {{.Tx2}},0 " style="fill:none;fill-opacity: 1;stroke:pink;stroke-width: 10.5"/>
  <path id="relB" d="M0,0 m-{{.Bx}},{{.By}} a{{.R1}},{{.R1}} 0 0,0  {{.Bx2}},0 " style="fill:none;fill-opacity: 1;stroke:yellow;stroke-width: 10.5"/>

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
