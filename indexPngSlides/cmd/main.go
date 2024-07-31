package main

import (
	"fmt"
	"os"
	"time"

	a "github.com/chlachula/go2/indexPngSlides"
)

func help(msg string) {
	helpText := `
Program indexPngSlides put not tall PNG slide images to the index.html page
Usage:
 go2 directory //containing png files
Example:
 go2 testdir
`
	if msg != "" {
		fmt.Println(msg)
	}
	fmt.Println(helpText)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	if len(os.Args) < 2 {
		help("No directory argument")
		os.Exit(1)
	}
	dirWithPNGs := os.Args[1]
	a.CreateIndexFile(dirWithPNGs)
}
