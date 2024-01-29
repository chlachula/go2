package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/dirreport"
)

func underscoredText(i int) string {
	if i < len(os.Args) {
		return strings.ReplaceAll(os.Args[i], "_", " ")
	} else {
		help("missing bottom text argument")
		os.Exit(1)
		return ""
	}
}
func help(msg string) {
	if msg != "" {
		fmt.Printf("%s \n\n", msg)
	}
	helptext := `Web displaying directory
 Usage:
 -h this help
 Example:
`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	help("")
	a.Dir = "hello"
	colonPort := ":8080"
	http.HandleFunc("/", a.ShowDir)
	print("Serving and listenning at " + colonPort + ". CTRL+C to stop.")
	http.ListenAndServe(colonPort, nil)
}
