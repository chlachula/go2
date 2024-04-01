package randomgroups

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

var version string = "1.0.0"

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

func removeIndex(s1 []Person, index int) []Person {
	s2 := make([]Person, 0)
	s2 = append(s2, s1[:index]...)
	return append(s2, s1[index+1:]...)
}
func createRandomWorkGroups(mainGroup []Person, groupSize int) []Workgroup {
	var group []Person = mainGroup
	var workgroups []Workgroup = make([]Workgroup, 0)

	number := len(group) / groupSize
	for i := 0; i < number; i++ {
		var workgroup = make(Workgroup, groupSize)
		for j := 0; j < groupSize; j++ {
			randomIndex := rand.IntN(len(group))
			workgroup[j] = group[randomIndex]
			group = removeIndex(group, randomIndex)
		}
		workgroups = append(workgroups, workgroup)
	}
	modulo := len(group) % groupSize
	for i := 0; i < modulo; i++ {
		randomIndex := rand.IntN(len(group))
		workgroups[i] = append(workgroups[i], group[randomIndex])
		group = removeIndex(group, randomIndex)
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
		for _, member := range members {
			fmt.Printf("     %s - %s\n", member.Nick, member.Name)
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
