package goh4

import (
	"testing"
)

func TestPositionMatcher(t *testing.T) {
	e := NewElement(Tag("p"))
	m := &PositionMatcher{Element: e}

	if m.Matches(NewElement(Tag("a"))) || m.Found {
		err(t, "incorrect position matcher matches", true, false)
	}

	_ = m.Matches(NewElement(Tag("a")))

	if !m.Matches(e) || !m.Found {
		err(t, "incorrect position matcher matches", false, true)
	}

	_ = m.Matches(NewElement(Tag("a")))

	if m.Pos != 2 {
		err(t, "incorrect position matcher position first round", m.Pos, 2)
	}
}

func TestFieldMatcher(t *testing.T) {
	var m FieldMatcher
	if !m.Matches(NewElement(Tag("input"), FormField)) {
		err(t, "incorrect field matcher matches", false, true)
	}
	if m.Matches(NewElement(Tag("a"))) {
		err(t, "incorrect field matcher matches", true, false)
	}
}

func TestClassMatcher(t *testing.T) {
	c := Class("fine")
	a1 := NewElement(Tag("a"))
	a1.Add(c)
	if !c.Matches(a1) {
		err(t, "incorrect class matcher matches", false, true)
	}
	a2 := NewElement(Tag("a"))
	if c.Matches(a2) {
		err(t, "incorrect class matcher matches", true, false)
	}
}

func TestNotMatcher(t *testing.T) {
	cl := Class("fine")
	c := Not(cl)
	a1 := NewElement(Tag("a"))
	a1.Add(cl)
	if c.Matches(a1) {
		err(t, "incorrect not matcher matches", true, false)
	}
	a2 := NewElement(Tag("a"))
	if !c.Matches(a2) {
		err(t, "incorrect not matcher matches", false, true)
	}
}

func TestIdMatcher(t *testing.T) {
	i := Id("fine")
	a1 := NewElement(Tag("a"))
	a1.Add(i)
	if !i.Matches(a1) {
		err(t, "incorrect id matcher matches", false, true)
	}
	a2 := NewElement(Tag("a"))
	if i.Matches(a2) {
		err(t, "incorrect id matcher matches", true, false)
	}
}

func TestOrMatcher(t *testing.T) {
	i := Or(Id("fine"), Class("well"))
	a1 := NewElement(Tag("a"))
	a1.Add(Id("fine"))

	if !i.Matches(a1) {
		err(t, "incorrect or id matcher matches", false, true)
	}

	a2 := NewElement(Tag("a"))
	a2.Add(Class("well"))
	if !i.Matches(a2) {
		err(t, "incorrect or class matcher matches", false, true)
	}
	a3 := NewElement(Tag("a"))
	if i.Matches(a3) {
		err(t, "incorrect or matcher matches", true, false)
	}
}

func TestAndMatcher(t *testing.T) {
	i := And(Id("fine"), Class("well"))
	a1 := NewElement(Tag("a"))
	a1.Add(Id("fine"), Class("well"))

	if !i.Matches(a1) {
		err(t, "incorrect and matcher matches", false, true)
	}
	a2 := NewElement(Tag("a"))
	a2.Add(Class("well"))

	if i.Matches(a2) {
		err(t, "incorrect and  class matcher matches", true, false)
	}
	a3 := NewElement(Tag("a"))
	a3.Add(Id("fine"))

	if i.Matches(a3) {
		err(t, "incorrect and id matcher matches", true, false)
	}
}

func TestTagMatcher(t *testing.T) {
	m := Tag("a")
	if !m.Matches(NewElement(Tag("a"))) {
		err(t, "incorrect tag matcher matches", false, true)
	}
	if m.Matches(NewElement(Tag("b"))) {
		err(t, "incorrect tag matcher matches", true, false)
	}
}

func TestAttrMatcher(t *testing.T) {
	m := SingleAttr{"width", "200"}
	div1 := NewElement(Tag("div"))
	div1.Add(Attr("width", "200"))
	if !m.Matches(div1) {
		err(t, "incorrect Attr matcher matches", false, true)
	}
	div2 := NewElement(Tag("div"))
	if m.Matches(div2) {
		err(t, "incorrect Attr matcher matches", true, false)
	}
}

func TestAttrsMatcher(t *testing.T) {
	m := Attr("width", "200", "height", "300")
	div1 := NewElement(Tag("div"))
	div1.Add(Attr("width", "200", "height", "300"))
	if !m.Matches(div1) {
		err(t, "incorrect Attrs matcher matches", false, true)
	}
	div2 := NewElement(Tag("div"))
	div2.Add(Attr("width", "200"))
	if m.Matches(div2) {
		err(t, "incorrect Attrs matcher matches", true, false)
	}
}

/*
func TestStyleMatcher(t *testing.T) {
	m := Style("width", "200")
	if !m.Matches(Div(Style("width", "200"))) {
		err(t, "incorrect Style matcher matches", false, true)
	}
	if m.Matches(Div()) {
		err(t, "incorrect Style matcher matches", true, false)
	}
}

func TestStylesMatcher(t *testing.T) {
	m := Style("width", "200", "height", "300")
	if !m.Matches(Div(Style("width", "200", "height", "300"))) {
		err(t, "incorrect Styles matcher matches", false, true)
	}
	if m.Matches(Div(Style("width", "200"))) {
		err(t, "incorrect Styles matcher matches", true, false)
	}
}
*/
func TestHtmlMatcher(t *testing.T) {
	a := NewElement(Tag("a"))
	a.Add(Text("hiho"))
	m := Html(a.String())
	div := NewElement(Tag("div"))
	div.Add(a)
	if !m.Matches(div) {
		err(t, "incorrect Html matcher matches", false, true)
	}

	a2 := NewElement(Tag("a"))
	a2.Add(Text("hoho"))
	div2 := NewElement(Tag("div"))
	div2.Add(a2)
	if m.Matches(div2) {
		err(t, "incorrect Html matcher matches", true, false)
	}
}
