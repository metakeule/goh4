package goh4

import (
	"testing"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %s, should be %s\n", is, shouldbe)
}

func TestNewElement(t *testing.T) {
	a := NewElement(Tag("a"))

	if a.invisible {
		err(t, "incorrect visibility", a.invisible, false)
	}

	if a.withoutDecoration {
		err(t, "incorrect decoration", a.withoutDecoration, false)
	}

	if a.withoutEscaping {
		err(t, "incorrect escaping", a.withoutEscaping, false)
	}

	if a.selfclosing {
		err(t, "incorrect selfclosing", a.selfclosing, false)
	}

	if a.idForbidden {
		err(t, "incorrect idForbidden", a.idForbidden, false)
	}

	if a.classForbidden {
		err(t, "incorrect classForbidden", a.classForbidden, false)
	}

	if a.Tag() != "a" {
		err(t, "incorrect tag", a.Tag(), "a")
	}

}

func TestElementAdd(t *testing.T) {
	span := Span()
	strong := Strong()
	css := NewCss(Class("good-looking"))

	a := NewElement(Tag("a"))
	a.Set(
		Id("myid"),
		Class("fine"),
		Class("fancy"),
		Text("<escaped>"),
		Html("<unchanged!>"),
		Comment("no comment"),
		span,
		span,
		strong,
		span,
		css,
		Styles{"color", "green", "zoom", "1"},
		Attrs{"height", "300", "width", "200"})

	if a.Attributes["height"] != "300" {
		err(t, "incorrect height", a.Attributes["height"], "300")
	}

	if a.Attributes["width"] != "200" {
		err(t, "incorrect width", a.Attributes["width"], "200")
	}

	if a.Id() != "myid" {
		err(t, "incorrect id", a.Id(), "myid")
	}

	classes := a.Classes()
	if classes[0] != "fine" {
		err(t, "incorrect class", classes[0], "fine")
	}
	if classes[1] != "fancy" {
		err(t, "incorrect class", classes[1], "fancy")
	}

	if classes[2] != "good-looking" {
		err(t, "incorrect class from css assignment", classes[2], "good-looking")
	}

	if a.style["color"] != "green" {
		err(t, "incorrect color", a.style["color"], "green")
	}

	if a.style["zoom"] != "1" {
		err(t, "incorrect zoom", a.style["zoom"], "1")
	}

	if span.Parent() != a {
		err(t, "incorrect parent", span.Parent(), a)
	}

	if a.Comment != "no comment" {
		err(t, "incorrect comment", a.Comment, "no comment")
	}

	if a.inner[0] != Text("&lt;escaped&gt;") {
		err(t, "incorrect inner content", a.inner[0], "&lt;escaped&gt;")
	}

	if a.inner[1] != Html("<unchanged!>") {
		err(t, "incorrect inner content", a.inner[1], "<unchanged!>")
	}

	if a.inner[2] != span {
		err(t, "incorrect inner content", a.inner[2], span)
	}

	if span.Path() != "a#myid.fine.fancy.good-looking span" {
		err(t, "incorrect path", span.Path(), "a#myid.fine.fancy.good-looking span")
	}

	spans := a.All(Tag("span"))

	if len(spans) != 3 {
		err(t, "incorrect number of spans", len(spans), 3)
	}

	if spans[0] != span {
		err(t, "incorrect result for all spans", spans[0], span)
	}

	_, strng := a.Any(Tag("strong"))

	if strong != strng {
		err(t, "incorrect result for any strong", strng, strong)
	}

}

// TODO add a method to set content before or after a certain inner tag or at a special position, e.g.
// AddTop(), AddBottom(), AddBefore(Stringer), AddAfter(Stringer), AddAtPosition(pos int)
// SetTop(), SetBottom(), SetAtPosition(pos int)
// add convenience method SetAttr(attr ...string), that always takes a pair of two strings and
// handles them as key and value
// wir brauchen auch einen platzhalter, der ein "pseudo-tag" darstellt, der nur dazu da ist, eine stelle zu
// markieren, die durch etwas anderes ersetzt wird, aber nicht als tag dargestellt wird

