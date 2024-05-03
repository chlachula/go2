package SkyMapLab

import "testing"

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
