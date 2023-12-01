package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/diffsecs"
)

func str2date(s string) (time.Time, error) {
	date, err := time.Parse("2006.1.2", s)
	return date, err
}
func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Date difference including leap seconds 
Usage:
go2 -h
go2 -d yyyy.mm.dd1 yyyy.mm.dd2
Examples:
go2 -d 2016.12.31 2017.1.1
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
	i := 1
	if strings.HasPrefix(os.Args[i], "-h") {
		help("")
		a.ShowLeapSeconds()
		os.Exit(0)
	} else if strings.HasPrefix(os.Args[i], "-v") {
		a.Verbose = true
	} else if strings.HasPrefix(os.Args[i], "-d") {
		if len(os.Args) < 4 {
			help("Less than 2 expected date arguments")
			os.Exit(1)
		}
		var d1, d2 time.Time
		var err error
		if d1, err = str2date(os.Args[i+1]); err != nil {
			help("1st date argument error: " + err.Error())
			os.Exit(1)
		}
		if d2, err = str2date(os.Args[i+2]); err != nil {
			help("2nd date argument error: " + err.Error())
			os.Exit(1)
		}
		a.ShowDatesDiffInSeconds(d1, d2)
	} else {
		help("Unknown argument " + os.Args[1])
		os.Exit(1)
	}

}
