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
	 .cross {
		stroke:black;
		stroke-width:0.5
	 }
    </style>
  %s
 </defs> 

  <g id="draw">
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

func SetVariables(top, bottom string) {
	TopText = top
	BottomText = bottom
	fmt.Printf("TOP:    %s\nBOTTOM: %s\n", TopText, BottomText)
}

func getSvgData(color bool) SvgDataType {
	data := SvgDataType{
		TopColor:    "green",
		BottomColor: "red",
		FontSize:    1.5,
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
	</g>
`
	return str
}
func raHourRoundScale() string {
	r1 := 150.0
	r2 := 154.0

	f0 := `
	<g id="raHourScale">
	  <circle cx="0" cy="0" r="150" stroke="black" stroke-width="0.5" fill="none" />
	  <circle cx="0" cy="0" r="152" stroke="black" stroke-width="0.5" fill="none" />
	  %s
	</g>
`
	s := "\n"
	f1 := "      <line x1=\"%.1f\" y1=\"%.1f\" x2=\"%.1f\" y2=\"%.1f\" class=\"cross\" />\n"
	for ra := 0; ra <= 23; ra++ {
		a := float64(ra*15) * math.Pi / 180.0
		x1 := -r1 * math.Sin(a)
		y1 := r1 * math.Cos(a)
		x2 := -r2 * math.Sin(a)
		y2 := r2 * math.Cos(a)
		s += fmt.Sprintf(f1, x1, y1, x2, y2)
	}
	return fmt.Sprintf(f0, s)
}
func dateRoundScale() string {
	str := ""
	return str
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
