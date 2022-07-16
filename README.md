# Go Generic lib for generic iteration

Iterate, modify and transform structs, arrays and maps in generic way in Go!

## Example

```go
type Person struct {
	FirstName  string
	LastName   string
	Colleagues []Person
}

// Person equals person where thair first and last name are the same
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
```

which prints:

```console
$ go run ./example/example_1
KEN THOMPSON
  knows Dennis Ritchie
  knows Brian Kernighan
DENNIS RITCHIE
  knows Ken Thompson
  knows Brian Kernighan
BRIAN KERNIGHAN
  knows Ken Thompson
  knows Dennis Ritchie
```

This library is inspired by Haskell [`Data.Data`](https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Data.html) and [`Data.Generics`](https://hackage.haskell.org/package/syb-0.7.2.1/docs/Data-Generics.html). See [this video](https://www.youtube.com/watch?v=Zj8KXD9MRA0) for short introduction. It was created to play with this idea in other language.
