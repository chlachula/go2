package webSvgCirclePathText

type SvgSmartphoneDataType = struct {
	Width  float64
	Height float64
	W0     float64
	H0     float64
	W1     float64
	H1     float64
	W2     float64
	H2     float64
	X2     float64
	Y2     float64
	R      float64
	D      float64
	W3R    float64
	H3R    float64
	W4D    float64
	H4D    float64
	MicY   float64
	MicW   float64
	MicH   float64
	CamX   float64
	CamY   float64
	CamR   float64
}

const svgSmartphoneTemplate1 = `
<svg xmlns="http://www.w3.org/2000/svg" 
     xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Width}}" height="{{.Height}}" viewBox="-{{.W0}} -{{.H0}} {{.Width}} {{.Height}}" >
<title>Smartphone</title>
<defs>
  <g id="smartphone">
  <rect x="-{{.W0}}" y="-{{.H0}}"  width="{{.Width}}" height="{{.H2}}" fill="#F3F3F3" />
  <rect x="-{{.W0}}" y="{{.Y2}}"   width="{{.Width}}" height="{{.H2}}" fill="#F3F3F3" />

  <rect x="-{{.W0}}" y="-{{.H0}}"  width="{{.W2}}" height="{{.Height}}" fill="#F3F3F3" />
  <rect x="{{.X2}}" y="-{{.H0}}"  width="{{.W2}}" height="{{.Height}}" fill="#F3F3F3" />
  <path stroke="none" stroke-width="0" fill="black"
   d="M0,-{{.H3R}} 
      h{{.W1}} a{{.R}},{{.R}} 0 0,1 {{.R}},{{.R}}
	  v{{.H1}} v{{.H1}} a{{.R}},{{.R}} 0 0,1 -{{.R}},{{.R}}
	  h-{{.W1}} h-{{.W1}} a{{.R}},{{.R}} 0 0,1 -{{.R}},-{{.R}}
	  v-{{.H1}} v-{{.H1}} a{{.R}},{{.R}} 0 0,1 {{.R}},-{{.R}}
	  h{{.W1}} 
	  v{{.R}} v{{.D}}
	  h-{{.W4D}} v{{.H4D}} v{{.H4D}} h{{.W4D}} h{{.W4D}} v-{{.H4D}} v-{{.H4D}} h-{{.W4D}}
	  z
	  "
   />
   
   <path id="Mic_on_top" stroke="none" stroke-width="0" fill="gray"
   d="M-{{.MicW}},-{{.MicY}} h{{.MicW}} {{.MicW}} v{{.MicH}}  h-{{.MicW}} -{{.MicW}} z"/>

   <path id="Button_on_side1" stroke="none" stroke-width="0" fill="black"
   d="M{{.W3R}},-{{.H1}} m0,{{.MicW}} 0,{{.MicW}} h{{.MicH}} v{{.MicW}} h-{{.MicH}} z"/>
   <path id="Button_on_side2" stroke="none" stroke-width="0" fill="black"
   d="M{{.W3R}},-{{.H1}} m0,{{.MicW}} 0,{{.MicW}} 0,{{.MicW}} 0,{{.MicW}} h{{.MicH}} v{{.MicW}} h-{{.MicH}} z"/>

   <circle id="camera" r="{{.CamR}}" cx="{{.CamX}}" cy="{{.CamY}}" fill="black" stroke="darkgray" stroke-width="0.5" />
  </g>
</defs> 
<circle r="55" cx="0" cy="0" fill="red" stroke="green" stroke-width="6" />
<circle r="25" cx="0" cy="0" fill="orange" stroke="blue" stroke-width="4" />
  <use xlink:href="#smartphone" />
</svg>
`

func getSmartphoneData(width, height int) SvgSmartphoneDataType {
	f := 0.8
	var d SvgSmartphoneDataType
	widthF := float64(width)
	heightF := float64(height)
	d.Width = widthF
	d.Height = heightF
	d.W0 = widthF * 0.5
	d.H0 = heightF * 0.5
	d.W1 = d.W0 * f
	d.H1 = d.H0 * f
	d.W2 = d.W0 - d.W1
	d.H2 = d.H0 - d.H1
	d.X2 = d.W0 - d.W2
	d.Y2 = d.H0 - d.H2
	d.R = d.W0 * f * 0.1
	d.D = d.W0 * f * 0.02
	d.W3R = d.W1 + d.R
	d.H3R = d.H1 + d.R
	d.W4D = d.W1 - d.D
	d.H4D = d.H1 - d.D
	d.MicY = d.H3R
	d.MicW = d.W1 * 0.2
	d.MicH = d.R * 0.2
	d.CamX = -d.W4D + 5*d.D
	d.CamY = -d.H4D + 5*d.D
	d.CamR = d.W0 * f * 0.04
	return d
}
