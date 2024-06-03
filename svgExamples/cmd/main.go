package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	a "github.com/chlachula/go2/svgExamples"
)

func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `Web: SVG on HTML example pages
 Usage:
 -h this help
 -p [port] preview to port 8080 
 Example:
 -p
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
			case "-p":
				http.HandleFunc("/", a.HandlerHome)
				http.HandleFunc("/svgExamples", a.HandlerSvgExamples)
				http.HandleFunc("/svgExamples/textPath1", a.HandlerHtmlTextPath1)
				http.HandleFunc("/img/svg/textPath1", a.HandlerImgSvgTextPath1)
				what := "svgExamples"
				print("Serving " + what + " at " + colonPort + ". CTRL+C to stop.")
				http.ListenAndServe(colonPort, nil)
			default:
				help("Unexpected argument " + arg)
			}
		}
	}
}
