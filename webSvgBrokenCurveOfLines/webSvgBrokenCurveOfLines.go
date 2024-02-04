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
	  >
    <title>To show title</title>

<defs>
    <style>
	<linearGradient id="grad1" x1="0%%" y1="0%%" x2="100%%" y2="0%%">
	  <stop offset="0%%" style="stop-color:rgb(255,255,0);stop-opacity:1" />
	  <stop offset="100%%" style="stop-color:rgb(255,0,0);stop-opacity:1" />
     </linearGradient>
	 <linearGradient id="verticalGrad1" x1="0%%" y1="0%%" x2="0%%" y2="100%%">
	   <stop offset="0%%" style="stop-color:rgb(199,255,199;stop-opacity:1" />
	   <stop offset="100%%" style="stop-color:rgb(089,089,089);stop-opacity:1" />
	 </linearGradient>
	 .font1 { 
		font-size: {{.FontSize}}px;
		font-family: Franklin Gothic, sans-serif;
		font-size: 45; 		
		font-weight: 9; 		
		letter-spacing: 2px;
	 }
	 .f1blue {
		fill: blue;
     }
    </style>
	%s
    <g id="LShape" >
	 <g transform="translate(350,200)">
	 <circle r="30" cx="0" cy="0" stroke="green" stroke-width="5" fill="none" />
	 <circle r="40" cx="0" cy="0" stroke="green" stroke-width="4.1" fill="none" />

	 <path stroke-width="20" stroke="red" fill="none" d="M0,0 v-100" />
	 <path stroke-width="20" stroke="yellow" fill="none" d="M0,0 h172" />
	 <path stroke-width="20" stroke="green" fill="none" d="M172,0 h207" />
	 <circle r="20" cx="0" cy="0" stroke="green" stroke-width="0" fill="green" />
	 <text x="30" y="-70" class="font1 f1blue">Inspire others</text>
	 <text x="30" y="-27" class="font1 f1blue">look &amp; feel good</text>
	 </g>
	</g>
</defs> 

  <rect width="1600" height="900" x="0" y="0" rx="0" ry="0" fill="url(#verticalGrad1)" />
%s
  <use xlink:href="#LShape" /> 
  <rect width="400" height="300" x="500" y="300" rx="10" ry="10" fill="pink" stroke="silver" stroke-width="2" />
  <rect width="400" height="300" x="600" y="400" rx="10" ry="10" fill="url(#verticalGrad1)" stroke="silver" stroke-width="2" />
  
  <ellipse cx="300" cy="500" rx="85" ry="55" fill="url(#grad1)"  stroke="black" stroke-width="2" />
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
	dx := math.Sqrt(distance * distance / (1.0 - k*k))
	dy := k * dx
	//	xk2 := math.Sqrt(distance*distance/(1.0 - k*k))
	//fmt.Printf("oposite k=%.6f xk2=%.6f \n", k, xk2)
	/*	a.Y = p.Y + k*distance
		a.X = p.X + distance*xk2
		b.Y = p.Y - k*distance
		b.X = p.X - distance*xk2
	*/
	a.Y = p.Y + dy
	a.X = p.X + dx
	b.Y = p.Y - dy
	b.X = p.X - dx
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
		fmt.Printf("%d. points a0=%v,b0=%v,a1=%v,b1=%v\n", i, a0, b0, a1, b1)
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
