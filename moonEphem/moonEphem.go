/*
	David G. Simpson: AN ALTERNATIVE LUNAR EPHEMERIS MODEL FOR ON-BOARD FLIGHT SOFTWARE USE
	https://caps.gsfc.nasa.gov/simpson/pubs/slunar.pdf
	Error about 0.341 deg, and a maximum error of 1.03 deg 2000 - 2100

...
where all angles are given in radians for convenience of use in software, t is the time
in Julian centuries from J2000 given by Eq. (10), and X, Y , and Z are the Cartesian
components of the lunar position vector, referred to the mean equator and equinox of
J2000.

for (n=0; n<3; n++)
{
x[n] = 0.0;
for (m=0; m<7; m++)
x[n] += a[n][m]*sin(w[n][m]*t+delta[n][m]);
}

https://aa.usno.navy.mil/faq/sun_approx  (1' precission)
D = JD – 2451545.0
where JD is the Julian date of interest. Then compute

Mean anomaly of the Sun:	g = 357.529 + 0.98560028 D
Mean longitude of the Sun:	q = 280.459 + 0.98564736 D
Geocentric apparent ecliptic longitude of the Sun (adjusted for aberration):
L = q + 1.915 sin g + 0.020 sin 2g
where all the constants (therefore g, q, and L) are in degrees. It may be necessary or desirable to reduce g, q, and L to the range 0° to 360°.

The Sun's ecliptic latitude, b, can be approximated by b=0.
The distance of the Sun from the Earth, R, in astronomical units (AU), can be approximated by
R = 1.00014 – 0.01671 cos g – 0.00014 cos 2g
Once the Sun's apparent ecliptic longitude, L, has been computed, the Sun's right ascension and declination can be obtained. First compute the mean obliquity of the ecliptic, in degrees:

e = 23.439 – 0.00000036 D
Then the Sun's right ascension, RA, and declination, d, can be obtained from

tan RA = cos e sin L / cos L
sin d = sin e sin L
*/
package moonEphem

import (
	"fmt"
	. "math"
	"time"
)

var degreesToRadians = Pi / 180.0
var radiansToDegrees = 180.0 / Pi

var a = [][]float64{
	{383.0, 31.5, 10.6, 6.2, 3.2, 2.3, 0.8},
	{351.0, 28.9, 13.7, 9.7, 5.7, 2.9, 2.1},
	{153.2, 31.5, 12.5, 4.2, 2.5, 3.0, 1.8},
}
var w = [][]float64{
	{8399.685, 70.990, 16728.377, 1185.622, 7143.070, 15613.745, 8467.263},
	{8399.687, 70.997, 8433.466, 16728.380, 1185.667, 7143.058, 15613.755},
	{8399.672, 8433.464, 70.996, 16728.364, 1185.645, 104.881, 8399.116},
}
var d = [][]float64{
	{5.381, 6.169, 1.453, 0.481, 5.017, 0.857, 1.010},
	{3.811, 4.596, 4.766, 6.165, 5.164, 0.300, 5.565},
	{3.807, 1.629, 4.595, 6.162, 5.167, 2.555, 6.248},
}

func to0_360(x float64) float64 {
	x360 := Remainder(x, 360.0)
	if x360 < 0.0 {
		x360 += 360.0
	}
	return x360
}
func sinD(a float64) float64 {
	return Sin(a * degreesToRadians)
}
func cosD(a float64) float64 {
	return Cos(a * degreesToRadians)
}
func tanD(a float64) float64 {
	return Tan(a * degreesToRadians)
}

