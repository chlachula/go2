package main

import (
	"fmt"
        "time"
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
}
