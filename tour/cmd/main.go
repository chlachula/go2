package main

import (
	"fmt"
	"time"

	a "github.com/chlachula/go2/tour"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Brief go language tour
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	help("")
	a.SystemInt()
	a.IntMinMax()
	a.FloatMinMax()
	a.GlobalVariables()
	a.Loops()
	a.Switches()

}
