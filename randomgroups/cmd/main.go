package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	a "github.com/chlachula/go2/randomgroups"
)

var group = []a.Person{}
var groupSize int = 2

var demoGroup = []a.Person{
	{Nick: "Ann1", Name: "Ann Anderson"},
	{Nick: "Bob2", Name: "Bob Brown"},
	{Nick: "Chuck3", Name: "Charles Jones"},
	{Nick: "Dave4", Name: "David Davis"},
	{Nick: "Eve5", Name: "Evelyne Smiths"},
	{Nick: "Fred6", Name: "Frederick Miller"},
	{Nick: "Greg7", Name: "Gregory Garcia"},
	{Nick: "Helen8", Name: "Helen Rodriguez"},
	{Nick: "Ian9", Name: "Ian Wilson"},
	{Nick: "Joe10", Name: "Joseph Williams"},
	{Nick: "Kam11", Name: "Kamila Thompson"},
	{Nick: "Lil12", Name: "Lillian Lopez"},
	{Nick: "Mat13", Name: "Mathew Johnson"},
	{Nick: "Neil14", Name: "Neil Young"},
}
var demoGroupSize int = 3

func PrettyJsonStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
func printDemoGroupJson() {
	if str, err := PrettyJsonStruct(demoGroup); err == nil {
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
	-f filename     #line oriented
	-g group-number #default is 2
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
			a.ShowRandomWorkGroups(demoGroup, demoGroupSize)
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
