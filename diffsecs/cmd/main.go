package main

import (
	"fmt"
	"os"
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
go2 yyyy.mm.dd1 yyyy.mm.dd2
Examples:
go2 2016.12.31 2017.1.1
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	if len(os.Args) < 3 {
		help("Less than 2 expected date arguments")
		os.Exit(1)
	}
	var d1, d2 time.Time
	var err error
	if d1, err = str2date(os.Args[1]); err != nil {
		help("1st argument error: " + err.Error())
		os.Exit(1)
	}
	if d2, err = str2date(os.Args[2]); err != nil {
		help("1st argument error: " + err.Error())
		os.Exit(1)
	}
	a.ShowLeapSeconds()
	a.SecondsDiff(d1, d2)
}
