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
