package main

import (
	"fmt"
	"github.com/metakeule/fat"
	. "github.com/metakeule/goh4/fat"
	. "github.com/metakeule/goh4/tag"
)

type Person struct {
	FirstName *fat.Field `fat.type:"string" goh4.type:"text"`
	LastName  *fat.Field `fat.type:"string"`
	Vita      *fat.Field `fat.type:"string" goh4.type:"html"`
}

var PERSON = fat.Proto(&Person{}).(*Person)

func NewPerson() *Person { return fat.New(PERSON, &Person{}).(*Person) }

func init() {
	Register(PERSON)
}

func main() {

	details := UL("\n",
		LI(Placeholder(PERSON.FirstName)), "\n",
		LI(Placeholder(PERSON.LastName)), "\n",
		LI(Placeholder(PERSON.Vita)), "\n",
	).Compile("details")

	paul := NewPerson()
	paul.FirstName.Set("<Pa>ul")
	paul.LastName.Set("Pa<n>zer")
	paul.Vita.Set("<p>hier die vita</p>")

	fmt.Println(details.MustReplace(Setters(paul)...).String())
}
