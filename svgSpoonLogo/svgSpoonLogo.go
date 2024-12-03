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
viewBox="-200 -200 800 400" height="400" width="800" 
     style="background:beige">
    >

  <path id="arc" d="M0,0 m-70.7,70.7 a100,100 0 1,1  141.4,0 " style="fill:none;stroke:orange;stroke-width: 1"></path>
  <text style="font-size:26px;">
    <textPath href="#arc" text-anchor="start" fill="green" >Left ~ start</textPath>
    <textPath href="#arc" startoffset="150" text-anchor="middle" fill="blue" >middle</textPath>
    <textPath href="#arc" startoffset="100%" text-anchor="end" fill="red" >Right ~ end</textPath>
  </text>

   <path id="rel" d="M0,0 m-100,100 h600" style="fill:none;fill-opacity: 1;stroke:silver;stroke-width: 5"/>

   <!--path id="rel" d="M0,0 m86.6025,50 a100,100 0 0,1  -173.205,0  a25,25 0 0,1  28.1211,-11.6481" style="fill:none;fill-opacity: 1;stroke:lightgreen;stroke-width: 5"/-->

   <path id="ray0"   d="M0,100 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray10a" d="M0,50 l17.3648,48.4808 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray10b" d="M0,50 l-17.3648,48.4808 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray20a" d="M0,50 l34.202,43.9693 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray20b" d="M0,50 l-34.202,43.9693 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray30a" d="M0,50 l50,36.6025 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray30b" d="M0,50 l-50,36.6025 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray40a" d="M0,50 l64.2788,26.6044 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray40b" d="M0,50 l-64.2788,26.6044 v-300" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>

   <path id="rel" d="M0,0 m-86.6025,50 a100,100 1 0,0 173.205,0   a25,25 0 0,1 28.1211,-11.6481 l380,60" style="fill:none;fill-opacity: 1;stroke:red;stroke-width: 5"/>


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
