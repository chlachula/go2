/*
Icons with moon phases are very often used on web calendars.
The synodic month is 29.5 days long, but for calendar purposes we can think of four weeks of 28 days.
Function CreateMoonPhaseSvgIcons generates 28 moon phase icons 90/7 = 12.9 degrees apart.
*/
package moonPhases

import (
	"fmt"
	"math"
	"os"
)

var svgBase = `<?xml version="1.0" encoding="utf-8"?>
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
     width="300" height="300" viewBox="-150 -150 300 300">
<desc>
  MIT License Josef Chlachula 2023
  Icon for Moon age %.2fdays (%.2fd..%.2fd) - angle  %.2f°(%.2f°..%.2f°)
</desc>	 
<style id="style1">
 .SunLight {
   fill: yellow;
   stroke: gold;
   stroke-width: 0.5;
 }
</style>
<defs>
 <g id="darkMoon"><circle cx="0" cy="0" r="100" /></g>
 <g id="m28f%02d"><use xlink:href="#darkMoon"/><path id="m02" d="m0,-100 a100,100 0 0,%d 0,200 a%.2f,100 0 0,%d 0,-200 z" class="SunLight"/></g>
</defs>

<use xlink:href="#m28f%02d" x="0" y="0" />

</svg>  
`
var htmlBase = `<html>
<title>All Phases</title>
<style>
 .center {
	margin-left: auto;
	margin-right: auto;
 }
 td {color:blue; text-align: center;}
</style>
<body bgcolor="lightblue">
 <table class="center">
	<caption>All 28 phases<br/>Synodic orbital period: 29.530589 d</caption>
	%s
 </table>
 </body>
</html>
`

const PhasesNumber = 28 //should be divisible by 4

func CreateAllPhasesPage() {
	tdFmt := "  <td>%.1f-%.1f°-%.1f<br/>%.1f - %.1f d<br/>%.1f d<br/>%d<br/><img src=\"%s\" width=\"150\" /></td>\n"
	text := ""
	numInQuarter := PhasesNumber / 4
	for m := 0; m < 4; m++ {
		text += "<tr>"
		for n := 0; n < numInQuarter; n++ {
			phaseNum := m*numInQuarter + n
			mAngleN, mAngle, mAngleP, mAgeN, mAge, mAgeP := moonAgePhaseNumbers(phaseNum)
			fname := fmt.Sprintf("moon28f%02d.svg", phaseNum)
			td := fmt.Sprintf(tdFmt, mAngleN, mAngle, mAngleP, mAgeN, mAgeP, mAge, phaseNum, fname)
			text += td
		}
		text += "</tr>\n"
	}
	page := fmt.Sprintf(htmlBase, text)
	filename := fmt.Sprintf("All%02dPhasesPage.html", PhasesNumber)
	createFile(filename, page)
}

func moonAgePhaseNumbers(n int) (mAngleN, mAngle, mAngleP, mAgeN, mAge, mAgeP float64) {
	synodicMoon := 29.530589
	a28half := 360.0 / float64(2*PhasesNumber)
	mAngle = a28half * float64(2*n)
	mAngleN = mAngle - a28half
	if mAngleN < 0.0 {
		mAngleN += 360.0
	}
	mAngleP = mAngle + a28half
	mAge = mAngle / 360.0 * synodicMoon
	mAgeN = mAngleN / 360.0 * synodicMoon
	mAgeP = mAngleP / 360.0 * synodicMoon
	return
}

// Create Moon phase SVG icon file in current directory
func CreateMoonPhaseSvgIcons() {
	for n := 0; n < PhasesNumber; n++ {
		createMoonPhaseSvgIcon(n)
	}
}

// Create Moon phase SVG icon file in current directory
func createMoonPhaseSvgIcon(n int) {
	if n < 0 || n > PhasesNumber {
		return
	}
	x := 100.0 * math.Cos(float64(n*2)*math.Pi/float64(PhasesNumber))
	HalfPhasesNumber := PhasesNumber / 2
	i1 := 0
	i2 := 0
	if n <= HalfPhasesNumber {
		i1 = 1
		if x < 0.0 {
			i2 = 1
		}
	} else {
		if x > 0.0 {
			i2 = 1
		}
	}
	/*
		synodicMoon := 29.530589
		a28half := 360.0 / float64(2*PhasesNumber)
		mAngle := a28half * float64(2*n)
		mAngleN := mAngle - a28half
		if mAngleN < 0.0 {
			mAngleN += 360.0
		}
		mAngleP := mAngle + a28half
		mAge := mAngle / 360.0 * synodicMoon
		mAgeN := mAngleN / 360.0 * synodicMoon
		mAgeP := mAngleP / 360.0 * synodicMoon
	*/
	mAngleN, mAngle, mAngleP, mAgeN, mAge, mAgeP := moonAgePhaseNumbers(n)
	text := fmt.Sprintf(svgBase, mAge, mAgeN, mAgeP, mAngle, mAngleN, mAngleP, n, i1, x, i2, n)
	filename := fmt.Sprintf("moon28f%02d.svg", n)
	createFile(filename, text)
}
func createFile(fname, ftext string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes := []byte(ftext)
	if _, err := f.Write(bytes); err != nil {
		return err
	}
	return nil
}
