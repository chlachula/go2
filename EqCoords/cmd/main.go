package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/EqCoords"
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
	helptext := `Equatorial coordinates conversion from J2000.0 
 Usage:
 -t date       
 -c coordinates
 Example:
 go2 -t 2024-04-21 -c 18h36m56.33635s,+38°47'01.2802" #Vega J2000.0
`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	if len(os.Args) < 2 {
		help("Not enough arguments")
	} else {
		for i := 1; i < len(os.Args); i++ {
			switch arg := os.Args[i]; arg {
			case "-h":
				help("")
				os.Exit(0)
			case "-t":
				i += 1
				a.SetOutputTime(os.Args[i])
			case "-c":
				i += 1
				a.ConvertCoordsStr(os.Args[i])
			default:
				help("Unexpected argument " + arg)
			}
		}
	}

}
