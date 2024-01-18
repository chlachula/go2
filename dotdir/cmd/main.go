package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/dotdir"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	text := `Program dotdir shows absolute path of the current directory.
	Current directory is represented by single '.' .
	Program shows also content of the directory on stdout and on webpage /dot .text
	It can be usefull to explore unknown environment like containers.
	Usage:
	no arguments .. display info on stdout and starts webserver on port 8080
	-h this help
	-p PORT_NUMBER .. defines webserver port number`
	fmt.Println(text)
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	colonPort := ":8080"
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "-h") {
			help("")
			return
		}
		if strings.HasPrefix(os.Args[1], "-p") && len(os.Args) > 2 {
			colonPort = ":" + os.Args[2]
		}
	}

	fmt.Println(a.DirAbsInfo("."))
	a.Web(colonPort)
}
