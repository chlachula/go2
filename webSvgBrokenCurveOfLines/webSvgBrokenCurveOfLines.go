package webSvgBrokenCurveOfLines

import (
	"fmt"
	"math"
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

type Point = struct {
	X float64
	Y float64
}

// slope and y-intercept
func lineSlopeYintercept(p1, p2 Point) (float64, float64) {
	k := (p2.Y - p1.Y) / (p2.X - p1.X)
	q := p1.Y - k*p1.X
	return k, q
}

// the perpendicular of the line at point p - slope and y-intercept
func perpendicularSlopeYintercept(k float64, p Point) (float64, float64) {
	k2 := -1 / k
	q2 := p.Y - k2*p.X
	return k2, q2
}

// oposite points on the line with slope k, y-intercept q throught point p
func opositePointsOnLineThroughPointInDistance(k, q float64, p Point, distance float64) (Point, Point) {
	var a, b Point
	xk2 := math.Sqrt(1.0 - k*k)
	a.Y = p.Y + k*distance
	a.X = p.X + distance*xk2
	b.Y = p.Y - k*distance
	b.X = p.X - distance*xk2
	return a, b
}
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
	points := []Point{{100, 100}, {200, 200}, {300, 100}, {400.0, 455.6}, {600.0, 432.78}}
	lines := ""
	width := 30.0
	w2 := width / 2.0
	for i := 1; i < len(points); i++ {
		p0 := points[i-1]
		p1 := points[i]
		line := fmt.Sprintf("<path d=\"M%.1f,%.1f L%.1f,%.1f\" stroke=\"red\" stroke-width=\"4\"/>\n", p0.X, p0.Y, p1.X, p1.Y)
		lines += line
		p0k, _ := lineSlopeYintercept(p0, p1)
		p0k2, p0q2 := perpendicularSlopeYintercept(p0k, p0)
		p1k2, p1q2 := perpendicularSlopeYintercept(p0k, p1)

		a0, b0 := opositePointsOnLineThroughPointInDistance(p0k2, p0q2, p0, w2)
		a1, b1 := opositePointsOnLineThroughPointInDistance(p1k2, p1q2, p1, w2)
		lines += fmt.Sprintf("<path d=\"M%.1f,%.1f L%.1f,%.1f %.1f,%.1f %.1f,%.1f z\" fill=\"none\" stroke=\"green\" stroke-width=\"4\"/>\n", a0.X, a0.Y, b0.X, b0.Y, b1.X, b1.Y, a1.X, a1.Y)

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
