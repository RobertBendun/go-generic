package main

import (
	"fmt"
	"generic"
	"math/rand"
	"strings"
	"time"
)

type Vec2 struct {
	X, Y int
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

type Person struct {
	FirstName string
	LastName  string
	Position  Vec2
	Flags     []bool
}

func prefixPrint(prefix string) func(interface{}) {
	return func(value interface{}) {
		generic.EachWithKey(value, func(key interface{}, value interface{}) {
			if !generic.IsAnyIterable(value) || generic.IsAnyString(value) {
				fmt.Printf("%s.%v = %v\n", prefix, key, value)
				return
			}
			prefixPrint(fmt.Sprintf("%v.%v", prefix, key))(value)
		})
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	person := Person{
		FirstName: "Barbara",
		LastName:  "Liskov",
		Position:  generic.New[Vec2](func() int { return 10 }),
		Flags:     generic.NewArray(4, func() bool { return rand.Intn(2) != 0 }),
	}

	generic.Modify(&person, func(s *string) {
		*s = strings.ToUpper(*s)
	})

	generic.EachWithKey(&person, func(key string, value interface{}) {
		fmt.Printf("person.%s = %v\n", key, value)
	})

	prefixPrint("person")(person)
}
