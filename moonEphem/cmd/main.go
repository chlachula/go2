package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	a "github.com/chlachula/go2/moonEphem"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Moon and Sun Ephemerides
Usage:
go2 date step number
Examples:
# Print JPL Horizons sample ephemerids
go2 e
# Sun Jan 1st 2023, step 10 days, 25 steps
go2 s 2023-01-01 10 25
# Moon Aug 29th 2025, step 1 day, 30 steps
go2 moo 2025-08-29 1 30
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	var body string
	if len(os.Args) > 1 {
		body = strings.ToLower(os.Args[1])
	}
	if strings.HasPrefix(body, "e") {
		a.PrintTestExampleEphems()
		return
	}

	if len(os.Args) < 5 {
		help("Less than 4 expected arguments")
		os.Exit(1)
	}

	DateLayout := "2006-01-02"
	date, err := time.Parse(DateLayout, os.Args[2])
	if err != nil {
		help("Error: " + err.Error())
		os.Exit(1)
	}
	stepDays, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		help("Error: " + err.Error())
		os.Exit(1)
	}
	stepsNumber, err := strconv.Atoi(os.Args[4])
	if err != nil {
		help("Error: " + err.Error())
		os.Exit(1)
	}

	if strings.HasPrefix(body, "m") {
		a.MoonEphemerides(date, stepDays, stepsNumber)
	} else if strings.HasPrefix(body, "s") {
		a.SunEphemerides(date, stepDays, stepsNumber)
	} else {
		help(fmt.Sprintf("The first letter of a celestial body '%s' corresponds to neither the Sun nor the Moon ", body))
		os.Exit(1)
	}
}
