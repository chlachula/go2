package main

import (
	"fmt"
	"time"

	a "github.com/chlachula/go2/moonEvents"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Events calendar with moon phases background
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	help("")
	a.Hello()
}
