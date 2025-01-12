package main

import (
	"fmt"
	"os"
	"time"

	a "github.com/chlachula/go2/starTime"
)

func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `Web displaying directory
 Usage:
 -t
 -w Wikipedia definition
 -h #this help
 Example:
 go2 -t 2025-01-11_23:15:30 2025-02-30
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
			case "-w":
				fmt.Print(a.WikiDefinition)
			case "-t":
				if i+2 >= len(os.Args) {
					help(fmt.Sprintf("Not enough arguments i=%d, len=%d", i, len(os.Args)))
					os.Exit(1)
				}
				a.Compute(os.Args[i+1], os.Args[i+2])
				i += 2
			default:
				help("Unexpected argument " + arg)
			}
		}
	}

}
