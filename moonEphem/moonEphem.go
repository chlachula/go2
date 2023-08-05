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
	"math"
	"time"
)

var degreesToRadians = math.Pi / 180.0
var radiansToDegrees = 180.0 / math.Pi

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
	x360 := math.Remainder(x, 360.0)
	if x360 < 0.0 {
		x360 += 360.0
	}
	return x360
}

// J2000.0 Moon cartese coordinates in metres, t is in Julian centuries
func MoonJ2000XYZ(t float64) (x [3]float64) {
	for n := 0; n < 3; n++ {
		x[n] = 0.0
		for m := 0; m < 2; m++ {
			x[n] = a[n][m] * math.Sin(w[n][m]*t+d[n][m])
		}
		x[n] *= 1.0e6
	}
	return x
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
	L := q + 1.915*math.Sin(gRad) + 0.020*math.Sin(2.0*gRad)
	return to0_360(L)
}

// distance to Sun in AU
func DistanceToSun(g float64) float64 {
	gRad := g * degreesToRadians
	R := 1.00014 - 0.01671*math.Cos(gRad) - 0.00014*math.Cos(2.0*gRad)
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
	numerator := math.Cos(e*degreesToRadians) * math.Sin(Lrad)
	denominator := math.Cos(Lrad)
	ra := math.Atan2(numerator, denominator)
	return ra * radiansToDegrees
}

// declination in deg
func DeclinationOfTheSun(d float64) float64 {
	//sin d = sin e sin L
	L := GeocentricApparentEclipticLongitudeOfTheSunAdjustedForAberration(d)
	Lrad := L * degreesToRadians
	e := EclipticLongitudeofTheSun(d)

	sinD := math.Sin(e*degreesToRadians) * math.Sin(Lrad)
	return math.Asin(sinD) * radiansToDegrees
}

// rectascention and declination of the Sun
func RA_Dec_OfTheSun(date time.Time) (ra float64, decl float64) {
	d := J2000Days(date)
	L := GeocentricApparentEclipticLongitudeOfTheSunAdjustedForAberration(d)
	Lrad := L * degreesToRadians
	e := EclipticLongitudeofTheSun(d)
	eRad := e * degreesToRadians

	numerator := math.Cos(eRad) * math.Sin(Lrad)
	denominator := math.Cos(Lrad)
	ra = math.Atan2(numerator, denominator) * radiansToDegrees
	ra = to0_360(ra)

	sinD := math.Sin(e*degreesToRadians) * math.Sin(Lrad)
	decl = math.Asin(sinD) * radiansToDegrees

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
	ra = math.Atan2(numerator, denominator) * radiansToDegrees
	ra = to0_360(ra)

	tanD := z / math.Sqrt(x*x+y*y)
	decl = math.Atan(tanD) * radiansToDegrees
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
	deg = math.Abs(deg)
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
