package goh4

import (
	"fmt"
	"html"
	"strings"
)

// here something special
type Stringer interface {
	String() string
}

// something that matches an Element
type Matcher interface {
	Matches(*Element) bool
}

type FieldMatcher int

func (ø FieldMatcher) Matches(t *Element) (m bool) {
	return t.field
}

type Context string

func (ø Context) String() string { return string(ø) }

type Comment string

func (ø Comment) String() string { return string(ø) }

type Class string

func (ø Class) String() string { return string(ø) }
func (ø Class) Matches(t *Element) bool {
	for _, c := range t.classes {
		if string(c) == string(ø) {
			return true
		}
	}
	return false
}

type Id string

func (ø Id) String() string { return string(ø) }
func (ø Id) Matches(t *Element) bool {
	if string(t.id) == string(ø) {
		return true
	}
	return false
}

type Text string

func (ø Text) String() string { return string(ø) }

type Html string

func (ø Html) String() string { return string(ø) }

// matching an html string, ignoring tabs and linefeeds
func (ø Html) Matches(t *Element) bool {
	inner := strings.Replace(t.InnerHtml(), "\n", "", -1)
	inner = strings.Replace(inner, "\t", "", -1)
	me := strings.Replace(ø.String(), "\n", "", -1)
	me = strings.Replace(me, "\t", "", -1)
	if inner == me {
		return true
	}
	return false
}

type Tag string

func (ø Tag) String() string { return string(ø) }

func (ø Tag) Matches(t *Element) bool {
	return ø == t.tag
}

type Tags []string

func (ø Tags) ToTagArr() (a []Tag) {
	a = []Tag{}
	for _, s := range ø {
		a = append(a, Tag(s))
	}
	return
}

func (ø Tags) String() string {
	return strings.Join(ø, ", ")
}

type Style struct {
	Key   string
	Value string
}

func (ø Style) String() string { return ø.Key + ": " + ø.Value + ";" }

func (ø Style) Matches(t *Element) bool {
	return t.style[ø.Key] == ø.Value
}

// helper to easily create multiple styles
// use is like this
// Styles{"width","200px","height","30px","color","green"}.ToStyleArr()
type Styles []string

func (ø Styles) ToStyleArr() (s []Style) {
	s = []Style{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, Style{ø[i], ø[i+1]})
	}
	return
}

func (ø Styles) Matches(t *Element) (m bool) {
	styles := ø.ToStyleArr()
	m = true
	for _, style := range styles {
		if !style.Matches(t) {
			m = false
		}
	}
	return
}

func (ø Styles) String() (s string) {
	ss := []string{}
	styles := ø.ToStyleArr()
	for _, style := range styles {
		ss = append(ss, style.String())
	}
	return strings.Join(ss, " ")
}

type Attr struct {
	Key   string
	Value string
}

func (ø Attr) Matches(t *Element) bool {
	if ø.Key == "id" {
		return Id(ø.Value).Matches(t)
	}
	if ø.Key == "class" {
		return Class(ø.Value).Matches(t)
	}
	if ø.Key == "style" {
		styles := strings.Split(ø.Value, ";")
		m := true
		for _, st := range styles {
			a := strings.Split(st, ":")
			style := Style{a[0], a[1]}
			if !style.Matches(t) {
				m = false
			}
		}
		return m
	}
	if t.Attributes[ø.Key] == ø.Value {
		return true
	}
	return false
}

func (ø Attr) String() string { return " " + ø.Key + `="` + html.EscapeString(ø.Value) + `"` }

// helper to easily create multiple attrs
// use is like this
// Attrs{"width","200px","height","30px","value","hiho"}.ToAttrArr()
type Attrs []string

func (ø Attrs) ToAttrArr() (s []Attr) {
	s = []Attr{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, Attr{ø[i], ø[i+1]})
	}
	return
}

func (ø Attrs) Matches(t *Element) (m bool) {
	attrs := ø.ToAttrArr()
	m = true
	for _, attr := range attrs {
		if !attr.Matches(t) {
			m = false
		}
	}
	return
}

func (ø Attrs) String() (s string) {
	ss := []string{}
	attrs := ø.ToAttrArr()
	for _, attr := range attrs {
		ss = append(ss, attr.String())
	}
	return strings.Join(ss, " ")
}

func panicf(message string, obj ...interface{}) {
	panic(fmt.Sprintf(message+"\n", obj...))
}
