package moonView

import "testing"

/*
   "2011" => "a003800/a003810",
   "2012" => "a003800/a003894",
   "2013" => "a004000/a004000",
   "2014" => "a004100/a004118",
   "2015" => "a004200/a004236",
   "2016" => "a004400/a004404",
   "2017" => "a004500/a004537",
   "2018" => "a004600/a004604",
   "2019" => "a004400/a004442",
   "2020" => "a004700/a004768",
   "2021" => "a004800/a004874",
   "2022" => "a004900/a004955",
   "2023" => "a005000/a005048"
   "2024" => "a005100/a005187"
   
   https://svs.gsfc.nasa.gov/vis/a000000/a005100/a005187/frames/730x730_1x1_30p/moon.0001.jpg
*/
func TestSvsMagicNumbers(t *testing.T) {
	year := 2011
	wantNN00 := 3800
	wantNNNN := 3810
	gotNN00, gotNNNN := svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}

	year = 2017
	wantNN00 = 4500
	wantNNNN = 4537
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}

	year = 2024 //last available year
	wantNN00 = 5100
	wantNNNN = 5187
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}
	year = 1955
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}
	year = 2055
	gotNN00, gotNNNN = svsMagicNumbers(year)
	if gotNN00 != wantNN00 || gotNNNN != wantNNNN {
		t.Errorf("For year %d WANT: %d/%d; GOT: %d/%d", year, wantNN00, wantNNNN, gotNN00, gotNNNN)
	}
}
