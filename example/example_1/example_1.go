package main

import (
	"fmt"
	"generic"
	"strings"
)

type Person struct {
	FirstName  string
	LastName   string
	Colleagues []Person
}

func (lhs Person) Equal(rhs Person) bool {
	return generic.All(lhs, rhs, func(lhs string, rhs string) bool {
		return lhs == rhs
	})
}

func main() {
	unixTeam := []Person{
		{FirstName: "Ken", LastName: "Thompson"},
		{FirstName: "Dennis", LastName: "Ritchie"},
		{FirstName: "Brian", LastName: "Kernighan"},
	}

	// Set all colleagues of person to people in unixTeam except that person
	generic.Modify(unixTeam, func(p *Person) {
		for _, other := range unixTeam {
			if !p.Equal(other) {
				p.Colleagues = append(p.Colleagues, other)
			}
		}
	})

	generic.Each(unixTeam, func(person Person) {
		// Get person with
		person = generic.Map(person, strings.ToUpper)

		fmt.Printf("%s %s\n", person.FirstName, person.LastName)
		generic.Each(person.Colleagues, func(colleague Person) {
			fmt.Printf("  knows %s %s\n", colleague.FirstName, colleague.LastName)
		})
	})
}
