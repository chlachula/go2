package webSvgBrokenCurveOfLines

import (
	"fmt"
	"net/http"
)

var DataLine string

const URL = "/broken-curve-of-lines"
const Title = "Broken Curve of straight lines"
const htmlMain = `<html>
<head><title>%s</title></head>
<body style="text-align: center;">
<h1><a href="%s">%s</a></h1>
</body>
</html>
`

const (
	htmlHead = `<html><head><title>%s</title>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<link rel="icon" type="image/ico" href="favicon.ico">
</head>
<body style="text-align: center;">
<a href="%s">%s</a> %s
`
	htmlEnd      = "\n<br/></body></html>"
	svgTemplate2 = `
<svg xmlns="http://www.w3.org/2000/svg" 
     xmlns:xlink="http://www.w3.org/1999/xlink" 
	 viewBox="0 0 1600 900"
	 style="background-color:lightgrey" >
    <title>To show title</title>

<defs>
    <style>
	.font1 { 
		font-size: {{.FontSize}}px;
		font-family: Franklin Gothic, sans-serif;
		font-weight: 90; 		
		letter-spacing: 2px;
	}
	%s
    </style>
</defs> 
%s
</svg>
`
)

func HandlerSvgBrokenCurveOfLines(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHead, Title, "/", "home", Title)
	/*
		if t, err := template.New("webpage1").Parse(svgTemplate2); err == nil {
			//data := getSvgData(true)
			if err = t.Execute(w, data); err != nil {
				fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
			}
		}
	*/
	type Point = struct {
		X float64
		Y float64
	}
	points := []Point{{100, 100}, {200, 200}, {300, 100}, {400.0, 455.6}, {600.0, 432.78}}
	lines := ""
	for i := 1; i < len(points); i++ {
		p0 := points[i-1]
		p1 := points[i]
		line := fmt.Sprintf("<path d=\"M%.1f,%.1f L%.1f,%.1f\" stroke=\"red\" stroke-width=\"4\"/>", p0.X, p0.Y, p1.X, p1.Y)
		lines += line
	}
	defs := fmt.Sprintf("<g id=\"brokenCurve\">\n%s\n</g>\n", lines)
	use := `  <use xlink:href="#brokenCurve" /> 
		 `
	s := fmt.Sprintf(svgTemplate2, defs, use)
	fmt.Fprint(w, s)
	fmt.Fprint(w, htmlEnd)
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlMain, Title, URL, Title)
}
