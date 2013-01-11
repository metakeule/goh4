package goh4

import "strings"

// something that matches an Element
type Matcher interface {
	Matches(*Element) bool
}

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

func (ø FieldMatcher) Matches(t *Element) (m bool) { return t.Is(Field) }

func (ø Class) Matches(t *Element) bool {
	for _, c := range t.classes {
		if c == ø {
			return true
		}
	}
	return false
}

// Bug(m) if Css has a Context, matching always fails
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
	return t.style[ø.Key] == ø.Value
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
	if t.attributes[ø.Key] == ø.Value {
		return true
	}
	return false
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
