package main

import (
	a "github.com/chlachula/go2/randomgroups"
)

var mainGroup = []a.Person{
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

func main() {
	groupSize := 3
	a.ShowRandomWorkGroups(mainGroup, groupSize)
}
