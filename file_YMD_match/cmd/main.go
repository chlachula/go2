package main

import (
	"fmt"
	"os"
	"time"

	a "github.com/chlachula/go2/file_YMD_match"
)

func help() {
	helptext := `
Usage:
go2 sourcePath targetDir 
Example:
go2  /e/Arc-Pics/2003/2003_06_22Ne/BoziTelo22 020.jpg /e/Arc-Pics/Zpracovat
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	if len(os.Args) < 3 {
		help()
		os.Exit(1)
	}
	fmt.Printf("%t\n", a.ExistsYMD(os.Args[1], os.Args[2]))
}
