package svgSpoonLogo

import (
	"fmt"
	"net/http"
	"text/template"
)

type SvgDataType = struct {
	Cbase   string
	Cgrad1  string
	Cgrad2  string
	Cground string
	Cring   string
	Cray    string
	Cspoon  string
	Ctop    string
	Cbot    string
	Chor    string
}

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

const svg1template = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
viewBox="-101 -101 202 202" height="202" width="202" 
     style="background:none">
    >
 <description>NCRAL 2025 logo, Josef Chlachula, 2024-12-03, MIT license</description>
 <defs>
     <linearGradient id="grad1" x1="0%" x2="0%" y1="0%" y2="100%">
      <stop offset="0%" stop-color="{{.Cgrad1}}" />
      <stop offset="100%" stop-color="{{.Cgrad2}}" />
    </linearGradient>
   <style>
     .base{fill:{{.Cbase}};fill-opacity: 1;stroke:none;stroke-width: 0.001}
     .ray {fill:none;fill-opacity: 1;stroke:{{.Cray}};stroke-width: 2.5}
     .spoon {fill:none;fill-opacity: 1;stroke:{{.Cspoon}};stroke-width: 5}
     .nothing {fill:none;fill-opacity: 1;stroke:none;stroke-width: 0.5}
     .roundFontTop {fill:{{.Ctop}};font-size:20px;}
     .roundFontBot {fill:{{.Cbot}};font-size:17px;}
     .horizonFont {fill:{{.Chor}};font-size:14px;}
   </style>
   
   <g id="spoon_and_rays">
   <path id="rel" d="M0,0 m-120,100 h400" style="fill:none;fill-opacity: 1;stroke:gray;stroke-width: 5"/>

   <path id="ray0"   d="M0,100 v-466" class="ray"/>
   <!--path id="ray10a" d="M0,50 l17.3648,48.4808 v-466"  style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray10b" d="M0,50 l-17.3648,48.4808 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/-->
   <path id="ray20a" d="M0,50 l34.202,43.9693 v-466"   class="ray"/>
   <path id="ray20b" d="M0,50 l-34.202,43.9693 v-466"  class="ray"/>
   <!--path id="ray30a" d="M0,50 l50,36.6025 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/>
   <path id="ray30b" d="M0,50 l-50,36.6025 v-466" style="fill:none;fill-opacity: 1;stroke:blue;stroke-width: 2"/-->
   <path id="ray40a" d="M0,50 l64.2788,26.6044  v-466"  class="ray"/>
   <path id="ray40b" d="M0,50 l-64.2788,26.6044 v-466"  class="ray"/>

   <!--path id="rel" d="M0,0 m-86.6025,50 a100,100 1 0,0 173.205,0   a25,25 0 0,1 28.1211,-11.6481 l380,60" style="fill:none;fill-opacity: 1;stroke:red;stroke-width: 5"/-->
   <path id="rel" d="M0,0 m-86.6025,50 a100,100 1 0,0 173.205,0   a25,25 0 0,1 28.1211,-11.6481 l190,30" class="spoon"/>
   </g>
   <g id="ring">
     <circle cx="0" cy="0" r="90" style="fill:none;fill-opacity: 1;stroke:{{.Cring}};stroke-width: 20"/>
     <circle cx="0" cy="0" r="100" style="fill:none;fill-opacity: 1;stroke:black;stroke-width: 0.5"/>
     <circle cx="0" cy="0" r="80" style="fill:none;fill-opacity: 1;stroke:black;stroke-width: 0.5"/>

    <path id="pathTop" d="M0,0 m-48.5,-84.0045 a97,97 0 0,1  97,0 " class="nothing"/>
    <path id="pathBot" d="M0,0 m-91.7153,-31.5801 a97,97 0 1,0  183.4306,0  " class="nothing"/>
    <text dy="0" dx="0" textLength="104.0" dominant-baseline="hanging" >
      <textPath xlink:href="#pathTop"  class="roundFontTop">NCRAL 2025</textPath>
    </text>    
    <text dy="0" dx="0" textLength="369.1">
      <textPath xlink:href="#pathBot" class="roundFontBot">* FIRST LIGHT * NEW FRONTIERS * NEW PEOPLE *</textPath>
    </text>
   </g>

   <g id="texts"  class="horizonFont">
      <text x="0" y="00" >Conference</text>
      <text x="0" y="20" >April 25/26</text>
      <text x="0" y="40" >Minneapolis</text>
   </g>

 </defs>
    <!--circle cx="0" cy="0" r="100" class="base"/-->
    <circle cx="0" cy="0" r="100" fill="url(#grad1)"/>
    
    <rect width="110" height="25" x="-55" y="60" fill="{{.Cground}}" />
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
  Logo variations 
  <br/>
  <br/>
  <br/>
  <a href="/img/svg/roundLogo_BW"><img src="/img/svg/roundLogo_BW" width="200"/><a>

  <a href="/img/svg/roundLogo_Color"><img src="/img/svg/roundLogo_Color" width="200"/></a>

  <a href="/img/svg/roundLogo_Color2"><img src="/img/svg/roundLogo_Color2" width="200"/></a>
  
  

 </body>
</html>
	`)
}
func HandlerHtmlRoundLogo1(w http.ResponseWriter, r *http.Request) {
	//	s := fmt.Sprintf(html1, svg1)
	s := fmt.Sprintf(html1, "svg1")
	fmt.Fprint(w, s)
}
func HandlerImgSvgRoundLogo_BW(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	data := SvgDataType{
		Cbase:   "white",
		Cgrad1:  "white",
		Cgrad2:  "white",
		Cground: "white",
		Cring:   "white",
		Cray:    "black",
		Cspoon:  "black",
		Ctop:    "black",
		Cbot:    "black",
		Chor:    "black",
	}
	if t, err := template.New("LogoBW").Parse(svg1template); err == nil {
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
}
func HandlerImgSvgRoundLogo_Color(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	data := SvgDataType{
		Cbase:   "darkblue",
		Cgrad1:  "yellow",
		Cgrad2:  "red",
		Cground: "black",
		Cring:   "silver",
		Cray:    "yellow",
		Cspoon:  "darkred",
		Ctop:    "black",
		Cbot:    "black",
		Chor:    "white",
	}
	if t, err := template.New("LogoColor").Parse(svg1template); err == nil {
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
}
func HandlerImgSvgRoundLogo_Color2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	data := SvgDataType{
		Cbase: "darkblue",
		//		Cgrad1:  "darkblue", Cgrad2:  "yellow",
		Cgrad1: "black", Cgrad2: "lightblue",
		Cground: "black",
		Cring:   "silver",
		Cray:    "yellow",
		Cspoon:  "red",
		Ctop:    "black",
		Cbot:    "black",
		Chor:    "white",
	}
	if t, err := template.New("LogoColor").Parse(svg1template); err == nil {
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
}
