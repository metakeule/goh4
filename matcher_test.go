package goh4

import (
	"testing"
)

func TestPositionMatcher(t *testing.T) {
	e := P()
	m := &PositionMatcher{Element: e}

	if m.Matches(A()) || m.Found {
		err(t, "incorrect position matcher matches", true, false)
	}

	_ = m.Matches(A())

	if !m.Matches(e) || !m.Found {
		err(t, "incorrect position matcher matches", false, true)
	}

	_ = m.Matches(A())

	if m.Pos != 2 {
		err(t, "incorrect position matcher position first round", m.Pos, 2)
	}
}

func TestFieldMatcher(t *testing.T) {
	var m FieldMatcher
	if !m.Matches(Input()) {
		err(t, "incorrect field matcher matches", false, true)
	}
	if m.Matches(A()) {
		err(t, "incorrect field matcher matches", true, false)
	}
}

func TestClassMatcher(t *testing.T) {
	c := Class("fine")
	if !c.Matches(A(Class("fine"))) {
		err(t, "incorrect class matcher matches", false, true)
	}
	if c.Matches(A()) {
		err(t, "incorrect class matcher matches", true, false)
	}
}

func TestNotMatcher(t *testing.T) {
	c := Not(Class("fine"))
	if c.Matches(A(Class("fine"))) {
		err(t, "incorrect not matcher matches", true, false)
	}
	if !c.Matches(A()) {
		err(t, "incorrect not matcher matches", false, true)
	}
}

func TestIdMatcher(t *testing.T) {
	i := Id("fine")
	if !i.Matches(A(Id("fine"))) {
		err(t, "incorrect id matcher matches", false, true)
	}
	if i.Matches(A()) {
		err(t, "incorrect id matcher matches", true, false)
	}
}

func TestOrMatcher(t *testing.T) {
	i := Or(Id("fine"), Class("well"))
	if !i.Matches(A(Id("fine"))) {
		err(t, "incorrect or id matcher matches", false, true)
	}
	if !i.Matches(A(Class("well"))) {
		err(t, "incorrect or class matcher matches", false, true)
	}
	if i.Matches(A()) {
		err(t, "incorrect or matcher matches", true, false)
	}
}

func TestAndMatcher(t *testing.T) {
	i := And(Id("fine"), Class("well"))
	if !i.Matches(A(Id("fine"), Class("well"))) {
		err(t, "incorrect and matcher matches", false, true)
	}
	if i.Matches(A(Class("well"))) {
		err(t, "incorrect and  class matcher matches", true, false)
	}
	if i.Matches(A(Id("fine"))) {
		err(t, "incorrect and id matcher matches", true, false)
	}
}

func TestTagMatcher(t *testing.T) {
	m := Tag("a")
	if !m.Matches(A()) {
		err(t, "incorrect tag matcher matches", false, true)
	}
	if m.Matches(B()) {
		err(t, "incorrect tag matcher matches", true, false)
	}
}

func TestAttrMatcher(t *testing.T) {
	m := attr{"width", "200"}
	if !m.Matches(Div(Attr("width", "200"))) {
		err(t, "incorrect Attr matcher matches", false, true)
	}
	if m.Matches(Div()) {
		err(t, "incorrect Attr matcher matches", true, false)
	}
}

func TestAttrsMatcher(t *testing.T) {
	m := Attr("width", "200", "height", "300")
	if !m.Matches(Div(Attr("width", "200", "height", "300"))) {
		err(t, "incorrect Attrs matcher matches", false, true)
	}
	if m.Matches(Div(Attr("width", "200"))) {
		err(t, "incorrect Attrs matcher matches", true, false)
	}
}

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

func TestHtmlMatcher(t *testing.T) {
	m := Html(A(Text("hiho")).String())
	if !m.Matches(Div(A(Text("hiho")))) {
		err(t, "incorrect Html matcher matches", false, true)
	}
	if m.Matches(Div(A(Text("hoho")))) {
		err(t, "incorrect Html matcher matches", true, false)
	}
}