func TestAddTop(t *testing.T) {
	a := A()
	div := Div(a)
	b := B()
	div.AddAtPosition(0, b)

	if div.inner[0] != b {
		err(t, "incorrect result for AddTop", div.inner[0], b)
	}

	if div.inner[1] != a {
		err(t, "incorrect result for AddTop", div.inner[1], a)
	}
}

func TestAddBottom(t *testing.T) {
	a := A()
	div := Div(a)
	b := B()
	div.Add(b)

	if div.inner[1] != b {
		err(t, "incorrect result for AddBottom", div.inner[1], b)
	}

	if div.inner[0] != a {
		err(t, "incorrect result for AddBottom", div.inner[0], a)
	}
}

func TestAddAtPosition(t *testing.T) {
	a := A()
	b := B()
	h := Hr()

	div := Div(a, b)

	div.AddAtPosition(1, h)

	if div.inner[1] != h {
		err(t, "incorrect result for AddAtPosition", div.inner[1], h)
	}

	if div.inner[2] != b {
		err(t, "incorrect result for AddAtPosition", div.inner[2], b)
	}

	if div.inner[0] != a {
		err(t, "incorrect result for AddAtPosition", div.inner[0], a)
	}

	div2 := Div(a, b, b, a, b)

	div2.AddAtPosition(3, h)

	if div2.inner[3] != h {
		err(t, "incorrect result for AddAtPosition", div2.inner[3], h)
	}

	if div2.inner[2] != b {
		err(t, "incorrect result for AddAtPosition", div2.inner[2], b)
	}

	if div2.inner[4] != a {
		err(t, "incorrect result for AddAtPosition", div2.inner[4], a)
	}
}

func TestSetAtPosition(t *testing.T) {
	a := A()
	b := B()
	h := Hr()

	div := Div(a, b)

	div.SetAtPosition(1, h)

	if div.inner[1] != h {
		err(t, "incorrect result for SetAtPosition", div.inner[1], h)
	}

}

func TestSetBottom(t *testing.T) {
	a := A()
	b := B()
	h := Hr()

	div := Div(a, b)

	div.SetBottom(h)

	if div.inner[1] != h {
		err(t, "incorrect result for SetBottom", div.inner[1], h)
	}

}

func TestPositionOf(t *testing.T) {
	a := A()
	b := B()
	h := Hr()
	s := Strong()
	u := Ul()

	div := Div(a, b, h, s)

	pos, _ := div.PositionOf(b)

	if pos != 1 {
		err(t, "incorrect result for PositionOf", pos, 1)
	}

	pos, _ = div.PositionOf(s)
	if pos != 3 {
		err(t, "incorrect result for PositionOf", pos, 3)
	}

	pos, found := div.PositionOf(u)
	if pos != -1 || found {
		err(t, "incorrect result for PositionOf", pos, -1)
	}

}

func TestAddBefore(t *testing.T) {
	a := A()
	b := B()
	h := Hr()
	s := Strong()

	div := Div(a, b, h)

	div.AddBefore(b, s)

	pos, _ := div.PositionOf(b)

	if pos != 2 {
		err(t, "incorrect result for AddBefore", pos, 2)
	}

	pos, _ = div.PositionOf(s)
	if pos != 1 {
		err(t, "incorrect result for AddBefore", pos, 1)
	}
}

func TestAddAfter(t *testing.T) {
	a := A()
	b := B()
	h := Hr()
	s := Strong()
	sp := Span()

	div := Div(a, b, h)

	div.AddAfter(b, s)
	div.AddAfter(h, sp)

	pos, _ := div.PositionOf(b)

	if pos != 1 {
		err(t, "incorrect result for AddAfter", pos, 1)
	}

	pos, _ = div.PositionOf(s)
	if pos != 2 {
		err(t, "incorrect result for AddAfter", pos, 2)
	}

	pos, _ = div.PositionOf(sp)
	if pos != 4 {
		err(t, "incorrect result for AddAfter", pos, 4)
	}
}
