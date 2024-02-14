package wikiday_births

import (
	"testing"
)

func TestMakeUrlString1(t *testing.T) {
	want := "January_1"
	got := makeWikiDayFilename(1, 1)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
func TestMakeUrlString2(t *testing.T) {
	want := "February_14"
	got := makeWikiDayFilename(2, 14)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
