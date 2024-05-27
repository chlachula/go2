package svgImagesJs

import "fmt"

var form1 = `<path id="lowerhalf_%s" d="M-100,0 a100,100 0 0,0 200,00" />
<g id="%s">
  <circle id="c1" cx="0" cy="0" r="100" stroke="gray" stroke-width="1" stroke-dasharray="5,5" d="M5 20 l215 0"  fill="none" />
  <text id="t1" x="0"  y="-75" text-anchor="middle" alignment-baseline="hanging"  style="font-family:Arial;font-size: 72;stroke:black;fill:black;">%s</text>
  <text id="t2" x="0"  y="0"   text-anchor="middle" alignment-baseline="baseline" style="font-family:Arial;font-size: 24;stroke:black;fill:black;">%s</text>
  <text id="t3" x="-67" y="67" text-anchor="start" alignment-baseline="baseline"  style="font-family:Arial;font-size: 32;stroke:black;fill:black;">%s</text>
  <text id="t4" x="67"  y="67" text-anchor="end" alignment-baseline="baseline"    style="font-family:Arial;font-size: 32;stroke:black;fill:black;">%s</text>
  <text id="t5" x="0" y="0" style="stroke: black;font-family: Arial;font-size: 11;" text-anchor="middle" alignment-baseline="baseline">
        <textPath id="t5path" xlink:href="#lowerhalf_%s" startOffset="50%">%s</textPath>
  </text>
</g>
`

func eyepieceLabel(id string, focus, units, magnification, viewAngle, bottomText string) string {
	s := fmt.Sprintf(form1, id, id, id, focus, units, magnification, viewAngle, bottomText)
	return s
}

func LabelsPage() {
	//   <use id="e1a" transform="translate(400, 400)  scale(7.5,7.5)" xlink:href="#k1a"></use><!-- asi 53mm -->
	eyepieceLabel("k1a", "32", "mm", "162x", "32'", "a b c d e f g h i j k l m")
}
