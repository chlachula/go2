package moonEphem

import (
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
func TestMoonJ2000XYZ(t *testing.T) {
	want := "a"
	got := "b"
	wantXYZ := MoonJ2000XYZ_legacy(0.0)
	gotXYZ := MoonJ2000XYZ(0.0)
	if got != want {
		t.Errorf("Unexpected error\nWANT legacy: %v\n        GOT: %v", wantXYZ, gotXYZ)
	}
}
