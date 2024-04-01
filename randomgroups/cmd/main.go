package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	a "github.com/chlachula/go2/randomgroups"
)

var groupSize int = 2
var groupsNumber int

func PrettyJsonStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
func printDemoGroupJson() {
	if str, err := PrettyJsonStruct(a.DemoGroup); err == nil {
		fmt.Printf("Demo json file example:\n\n%s\n\n", str)
	} else {
		fmt.Printf("Error:\n %s\n", err.Error())
	}
}
func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Splitting into number of groups
	Usage:
	-h this help
	-d demo
	-f filename         #json
	-g group-max-size   #default is 2 
	Examples:
	-d 
	-g 4 -f people.json
	`
	fmt.Println(helptext)
}
func main() {

	if len(os.Args) < 2 {
		help("not enought arguments")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		switch arg := os.Args[i]; arg {
		case "-h":
			help("")
			os.Exit(0)
		case "-d":
			printDemoGroupJson()
			a.ShowRandomWorkGroups(a.DemoGroup, groupSize)
			os.Exit(0)
		case "-g":
			i += 1
			if i >= len(os.Args) {
				help("not -g argument")
				os.Exit(1)
			}
			groupSize, _ = strconv.Atoi(os.Args[i])
		case "-f":
			i += 1
			if i >= len(os.Args) {
				help("not -g argument")
				os.Exit(1)
			}
			a.SplitRandomlyPeopleInFile(os.Args[i], groupSize)
			os.Exit(0)
		default:
			help("unexpected argument " + arg)
		}
	}
}
