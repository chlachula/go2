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
	helptext := `Events calendar with moon phases background
	-h this help
	-p [port] preview to port 8080 
	-c generate google calendar  
	`
	fmt.Println(helptext)
}
func main() {
	argsWithoutProg := os.Args[1:]
	http.HandleFunc("/", a.EventHandler)
	colonPort := ":8080"
	if argsWithoutProg[0] == "-p" {
		if len(argsWithoutProg) > 1 {
			colonPort = ":" + argsWithoutProg[1]
		}
		print("...listening at " + colonPort)
		log.Fatal(http.ListenAndServe(colonPort, nil))
	}
}
