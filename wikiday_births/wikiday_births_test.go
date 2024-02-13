package wikiday_births

import (
	"testing"
)

func TestMakeUrlString1(t *testing.T) {
	want := "https://en.wikipedia.org/wiki/January_1"
	got := makeUrlString(1, 1)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
func TestMakeUrlString2(t *testing.T) {
	want := "https://en.wikipedia.org/wiki/February_14"
	got := makeUrlString(2, 14)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
