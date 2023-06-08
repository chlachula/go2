package main

import (
	"fmt"
	"runtime"
	"time"

	a "github.com/chlachula/go2/windrives"
)

func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	if runtime.GOOS != "windows" {
		fmt.Println("This program is suited only for Windows OS!")
		return
	}
	fmt.Printf("Available letters: %v\n", a.WinAvailableLetterDrives())
}
