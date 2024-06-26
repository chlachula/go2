package main

import (
	"fmt"
	"os"
	"time"

	a "github.com/chlachula/go2/pi"
)

func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `PI computing
 Usage:
 -b  Bailey-Borwein-Plouffe formula of the calculating PI https://en.wikipedia.org/wiki/Bailey%E2%80%93Borwein%E2%80%93Plouffe_formula
 -l  Leibniz formula of the calculating PI
 -h  #this help
 Examples:
 go2 -b #Bailey-Borwein-Plouffe method
 go2 -l #Leibniz method 
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
			case "-b":
				// other algorithms like the Bailey-Borwein-Plouffe formula are more efficient
				a.BaileyBorweinPlouffe(11)
			case "-l":
				a.Leibniz()
			default:
				help("Unexpected argument " + arg)
			}
		}
	}

}
