package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/wikiday_births"
)

func quotesRemove(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Displays births and deads of the day [en]
	Usage:
	go2 -h #this help
	go2 -d day #in format mm/dd
	Example:
	go2 -d 12/31 # births and deads for December 31st
	go2 -d 2/28  # births and deads for February 28th
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	if len(os.Args) < 2 {
		help("Not enough arguments")
		os.Exit(1)
	}
	if strings.HasPrefix(os.Args[1], "-h") {
		help("")
		os.Exit(0)
	}
	if len(os.Args) < 3 {
		help("Not enough arguments")
		os.Exit(1)
	}

	a.WikiDay(os.Args[2])

}
