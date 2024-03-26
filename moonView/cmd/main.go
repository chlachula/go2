package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	a "github.com/chlachula/go2/moonView"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Moon views:
	Based on 730 px images from NASA SVS with 1 hour steps
	-h this help
	-p [port] preview to port 8080 
	`
	fmt.Println(helptext)
}
func main() {
	argsWithoutProg := os.Args[1:]
	http.HandleFunc("/", a.EventHandler)
	colonPort := ":8080"

	if len(argsWithoutProg) == 0 {
		help("not enough arguments")
		os.Exit(1)
	}

	if argsWithoutProg[0] == "-p" {
		if len(argsWithoutProg) > 1 {
			colonPort = ":" + argsWithoutProg[1]
		}
		print("...listening at " + colonPort)
		log.Fatal(http.ListenAndServe(colonPort, nil))
	} else if argsWithoutProg[0] == "-h" {
		help("")
	} else {
		help("unknown argument " + argsWithoutProg[0])
		os.Exit(1)
	}
}
