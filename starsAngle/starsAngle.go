package starsAngle


// Angular distance of two stars
package main

import (
	"fmt"
	. "math"
)

type polarCoordinates struct {
	RA float64
	De float64
}

var (
	degreesToRadians = Pi / 180.0
	sToRadians       = Pi / 180.0 / 240.0 // 360/86400=1/240
	radiansToDegrees = 180.0 / Pi
)

func dmsToRad(d, m, s float64) float64 {
	deg := d + (m+s/60.0)/60.0
	return deg * degreesToRadians
}

func dms2rad(d, m int, s float64) float64 {
	return dmsToRad(float64(d), float64(m), s)
}

func hmsToRad(h, m, s float64) float64 {
	seconds := (h*60.0+m)*60.0 + s
	return seconds * sToRadians
}
func hms2rad(h, m int, s float64) float64 {
	return hmsToRad(float64(h), float64(m), s)
}

func AngularDistance(A, B polarCoordinates) float64 {
	cosT := Cos(A.De)*Cos(B.De)*Cos(A.RA-B.RA) + Sin(A.De)*Sin(B.De)
	return Acos(cosT)
}

func main() {
	a := polarCoordinates{ //HD 92956 8.13mag 10h 45m07.2s   -06°12'17.8"
		RA: hms2rad(10, 45, 7.2),
		De: -dms2rad(6, 12, 17.8),
	}
	b := polarCoordinates{ //HD 90362 5.69mag 10h 26m57.1s   -07°10'53.4"
		RA: hms2rad(10, 26, 57.1),
		De: -dms2rad(7, 10, 53.4),
	}
	fmt.Println("AngularDistance(a,b) in degrees =", AngularDistance(a, b)*radiansToDegrees)
}