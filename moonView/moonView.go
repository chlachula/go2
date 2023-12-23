package moonView

import (
	"fmt"
	"net/http"
)

func svsMagicNumbers(y int) (int, int) {
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

	*/
	svsMagic := []int{
		3810, //2011
		3894, //2012
		4000, //2013
		4118, //2014
		4236, //2015
		4404, //2016
		4537, //2017
		4604, //2018
		4442, //2019
		4768, //2020
		4874, //2021
		4955, //2022
		5048, //2023
		5187, //2024
	}
	y1 := 2011
	i := len(svsMagic) - 1
	y2 := y1 + i

	if y >= y1 && y <= y2 {
		i = y - y1
	}
	nnnn := svsMagic[i]
	nn := nnnn / 100
	nn00 := nn * 100
	return nn00, nnnn
}
func EventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, part1)
	fmt.Fprint(w, part2)
	imgURL := "https://svs.gsfc.nasa.gov/vis/a000000/a005000/a005048/frames/730x730_1x1_30p/moon.8456.jpg"
	println(imgURL)
	fmt.Fprintf(w, part3, imgURL)
	fmt.Fprint(w, part4)
}
