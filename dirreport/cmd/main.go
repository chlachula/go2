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
 -d dir-name #set directory to show
 -i include also hidden subdirs starting with dot .
 -p [port] #post to start web, default is 8080
 -h #this help
 Example:
 go2 -d /home/users/joedoe -p 8081
`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	colonPort := ":8080"
	if len(os.Args) < 2 {
		help("Not enough arguments")
	} else {
		for i := 1; i < len(os.Args); i++ {
			switch arg := os.Args[i]; arg {
			case "-h":
				help("")
				os.Exit(0)
			case "-d":
				i += 1
				a.Dir = os.Args[i]
			case "-i":
				a.ExcludeDotDirs = false
			case "-p":
				a.SetDirInf()
				http.HandleFunc("/", a.HandleHome)
				http.HandleFunc("/show-dir", a.HandleShowDir)
				http.HandleFunc("/show-dir2", a.HandleShowDir2)
				http.HandleFunc("/show-file", a.HandleShowFile)
				print("Serving and listenning at " + colonPort + ". CTRL+C to stop.")
				http.ListenAndServe(colonPort, nil)
			default:
				help("Unexpected argument " + arg)
			}
		}
	}

}
