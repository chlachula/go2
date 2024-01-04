package main

import (
	"fmt"
	"os"
	"time"

	a "github.com/chlachula/go2/aptitles"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `APOD titles
	-h this help
	-c create APOD archive with titles to local json file
	-l load local json APOD archive 
	-t yymmdd seach title for give date
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

	switch a1 := os.Args[1][:2]; a1 {
	case "-h":
		help("")
	case "-c":
		a.Create()
	case "-l":
		a.LoadAPODarchive()
	case "-t":
		if len(os.Args) < 3 {
			help("Missing -t date argument in format yymmdd")
			os.Exit(1)
		}
		yymmdd := os.Args[2]
		if title, err := a.SearchTitle(yymmdd); err != nil {
			fmt.Println(yymmdd, "error:", err.Error())
			os.Exit(1)
		} else {
			fmt.Println(yymmdd, ":", title)
		}
	default:
		help("Not enough arguments")
		os.Exit(1)
	}

}
