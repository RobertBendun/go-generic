# Go Generic - generic iteration library

Iterate, modify and transform structs, arrays and maps in generic way in Go, with type of function determining which fields can be accessed.

For functions used for iteration `func(arg T)`:

- when `T` is a __concrete type__ like `int`, `[]float`, `struct {}` function is beeing called only when type of field is convertible to `T`,
- when `T` is an __interface__ then function is beeing called only for those values (fields, elements) that implements that interface.

Rules above means that generic iteration function, that will traverse all struct fields or map and slice values is `func(value interface{})`

Situation is the same for `func(key K, val V)`, for `K` beeing convertible from or implementing `string` for structs, `int` for slices and `T` for `map[T]_` 

## Example

For more examples see [example](example/) directory.

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
		{FirstName: "Ken",    LastName: "Thompson"},
		{FirstName: "Dennis", LastName: "Ritchie"},
		{FirstName: "Brian",  LastName: "Kernighan"},
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
		// Make person first and last name uppercase for nice heading
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

## See also

This library is inspired by Haskell [`Data.Data`](https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Data.html) and [`Data.Generics`](https://hackage.haskell.org/package/syb-0.7.2.1/docs/Data-Generics.html). See [this video](https://www.youtube.com/watch?v=Zj8KXD9MRA0) for short introduction.
