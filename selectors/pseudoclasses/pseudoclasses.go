package pseudoclasses

import (
	"fmt"
)

type pseudo string

type Selecter interface {
	Selector() string
}

func (ø pseudo) Selector() string { return ":" + string(ø) }

// root of the document
func Root() pseudo { return pseudo("root") }

// the n-th child of its parent
func NthChild(n int) pseudo { return pseudo(fmt.Sprintf("nth-child(%d)", n)) }

// the n-th child of its parent, counting from the last one
func NthLastChild(n int) pseudo { return pseudo(fmt.Sprintf("nth-last-child(%d)", n)) }

// the n-th sibling of its type
func NthOfType(n int) pseudo { return pseudo(fmt.Sprintf("nth-of-type(%d)", n)) }

// first child of its parent
func FirstChild() pseudo { return pseudo("first-child") }

// last child of its parent
func LastChild() pseudo { return pseudo("last-child") }

// first sibling of its type
func FirstOfType() pseudo { return pseudo("first-of-type") }

// last sibling of its type
func LastOfType() pseudo { return pseudo("last-of-type") }

// only child of its parent
func OnlyChild() pseudo { return pseudo("only-child") }

// only sibling of its type
func OnlyOfType() pseudo { return pseudo("only-of-type") }

// has no children (including text nodes)
func Empty() pseudo { return pseudo("empty") }

// is the source anchor of a hyperlink of which the target is not yet visited
func Link() pseudo { return pseudo("link") }

// is the source anchor of a hyperlink of which the target is already visited (:visited)
func Visited() pseudo { return pseudo("visited") }

// E:active
func Active() pseudo { return pseudo("active") }

// E:hover
func Hover() pseudo { return pseudo("hover") }

// during certain user actions
func Focus() pseudo { return pseudo("focus") }

// being the target of the referring URI
func Target() pseudo { return pseudo("target") }

// in language l
func Lang(l string) pseudo { return pseudo(fmt.Sprintf("lang(%s)", l)) }

// E:enabled
func Enabled() pseudo { return pseudo("enabled") }

// is disabled
func Disabled() pseudo { return pseudo("disabled") }

// is checked (for instance a radio-button or checkbox)
func Checked() pseudo { return pseudo("checked") }

// the first formatted line
func FirstLine() pseudo { return pseudo(":first-line") }

// the first formatted letter
func FirstLetter() pseudo { return pseudo("::first-letter") }

// generated content before
func Before() pseudo { return pseudo("::before") }

// generated content after
func After() pseudo { return pseudo("::after") }

// does not match simple selector
func Not(selector Selecter) pseudo { return pseudo(fmt.Sprintf("not(%s)", selector.Selector())) }
