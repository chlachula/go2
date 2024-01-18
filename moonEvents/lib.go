package moonEvents

import (
	"encoding/json"
	"os"
)

func LoadLocations(filename string) ([]LocationInfo, error) {
	var locations []LocationInfo
	if bytes, err := os.ReadFile(filename); err != nil { //Read entire file content. No need to close
		return locations, err
	} else {
		if err = json.Unmarshal(bytes, &locations); err != nil {
			return locations, err
		} else {
			return locations, nil
		}
	}
}
func LoadEventsYear(filename string) (EventsYearType, error) {
	var eventsYear EventsYearType
	if bytes, err := os.ReadFile(filename); err != nil { //Read entire file content. No need to close
		return eventsYear, err
	} else {
		if err = json.Unmarshal(bytes, &eventsYear); err != nil {
			return eventsYear, err
		} else {
			return eventsYear, nil
		}
	}
}
func LoadTextFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename) //Read entire file content. No need to close
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
