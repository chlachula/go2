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
 <defs>
   <g id="spoon_and_rays">
   <path id="rel" d="M0,0 m-120,100 h400" style="fill:none;fill-opacity: 1;stroke:gray;stroke-width: 5"/>

   <path id="ray0"   d="M0,100 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <!--path id="ray10a" d="M0,50 l17.3648,48.4808 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray10b" d="M0,50 l-17.3648,48.4808 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/-->
   <path id="ray20a" d="M0,50 l34.202,43.9693 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray20b" d="M0,50 l-34.202,43.9693 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <!--path id="ray30a" d="M0,50 l50,36.6025 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray30b" d="M0,50 l-50,36.6025 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/-->
   <path id="ray40a" d="M0,50 l64.2788,26.6044 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray40b" d="M0,50 l-64.2788,26.6044 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>

   <!--path id="rel" d="M0,0 m-86.6025,50 a100,100 1 0,0 173.205,0   a25,25 0 0,1 28.1211,-11.6481 l380,60" style="fill:none;fill-opacity: 1;stroke:red;stroke-width: 5"/-->
   <path id="rel" d="M0,0 m-86.6025,50 a100,100 1 0,0 173.205,0   a25,25 0 0,1 28.1211,-11.6481 l190,30" style="fill:none;fill-opacity: 1;stroke:red;stroke-width: 5"/>
   </g>
   <g id="ring">
     <circle cx="0" cy="0" r="90" style="fill:none;fill-opacity: 1;stroke:silver;stroke-width: 20"/>
     <circle cx="0" cy="0" r="100" style="fill:none;fill-opacity: 1;stroke:black;stroke-width: 0.5"/>
     <circle cx="0" cy="0" r="80" style="fill:none;fill-opacity: 1;stroke:black;stroke-width: 0.5"/>

    <path id="pathTop" d="M0,0 m-48.5,-84.0045 a97,97 0 0,1  97,0 " style="fill:none;fill-opacity: 1;stroke:none;stroke-width: 0.5"/>
    <path id="pathBot" d="M0,0 m-91.7153,-31.5801 a97,97 0 1,0  183.4306,0  " style="fill:none;fill-opacity: 1;stroke:none;stroke-width: 0.5"/>
    <text dy="0" dx="0" textLength="104.0" dominant-baseline="hanging" >
      <textPath xlink:href="#pathTop"  font-size="19" >NCRAL 2025</textPath>
    </text>    
    <text dy="0" dx="0" textLength="369.1">
      <textPath xlink:href="#pathBot">* FIRST LIGHT * NEW FRONTIERS * NEW PEOPLE *</textPath>
    </text>
   </g>

   <g id="texts"  font-size="14">
      <text x="0" y="00" >Conference</text>
      <text x="0" y="20" >April 25/26</text>
      <text x="0" y="40" >Minneapolis</text>
   </g>

 </defs>
    <circle cx="0" cy="0" r="100" style="fill:white;fill-opacity: 1;stroke:none;stroke-width: 0.001"/>
    <use xlink:href="#spoon_and_rays" transform="translate(-22, 30)  scale(0.3,0.3)" />
    <use xlink:href="#ring"  />
    <use xlink:href="#texts"  transform="translate(3, -19)"  />

  Sorry, your browser does not support inline SVG.
</svg>
`

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html>
 <head>
      <meta http-equiv="refresh" content="0; url=/svgImages">
 	  <title>Redirect to svgImages</title>
 </head>
 <body>
  <h1>Click to: <a href="/svgImages">svgImages</a></h1>
 </body>
</html>
	`)
}
func HandlerSvgImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html>
 <head>
 	  <title>SVG Logo images</title>
 </head>
 <body>
  <h1>SVG Images</h1>
  Logo inside of
  <a href="/svgImages/roundLogo1">html</a>
  and as a standalone svg
  <a href="/img/svg/roundLogo1">image</a>
 </body>
</html>
	`)
}
func HandlerHtmlRoundLogo1(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf(html1, svg1)
	fmt.Fprint(w, s)
}
func HandlerImgSvgRoundLogo1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, svg1)
}
