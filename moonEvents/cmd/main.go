package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	a "github.com/chlachula/go2/moonEvents"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Events calendar with moon phases background
	Loads data from default files locations.json and events.json
	Usage:
	-h this help
	-p [port] preview to port 8080 
	-c generate google calendar  
	`
	fmt.Println(helptext)
}
func main() {
	a.LoadData()
	sort.Sort(a.ByDate(a.EventsData.Event))

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		help("")
	}
	if argsWithoutProg[0] == "-c" {
		filename := fmt.Sprintf("GoogleCal-events-%d.csv", a.EventsData.YYYY)
		a.GenerateGoogleCalendarListCSV(a.EventsData, filename)
	} else if argsWithoutProg[0] == "-p" {
		http.HandleFunc("/", a.EventHandler)
		http.HandleFunc("/api", a.ApiHandler)
		http.Handle("/files", http.FileServer(http.Dir("public/")))
		colonPort := ":8080"
		if len(argsWithoutProg) > 1 {
			colonPort = ":" + argsWithoutProg[1]
		}
		print("...listening at " + colonPort)
		log.Fatal(http.ListenAndServe(colonPort, nil))
	}
}
