package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/webSvgBrokenCurveOfLines"
)

func underscoredText(i int) string {
	if i < len(os.Args) {
		return strings.ReplaceAll(os.Args[i], "_", " ")
	} else {
		help("missing argument")
		os.Exit(1)
		return ""
	}
}
func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `Web: Broken Curve Of streight lines
 Usage:
 -h this help
 -d data points
 -p [port] preview to port 8080 
 Example:
 -d 1,1_2,2_1,1_1,1.5_1.2,1.6 -p
`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	colonPort := ":8080"
	if len(os.Args) < 2 {
		help("Not enough arguments")
	} else {
		for i := 1; i < len(os.Args); i++ {
			switch arg := os.Args[i]; arg {
			case "-h":
				help("")
				os.Exit(0)
			case "-d":
				i += 1
				a.DataLine = underscoredText(i)
			case "-p":
				http.HandleFunc("/", a.HandlerRoot)
				http.HandleFunc(a.URL, a.HandlerSvgBrokenCurveOfLines)
				what := "SVG Broken Curve Of straight lines"
				print("Serving " + what + "at " + colonPort + ". CTRL+C to stop.")
				http.ListenAndServe(colonPort, nil)
			default:
				help("Unexpected argument " + arg)
			}
		}
	}
}
