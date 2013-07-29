package goh4

import (
	"strings"
)

// something that matches an Element
type Matcher interface {
	Matches(*Element) bool
}

type or []Matcher

// matches if any of the Matchers matches
func (ø or) Matches(e *Element) bool {
	for _, m := range ø {
		if m.Matches(e) {
			return true
		}
	}
	return false
}

func Or_(m ...Matcher) or {
	return or(m)
}

type and []Matcher

// matches if all of the Matchers matches
func (ø and) Matches(e *Element) bool {
	for _, m := range ø {
		if !m.Matches(e) {
			return false
		}
	}
	return true
}

func And_(m ...Matcher) and {
	return and(m)
}

type not struct{ Matcher }

// matches if inner matcher does not match
func (ø *not) Matches(e *Element) bool { return !ø.Matcher.Matches(e) }

func Not_(m Matcher) *not { return &not{m} }

type PositionMatcher struct {
	Element *Element
	Pos     int
	Found   bool
}

func (ø *PositionMatcher) Matches(e *Element) (f bool) {
	// no recursive findings
	if e.Parent() != ø.Element.Parent() {
		return
	}
	if ø.Element == e {
		ø.Found = true
		return true
	}
	if !ø.Found {
		ø.Pos += 1
	}
	return
}

type FieldMatcher int

func (ø FieldMatcher) Matches(t *Element) (m bool) { return t.Is(FormField) }

func (ø Class) Matches(t *Element) bool {
	for _, c := range t.classes {
		if c == ø {
			return true
		}
	}
	return false
}

// if Css has a Context, matching always fails
/*
func (ø *Css) Matches(t *Element) (m bool) {
	if ø.Context != "" {
		// if Css has a Context, matching always fails
		return false
	}
	if ø.class != "" {
		if !t.HasClass(ø.class) {
			return false
		}
	}

	if len(ø.Tags) > 0 {
		for _, tt := range ø.Tags {
			if ø.matchTag(t, Tag(tt)) {
				return true
			}
		}
	} else {
		return true
	}

	return
}
*/

func (ø Id) Matches(t *Element) bool {
	if string(t.id) == string(ø) {
		return true
	}
	return false
}

// matching an html string, ignoring whitespace
func (ø Html) Matches(t *Element) bool {
	inner := removeWhiteSpace(t.InnerHtml())
	me := removeWhiteSpace(ø.String())
	if inner == me {
		return true
	}
	return false
}

func (ø Tag) Matches(t *Element) bool {
	return ø == t.tag
}

func (ø Style) Matches(t *Element) bool {
	return t.style[ø.Property] == ø.Value
}

/*
func (ø styles) Matches(t *Element) (m bool) {
	m = true
	for _, styl := range ø {
		if !styl.Matches(t) {
			m = false
		}
	}
	return
}
*/

func (ø SingleAttr) Matches(t *Element) bool {
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
			styl := Style{a[0], a[1]}
			if !styl.Matches(t) {
				m = false
			}
		}
		return m
	}
	if t.attributes[ø.Key] == ø.Value {
		return true
	}
	return false
}

func (ø Attrs) Matches(t *Element) (m bool) {
	m = true
	for _, attr := range ø {
		if !attr.Matches(t) {
			m = false
		}
	}
	return
}
