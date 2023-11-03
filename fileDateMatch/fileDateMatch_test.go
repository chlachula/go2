package fileDateMatch

import (
	"testing"
	"time"
)

func TestDateInsideFile(t *testing.T) {
	want := time.Date(2023, time.November, 5, 0, 0, 0, 0, time.Local)
	var lines = []string{"first line", `<date yyyy="2023" mm="11" dd="5" />`, "last line"}
	got, err := dateInsideFile(lines)

	if err != nil {
		t.Errorf("unexpected error -  %s", err.Error())
	}
	if want.Format(time.DateOnly) != got.Format(time.DateOnly) {
		t.Errorf("Unexpected date %s instead of  %s", got.Format(time.DateOnly), want.Format(time.DateOnly))
	}
}
