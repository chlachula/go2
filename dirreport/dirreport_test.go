package dirreport

import "testing"

func TestNum10p3str(t *testing.T) {
	var n int64
	want := "  0 "
	got := num10p3str(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 1000
	want = "1.0K"
	got = num10p3str(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 9900
	want = "9.9K"
	got = num10p3str(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 9949 * 1000 * 1000 * 1000
	want = "9.9T"
	got = num10p3str(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
	n = 9900 * 1000 * 1000 * 1000 * 1000
	want = "9.9P" // Peta
	got = num10p3str(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}

	n = int64(^uint64(0) >> 1) // max int64
	want = "9.2E"              // Exa
	got = num10p3str(n)
	if want != got {
		t.Errorf("want %s, got %s for number %d", want, got, n)
	}
}
