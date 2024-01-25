package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/webSvgCirclePathText"
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
func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `Web serving embeded directory tree in root /
 Usage:
 -h this help
 -t top_text    #underscores are transformed to space
 -b bottom_text
 -p [port] preview to port 8080 
 Example:
 -b StartUp1 -t Serving_to_our_customers_all_day -p
 -t Serving_to_our_customers -b Company -p 8081
`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	topText := "Top longer descriptive text"
	bottomText := "Bottom text"

	colonPort := ":8080"
	if len(os.Args) < 2 {
		help("Not enough arguments")
	} else {
		for i := 1; i < len(os.Args); i++ {
			switch arg := os.Args[i]; arg {
			case "-h":
				help("")
				os.Exit(0)
			case "-b":
				i += 1
				bottomText = underscoredText(i)
			case "-t":
				i += 1
				topText = underscoredText(i)
			case "-p":
				a.SetVariables(topText, bottomText)
				if len(os.Args) > i+1 {
					colonPort = ":" + os.Args[i+1]
				}
				http.HandleFunc("/", a.SvgHandler)
				print("Serving SVG page Listenning at " + colonPort + ". CTRL+C to stop.")
				http.ListenAndServe(colonPort, nil)
			default:
				help("Unexpected argument " + arg)
			}
		}
	}
}
