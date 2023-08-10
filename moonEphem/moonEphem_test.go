package moonEphem

import (
	"math"
	"testing"
)

func TestRAstring_180(t *testing.T) {
	want := "12 00  0.00"
	got := RAstring(180.0)
	if got != want {
		t.Errorf("Unexpected error - WANT: %s; GOT: %s", want, got)
	}
}
func TestRAstring_75(t *testing.T) {
	want := "00 30  0.00"
	got := RAstring(7.5)
	if got != want {
		t.Errorf("Unexpected error - WANT: %s; GOT: %s", want, got)
	}
}
func TestRAstring_1s(t *testing.T) {
	want := "00 00  1.00"
	got := RAstring(1.0 / 3600 * 15.0) //1 / 240
	if got != want {
		t.Errorf("Unexpected error - WANT: %s; GOT: %s", want, got)
	}
}
func TestDeclStringMinus45(t *testing.T) {
	want := "-45 00  0.00"
	got := DeclString(-45.0)
	if got != want {
		t.Errorf("Unexpected error - WANT: %s; GOT: %s", want, got)
	}
}
func TestDeclStringPlus30Min(t *testing.T) {
	want := "+00 30  0.00"
	got := DeclString(0.5)
	if got != want {
		t.Errorf("Unexpected error - WANT: %s; GOT: %s", want, got)
	}
}
func TestDeclStringPlus1s(t *testing.T) {
	want := "+27 30  1.00"
	got := DeclString(27.5 + (1.0 / 3600.0))
	if got != want {
		t.Errorf("Unexpected error - WANT: %s; GOT: %s", want, got)
	}
}

// MoonJ2000XYZ_legacy(t float64) (xyz [3]float64)
func NoTestMoonJ2000XYZ(t *testing.T) {
	want := "a"
	got := "b"
	wantXYZ := MoonJ2000XYZ_legacy(0.0)
	gotXYZ := MoonJ2000XYZ(0.0)
	if got != want {
		rw := SpaceDiagonal(wantXYZ)
		rg := SpaceDiagonal(gotXYZ)
		rate := (rw - rg) / rg * 100.0
		t.Errorf("Unexpected error\nWANT legacy: %v\nspaceDiagoval %.4e\n        GOT: %v\nspaceDiagoval %.4e\nrate = %.2f \n", wantXYZ, rw, gotXYZ, rg, rate)
	}
}
func TestAngularDistance(t *testing.T) {
	want := math.Pi * 0.5
	got := AngularDistance(0.0, 0.0, math.Pi*0.5, 0.0)
	if got != want {
		t.Errorf("Unexpected error1 - WANT: %f; GOT: %f", want, got)
	}
	got = AngularDistance(0.0, 0.0, 0.0, math.Pi*0.5)
	if got != want {
		t.Errorf("Unexpected error2 - WANT: %f; GOT: %f", want, got)
	}
	got = AngularDistance(0.0, 0.0, 0.0, -math.Pi*0.5)
	if got != want {
		t.Errorf("Unexpected error3 - WANT: %f; GOT: %f", want, got)
	}

	want = math.Pi
	got = AngularDistance(0.0, math.Pi*0.5, math.Pi, -math.Pi*0.5)
	if got != want {
		t.Errorf("Unexpected error4 - WANT: %f; GOT: %f", want, got)
	}
	got = AngularDistance(1.0, math.Pi*0.25, 1.0+math.Pi, -math.Pi*0.25)
	if got != want {
		t.Errorf("Unexpected error5 - WANT: %f; GOT: %f", want, got)
	}

	// http://www.astronomycafe.net/FAQs/q1890x.html#:~:text=You%20can%20check%20this%20by,calculator%20over%20at%20Celestial%20WOnders.
	//1 Sirius     is at 6h 41m and -16d 35' so ra1 = 100.2 degrees and d1 = -16.58 degrees.
	//2 Betelgeuse is at 5h 50m and  +7d 23' so ra2 =  87.5 degrees and d2 =   7.38 degrees
	ra1 := 100.2 * degreesToRadians
	d1 := -16.58 * degreesToRadians
	ra2 := 87.5 * degreesToRadians
	d2 := 7.38 * degreesToRadians
	want = 27.1 //degrees
	got = math.Round(AngularDistance(ra1, d1, ra2, d2)*radiansToDegrees*10.0) / 10.0
	if got != want {
		t.Errorf("Unexpected error6 - WANT: %f; GOT: %f", want, got)
	}
}
