package dirreport

import "testing"

func TestStrSize(t *testing.T) {
	var n int64
	want := "  0 "
	got := strSize(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 1000
	want = "1.0K"
	got = strSize(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 9900
	want = "9.9K"
	got = strSize(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 9949 * 1000 * 1000 * 1000
	want = "9.9T"
	got = strSize(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 9900 * 1000 * 1000 * 1000 * 1000
	want = "9.9P" // Peta
	got = strSize(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
}
