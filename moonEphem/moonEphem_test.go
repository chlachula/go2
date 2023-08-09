package moonEphem

import (
	"math"
	"testing"
)

var horizonMoonEphemerids string = `******************************************************************************************************************************************************************************
Date__(UT)__HR:MN     R.A._____(ICRF)_____DEC    APmag   S-brt             delta      deldot     S-O-T /r     S-T-O  Sky_motion  Sky_mot_PA  RelVel-ANG  Lun_Sky_Brt  sky_SNR
******************************************************************************************************************************************************************************
$$SOE
2024-Jan-01 00:00     10 35 11.55 +12 45 08.3  -11.150   4.560  0.00270481348636   0.0084978  123.9718 /L   55.8978   29.622877   115.41135   0.5059510         n.a.     n.a.
2024-Jan-04 00:00     12 42 21.31 -03 39 02.0  -10.126   5.151  0.00268330633882  -0.0340945   91.5991 /L   88.2446   29.986838   117.97343   -2.007440         n.a.     n.a.
2024-Jan-07 00:00     15 00 07.62 -19 29 07.5   -8.671   5.821  0.00259315380892  -0.0653085   57.6737 /L  122.1983   32.205777   110.69968   -3.704584         n.a.     n.a.
2024-Jan-10 00:00     17 54 21.23 -28 05 21.5   -6.120   6.318  0.00248104062403  -0.0559723   20.6437 /L  159.3050   35.374475   92.242331   -3.025521         n.a.     n.a.
2024-Jan-13 00:00     21 06 02.80 -21 25 21.2   -6.222   6.325  0.00242269210216  -0.0078065   21.1966 /T  158.7525   37.102198   71.115932   -0.417062         n.a.     n.a.
2024-Jan-16 00:00     23 51 17.02 -03 02 40.6   -9.015   5.736  0.00245101565831   0.0359554   61.8113 /T  118.0628   36.073509   61.912754   1.9240867         n.a.     n.a.
2024-Jan-19 00:00     02 22 42.35 +15 57 32.0  -10.571   4.988  0.00252783833932   0.0477457  100.5636 /T   79.2917   33.863369   66.469104   2.6430095         n.a.     n.a.
2024-Jan-22 00:00     05 06 20.30 +27 11 17.0  -11.586   4.317  0.00260580973598   0.0409462  136.7856 /T   43.1106   32.019498   82.258980   2.3290755         n.a.     n.a.
2024-Jan-25 00:00     07 55 03.37 +25 52 04.7  -12.415   3.586  0.00266806097357   0.0304611  170.2627 /T    9.7109   30.682219   101.84989   1.7689081         n.a.     n.a.
2024-Jan-28 00:00     10 21 28.92 +14 06 45.6  -11.983   3.942  0.00270719828836   0.0126337  154.9037 /L   25.0298   29.747173   114.81101   0.7487826         n.a.     n.a.
2024-Jan-31 00:00     12 28 55.43 -02 09 30.1  -11.112   4.586  0.00270312465659  -0.0196720  122.7478 /L   57.1203   29.649221   118.22393   -1.160192         n.a.     n.a.
$$EOE`

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
