package combinators

import "fmt"

type Selecter interface {
	Selector() string
}

type combinator struct {
	first    Selecter
	operator string
	last     Selecter
}

func (ø combinator) Selector() string {
	if ø.operator == "" {
		return fmt.Sprintf(`%s %s`, ø.first.Selector(), ø.last.Selector())
	}

	return fmt.Sprintf(`%s %s %s`, ø.first.Selector(), ø.operator, ø.last.Selector())
}

// F element descendant of an E element
func Descendant(first Selecter, last Selecter) combinator { return combinator{first, "", last} }

// F element child of an E element
func Child(first Selecter, last Selecter) combinator { return combinator{first, ">", last} }

// F element immediately preceded by an E element
func DirectFollows(first Selecter, last Selecter) combinator { return combinator{first, "+", last} }

// F element preceded by an E element
func Follows(first Selecter, last Selecter) combinator { return combinator{first, "~", last} }
