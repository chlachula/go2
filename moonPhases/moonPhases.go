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

const PhasesNumber = 28

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
	text := fmt.Sprintf(svgBase, n, i1, x, i2, n)
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
