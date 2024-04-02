package randomgroups

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

var version string = "1.0.0"
var verbose bool = true
var DemoGroup = []Person{
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

type Person struct {
	Nick string
	Name string
}
type Workgroup []Person

func printVerbose(s string) {
	if verbose {
		print(s)
	}
}
func removeIndex(s1 []Person, index int) []Person {
	s2 := make([]Person, 0)
	s2 = append(s2, s1[:index]...)
	return append(s2, s1[index+1:]...)
}
func createRandomWorkGroups(mainGroup []Person, groupSize int) []Workgroup {
	var group []Person = mainGroup // copy of the mainGroup

	// determine number of groups and create them
	numberOfGroups := len(group) / groupSize
	var workgroups []Workgroup = make([]Workgroup, numberOfGroups)
	for i := 0; i < numberOfGroups; i++ {
		workgroups[i] = make(Workgroup, groupSize)
	}
	modulo := len(group) % groupSize
	if modulo != 0 {
		numberOfGroups += 1
		workgroups = append(workgroups, make(Workgroup, modulo))
	}
	printVerbose(fmt.Sprintf("people count = %d, maxSize = %d, numberOfGroups = %d, modulo = %d\n", len(group), groupSize, numberOfGroups, modulo))

	// randomly populate groups
	for i := 0; i < numberOfGroups; i++ {
		for j := 0; j < len(workgroups[i]); j++ {
			leng := len(group)
			randomIndex := rand.IntN(leng)
			printVerbose(fmt.Sprintf("i = %d, j = %d, len(group) = %d\n", i, j, len(group)))
			workgroups[i][j] = group[randomIndex]
			group = removeIndex(group, randomIndex)
		}
	}

	return workgroups
}
func ShowRandomWorkGroups(mainGroup []Person, groupSize int) {
	w := createRandomWorkGroups(mainGroup, groupSize)
	fmt.Printf("Group of %d people randomly divided into %d groups of %d-%d people\n", len(mainGroup), len(w), groupSize, groupSize+1)
	fmt.Printf("Generated on %s by program version %s\n", time.Now().Format(time.RFC3339), version)
	for i := 0; i < len(w); i++ {
		fmt.Printf("%d. Group:\n", i+1)
		members := w[i]
		for j, member := range members {
			fmt.Printf("     %d. %s - %s\n", (j + 1), member.Nick, member.Name)
		}
	}

}
func LoadGroup(filename string) ([]Person, error) {
	bytes, err := os.ReadFile(filename) //Read entire file content. No need to close
	if err != nil {
		return nil, err
	}
	var group []Person
	if err = json.Unmarshal(bytes, &group); err == nil {
		return group, nil
	} else {
		return nil, err
	}
}
func SplitRandomlyPeopleInFile(filename string, size int) {
	if people, err := LoadGroup(filename); err == nil {
		ShowRandomWorkGroups(people, size)
	} else {
		fmt.Println(err.Error())
	}
}
