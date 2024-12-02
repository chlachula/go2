package svgSpoonLogo

import (
	"fmt"
	"net/http"
)

const html1 = `
<!DOCTYPE html>
<html>
<head>
<title>SVG spoon logo</title>

</head>
 
<body>
<h2>SVG spoon logo</h2>
%s
</body>
</html>`

const svg1 = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
viewBox="-200 -200 400 400" height="400" width="400" 
     style="background:beige">
    >

  <path id="arc" d="M0,0 m-70.7,70.7 a100,100 0 1,1  141.4,0 " style="fill:none;stroke:orange;stroke-width: 1"></path>
  <text style="font-size:26px;">
  <textPath href="#arc" text-anchor="start" fill="green" >Left ~ start</textPath>
  <textPath href="#arc" startoffset="150" text-anchor="middle" fill="blue" >middle</textPath>
  <textPath href="#arc" startoffset="100%" text-anchor="end" fill="red" >Right ~ end</textPath>
  </text>

  Sorry, your browser does not support inline SVG.
</svg>
`

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html>
 <head>
      <meta http-equiv="refresh" content="0; url=/svgExamples">
 	  <title>Redirect to svgExamples</title>
 </head>
 <body>
  <h1>Click to: <a href="/svgExamples">svgExamples</a></h1>
 </body>
</html>
	`)
}
func HandlerSvgExamples(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html>
 <head>
 	  <title>Svg Examples</title>
 </head>
 <body>
  <h1>Svg Examples</h1>
  textPath element inside of
  <a href="/svgExamples/textPath1">html</a>
  and as a standalone svg
  <a href="/img/svg/textPath1">image</a>
 </body>
</html>
	`)
}
func HandlerHtmlTextPath1(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf(html1, svg1)
	fmt.Fprint(w, s)
}
func HandlerImgSvgTextPath1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, svg1)
}