// J2000.0 Moon cartese coordinates in metres, t is in Julian centuries
func MoonJ2000XYZ_legacy(t float64) (xyz [3]float64) {
	lambda := 218.32 + 481267.883*t +
		6.29*sinD((477198.85*t+134.9)) -
		1.27*sinD((-413335.38*t+259.2)) +
		0.66*sinD((890534.23*t+235.7)) +
		0.21*sinD((954397.70*t+269.9)) +
		0.19*sinD((35999.05*t+357.5)) +
		0.11*sinD((966404.05*t+186.6))

	beta := 5.13*sinD((483202.03*t+93.3)) +
		0.28*sinD((960400.87*t+228.2)) -
		0.28*sinD((6003.18*t+318.3)) -
		0.17*sinD((-407332.20*t+217.6))
	piM := 0.9508 +
		0.0518*cosD(477198.85*t+134.9) +
		0.0095*cosD(-413335.38*t+259.2) +
		0.0078*cosD(890534.23*t+235.7) +
		0.0028*cosD(954397.70*t+269.9)
	Rearth := 6378140.0 // m
	r := Rearth / sinD(piM)
	eps0 := 23.43929111111111111 // 23° 26' 21.448"
	//precession constants a,b,c
	a := t * (1.396971 + 0.0003086*t)
	b := t * (0.013056 - 0.0000092*t)
	c := 5.12362 - t*(1.155358+0.0001964*t)
	beta0 := beta - b*sinD((lambda+c))
	lambda0 := lambda - a + b*cosD((lambda+c))*tanD(beta0)
	xyz[0] = r * cosD(beta0) * cosD(lambda0)
	xyz[1] = r * (cosD(beta0)*sinD(lambda0)*cosD(eps0) - sinD(beta0)*sinD(eps0))
	xyz[2] = r * (cosD(beta0)*sinD(lambda0)*sinD(eps0) + sinD(beta0)*cosD(eps0))
	return xyz
}

// J2000.0 Moon cartese coordinates in metres, t is in Julian centuries
func MoonJ2000XYZ(t float64) (xyz [3]float64) {
	for n := 0; n < 3; n++ {
		xyz[n] = 0.0
		for m := 0; m < 2; m++ {
			xyz[n] = a[n][m] * Sin(w[n][m]*t+d[n][m])
		}
		xyz[n] *= 1.0e6
	}
	return xyz
}

// julian days  since J2000
func J2000Days(date time.Time) float64 {
	d2000 := time.Date(2000, time.January, 1, 12, 0, 0, 0, time.UTC)
	delta := date.Sub(d2000)
	days := delta.Seconds() / 86400.0
	return days
}

// julian centuries since J2000
func J2000Centuries(date time.Time) float64 {
	julianCenturyDays := 36525.0
	return J2000Days(date) / julianCenturyDays
}

// mean anomaly in deg
func MeanAnomalyOfTheSun(d float64) float64 {
	g := 357.529 + 0.98560028*d
	return to0_360(g)
}

// mean longitide in deg
func MeanLongitudeOfTheSun(d float64) float64 {
	q := 280.459 + 0.98564736*d
	return to0_360(q)
}

// ecliptic longiture in deg
func GeocentricApparentEclipticLongitudeOfTheSunAdjustedForAberration(d float64) float64 {
	gRad := MeanAnomalyOfTheSun(d) * degreesToRadians
	q := MeanLongitudeOfTheSun(d)
	L := q + 1.915*Sin(gRad) + 0.020*Sin(2.0*gRad)
	return to0_360(L)
}

// distance to Sun in AU
func DistanceToSun(g float64) float64 {
	gRad := g * degreesToRadians
	R := 1.00014 - 0.01671*Cos(gRad) - 0.00014*Cos(2.0*gRad)
	return R
}

// Sun's ecliptic latitude, b, can be approximated by b=0.
func EclipticLongitudeofTheSun(d float64) float64 {
	e := 23.439 - 0.00000036*d
	return to0_360(e)
}

// RA in deg
func RightAccessionOfTheSun(d float64) float64 {

	//tan RA = cos e sin L / cos L
	L := GeocentricApparentEclipticLongitudeOfTheSunAdjustedForAberration(d)
	Lrad := L * degreesToRadians
	e := EclipticLongitudeofTheSun(d)
	numerator := Cos(e*degreesToRadians) * Sin(Lrad)
	denominator := Cos(Lrad)
	ra := Atan2(numerator, denominator)
	return ra * radiansToDegrees
}

