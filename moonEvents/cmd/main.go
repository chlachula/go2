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
	-l locations-file.json
	-e events-file.json
	-c generate google calendar  
  Examples:
	go2 -p 
	go2 -l locs.json -e ev2024.json -p 8081
	go2 -l locs.json -e ev2024.json -c	
`
	fmt.Println(helptext)
}
func main() {
	locationsJson := "locations.json"
	eventsJson := "events.json"
	colonPort := ":8080"
	startWeb := true

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		help("")
		os.Exit(1)
	}
	for i := 0; i < len(argsWithoutProg); i++ {
		switch arg := argsWithoutProg[i]; arg {
		case "-h":
			help("")
			os.Exit(0)
		case "-l":
			i += 1
			locationsJson = argsWithoutProg[i]
		case "-e":
			i += 1
			eventsJson = argsWithoutProg[i]
		case "-c":
			startWeb = false
		case "-p":
			startWeb = true
			if len(argsWithoutProg) > i+1 {
				colonPort = ":" + argsWithoutProg[i+1]
			}
		default:
			fmt.Printf("Unexpected argument %s\n", arg)
		}
	}

	a.LoadData(locationsJson, eventsJson)
	sort.Sort(a.ByDate(a.EventsData.Event))

	if startWeb {
		http.HandleFunc("/", a.EventHandler)
		http.HandleFunc("/api", a.ApiHandler)
		http.Handle("/files", http.FileServer(http.Dir("public/")))
		print("...listening at " + colonPort)
		log.Fatal(http.ListenAndServe(colonPort, nil))
	} else {
		filename := fmt.Sprintf("GoogleCal-events-%d.csv", a.EventsData.YYYY)
		a.GenerateGoogleCalendarListCSV(a.EventsData, filename)
	}
}
