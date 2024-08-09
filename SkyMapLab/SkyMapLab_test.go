package SkyMapLab

import (
	"fmt"
	"math"
	"testing"
)

func TestLoadECSV(t *testing.T) {
	filename := "data/test01.csv"
	if rows, err := LoadECSV(filename); err != nil {
		t.Errorf("error loading file " + filename + ": " + err.Error())
	} else {
		want := 4
		got := len(rows)
		if want != got {
			t.Errorf("want %d rows, got %d for file %s", want, got, filename)
		}
		if rows[3][1] != "April" {
			t.Errorf("4th row, 2nd columns should have value ' April', got %s", rows[3][1])
		}
	}

}
func TestEclipticalToEquatorial(t *testing.T) {
	𝜀Deg2025 = 23.436040
	toRad := math.Pi / 180.0
	for La := 0.0; La < 360.1; La = La + 30.0 {
		LaR := La * toRad
		ra, de := EclipticalToEquatorial(LaR, 0.0, 𝜀Deg2025)
		raD := ra / toRad
		deD := de / toRad
		fmt.Printf("La:%3.0f    ra:%6.2f,de:%6.2f\n", La, raD, deD)
	}
}
func TestAzimutalToEquatoreal_I(t *testing.T) {
	toRad := math.Pi / 180.0
	toDeg := 180.0 / math.Pi
	fi := 50.0
	fiR := fi * toRad
	fmt.Printf("fi:%3.0f\n", fi)
	for Az := 0.0; Az < 360.1; Az = Az + 30.0 {
		AzR := Az * toRad
		t, de := AzimutalToEquatoreal_I(AzR, 0.0, fiR)
		tD := t * toDeg
		deD := de * toDeg
		fmt.Printf("Az:%3.0f    t:%6.2f,de:%6.2f\n", Az, tD, deD)
	}
}

func TestCartesianXY(t *testing.T) {
	gotX, gotY := cartesianXY(1.0, 0.0)
	if gotX != 0.0 {
		t.Errorf("unexpected X %.4f instead of 0.0", gotX)
	}
	if gotY != 1.0 {
		t.Errorf("unexpected Y %.4f instead of 1.0", gotY)
	}
}