// declination in deg
func DeclinationOfTheSun(d float64) float64 {
	//sin d = sin e sin L
	L := GeocentricApparentEclipticLongitudeOfTheSunAdjustedForAberration(d)
	Lrad := L * degreesToRadians
	e := EclipticLongitudeofTheSun(d)

	sinD := Sin(e*degreesToRadians) * Sin(Lrad)
	return Asin(sinD) * radiansToDegrees
}

// rectascention and declination of the Sun
func RA_Dec_OfTheSun(date time.Time) (ra float64, decl float64) {
	d := J2000Days(date)
	L := GeocentricApparentEclipticLongitudeOfTheSunAdjustedForAberration(d)
	Lrad := L * degreesToRadians
	e := EclipticLongitudeofTheSun(d)
	eRad := e * degreesToRadians

	numerator := Cos(eRad) * Sin(Lrad)
	denominator := Cos(Lrad)
	ra = Atan2(numerator, denominator) * radiansToDegrees
	ra = to0_360(ra)

	sinD := Sin(e*degreesToRadians) * Sin(Lrad)
	decl = Asin(sinD) * radiansToDegrees

	//fmt.Printf("DEBUG days=%.2f L=%.2f e=%.2f ra=%.2f decl=%.2f       ", d, L, e, ra, decl)
	return
}

// rectascention and declination of the Sun
func RA_Dec_OfTheMoon(date time.Time) (ra float64, decl float64) {
	t := J2000Centuries(date)
	fmt.Println("DEBUG time", date, t)
	xyz := MoonJ2000XYZ(t)
	x := xyz[0]
	y := xyz[1]
	z := xyz[2]
	numerator := y
	denominator := x
	ra = Atan2(numerator, denominator) * radiansToDegrees
	ra = to0_360(ra)

	tanD := z / Sqrt(x*x+y*y)
	decl = Atan(tanD) * radiansToDegrees
	return
}
func RAstring(deg float64) string {
	hh := int(deg / 15.0) // 355deg = 23h 40m
	mm := int(4.0 * (deg - float64(hh)*15.0))
	ss := 60.0 * (4.0*deg - float64(mm) - float64(hh)*60.0)
	return fmt.Sprintf("%02d %02d %5.2f", hh, mm, ss)
}
func DeclString(deg float64) string {
	sign := "+"
	if deg < 0.0 {
		sign = "-"
	}
	deg = Abs(deg)
	dd := int(deg)
	mm := int(60.0 * (deg - float64(dd)))
	ss := 3600.0*deg - 60.0*(float64(mm)+60.0*float64(dd))
	return fmt.Sprintf("%s%02d %02d %5.2f", sign, dd, mm, ss)
}
func SunEphemerides(date time.Time, stepDays float64, stepsNumber int) {
	fmt.Printf("%s\n%-19s %-11s %-11s\n", "Sun", "Date (UTC)", "Right asc.", "Declination")
	for i := 0; i < stepsNumber; i++ {
		//		iDate := date.Add(24.0*untyped float(time.Hour)*time.Duration(stepDays))
		// - 28104*time.Second
		iDate := date.Add(time.Second * time.Duration(86400.0*stepDays*float64(i)))
		//19 04 28.27 -22 36 33.6
		ra, decl := RA_Dec_OfTheSun(iDate)
		fmt.Println(iDate.Format("2006-01-02 15:04:05"), RAstring(ra), DeclString(decl))
	}
}
func MoonEphemerides(date time.Time, stepDays float64, stepsNumber int) {
	fmt.Printf("%s\n%-19s %-11s %-11s\n", "Moon", "Date (UTC)", "Right asc.", "Declination")
	for i := 0; i < stepsNumber; i++ {
		iDate := date.Add(time.Second * time.Duration(86400.0*stepDays*float64(i)))
		ra, decl := RA_Dec_OfTheMoon(iDate)
		fmt.Println(iDate.Format("2006-01-02 15:04:05"), RAstring(ra), DeclString(decl))
	}
}
