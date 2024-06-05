package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	a "github.com/chlachula/go2/SkyMapLab"
)

func underscoredText(i int) string {
	if i < len(os.Args) {
		return strings.ReplaceAll(os.Args[i], "_", " ")
	} else {
		help("missing bottom text argument")
		os.Exit(1)
		return ""
	}
}
func toFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `SkyMap Lab to create mostly visual objects sky maps in SVG
 Usage:
 -h this help
 -l latitude
 -r radius
 -p [port] preview to port 8080 
 Example:

`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	radius := 150.0
	lat := 44.0

	colonPort := ":8080"
	if len(os.Args) < 2 {
		help("Not enough arguments")
	} else {
		for i := 1; i < len(os.Args); i++ {
			switch arg := os.Args[i]; arg {
			case "-h":
				help("")
				os.Exit(0)
			case "-l":
				if len(os.Args) > i+1 {
					lat = toFloat64(os.Args[i+1])
					i += 1
				}
			case "-r":
				if len(os.Args) > i+1 {
					radius = toFloat64(os.Args[i+1])
					i += 1
				}
			case "-p":
				if len(os.Args) > i+1 {
					colonPort = ":" + os.Args[i+1]
				}
				a.LoadStars("data/SkyMap-stars.json")
				a.LoadConstellations("data/SkyMap-constellations.json")
				a.SetMapStyle(radius, lat, a.MapColorsRed)
				http.HandleFunc("/", a.HandlerHome)
				http.HandleFunc("/img/svg/skymap/{colorId}/{latId}", a.HandlerSkyMapGeneral)
				http.HandleFunc("/SkyMapLab", a.HandlerSkyMapLab)
				print("Serving SVG page Listenning at " + colonPort + ". CTRL+C to stop.")
				http.ListenAndServe(colonPort, nil)
			default:
				help("Unexpected argument " + arg)
			}
		}
	}
}
