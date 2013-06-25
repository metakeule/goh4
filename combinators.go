package goh4

import (
	"strings"
)

type combinator struct {
	selectors []Selecter
	operator  string
}

func (ø combinator) Selector() string {
	s := []string{}
	for _, sel := range ø.selectors {
		s = append(s, sel.Selector())
	}
	return strings.Join(s, ø.operator)
}

func (ø combinator) Add(s Selecter) SelecterAdder {
	return combinator{append(ø.selectors, s), ø.operator}
}

// F element descendant of an E element
func Descendant(selectors ...Selecter) combinator { return combinator{selectors, " "} }

// F element child of an E element
func Child(selectors ...Selecter) combinator { return combinator{selectors, " > "} }

// F element immediately preceded by an E element
func DirectFollows(selectors ...Selecter) combinator { return combinator{selectors, " + "} }

// F element preceded by an E element
func Follows(selectors ...Selecter) combinator { return combinator{selectors, " ~ "} }

// for each given selector the rules apply
func Each(selectors ...Selecter) combinator { return combinator{selectors, ",\n"} }
