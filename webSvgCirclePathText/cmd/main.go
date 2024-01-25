package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	a "github.com/chlachula/go2/webSvgCirclePathText"
)

func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `Web serving embeded directory tree in root /
 Usage:
  -h this help
  -p [port] preview to port 8080 
`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	a.SetVariables("Upper longer descriptive text", "Bottom text")

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
