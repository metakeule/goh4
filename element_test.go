package goh4

import (
	"fmt"
	"testing"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %s, should be %s\n", is, shouldbe)
}

func TestNewElement(t *testing.T) {
	a := NewElement(Tag("a"))

	if a.Is(Invisible) {
		err(t, "incorrect visibility", a.Is(Invisible), false)
	}

	if a.Is(WithoutDecoration) {
		err(t, "incorrect decoration", a.Is(WithoutDecoration), false)
	}

	if a.Is(WithoutEscaping) {
		err(t, "incorrect escaping", a.Is(WithoutEscaping), false)
	}

	if a.Is(SelfClosing) {
		err(t, "incorrect selfclosing", a.Is(SelfClosing), false)
	}

	if a.Is(IdForbidden) {
		err(t, "incorrect idForbidden", a.Is(IdForbidden), false)
	}

	if a.Is(ClassForbidden) {
		err(t, "incorrect classForbidden", a.Is(ClassForbidden), false)
	}

	if a.Tag() != "a" {
		err(t, "incorrect tag", a.Tag(), "a")
	}

}

func TestElementRemoveAttribute(t *testing.T) {
	e := A(Attr("href", "#"))

	e.RemoveAttribute("href")

	if e.attributes["href"] != "" {
		err(t, "incorrect attribute removal", e.attributes["href"], "")
	}
}

// apply the css to the element
/*
func TestElementApplyCss(t *testing.T) {
	el := Div()
	el.ApplyCss(NewCss(Class("uuh")))

	if !el.HasClass(Class("uuh")) {
		err(t, "incorrect ApplyCss, has class", false, true)
	}

	if e := el.ApplyCss(NewCss()); e == nil {
		err(t, "incorrect ApplyCss, error no class", "no class in css rule", nil)
	}

	if e := el.ApplyCss(NewCss(Class("uuh"), Tag("ul"))); e == nil {
		err(t, "incorrect ApplyCss, error not affected", "element not affected by css rule", nil)
	}
}
*/

func TestElementAdd(t *testing.T) {
	span := Span()
	strong := Strong()
	//	css := NewCss(Class("good-looking"))

	a := NewElement(Tag("a"))
	a.Add(
		Id("myid"),
		Class("fine"),
		Class("fancy"),
		"<escaped>",
		Html("<unchanged!>"),
		Comment("no comment"),
		span,
		span,
		strong,
		span,
		//		css,
		//		Style("color", "green", "zoom", "1"),
		Attr("height", "300", "width", "200"))

	if a.attributes["height"] != "300" {
		err(t, "incorrect height", a.attributes["height"], "300")
	}

	if a.attributes["width"] != "200" {
		err(t, "incorrect width", a.attributes["width"], "200")
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

	/*
		if classes[2] != "good-looking" {
			err(t, "incorrect class from css assignment", classes[2], "good-looking")
		}
	*/
	/*
		if a.style["color"] != "green" {
			err(t, "incorrect color", a.style["color"], "green")
		}
	*/
	/*
		if a.style["zoom"] != "1" {
			err(t, "incorrect zoom", a.style["zoom"], "1")
		}
	*/
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
	/*
		if span.Path() != "a#myid.fine.fancy.good-looking span" {
			err(t, "incorrect path", span.Path(), "a#myid.fine.fancy.good-looking span")
		}
	*/
	spans := a.All(Tag("span"))

	if len(spans) != 3 {
		err(t, "incorrect number of spans", len(spans), 3)
	}

	if spans[0] != span {
		err(t, "incorrect result for all spans", spans[0], span)
	}

	strng := a.Any(Tag("strong"))

	if strong != strng {
		err(t, "incorrect result for any strong", strng, strong)
	}

}

type notAnElementer int

func (ø notAnElementer) String() string {
	return "but a stringer"
}

func TestElementAddErrors(t *testing.T) {
	elem := NewElement("paranoid", SelfClosing, ClassForbidden, IdForbidden)

	if e := elem.Add(Text("hiho")); e == nil {
		err(t, "incorrect result for ElementAdd with selfclosing", "content not allowed", nil)
	}

	if e := elem.Add(Id("hiho")); e == nil {
		err(t, "incorrect result for ElementAdd with IdForbidden", "id forbidden", nil)
	}

	if e := elem.Add(Class("hiho")); e == nil {
		err(t, "incorrect result for ElementAdd with ClassForbidden", "class forbidden", nil)
	}

	if e := elem.Add(notAnElementer(0)); e == nil {
		err(t, "incorrect result for ElementAdd with notAnElementer", "can't set me as parent", nil)
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

	if e := div.AddAtPosition(3, h); e == nil {
		err(t, "incorrect result for AddAtPosition error", "out of range", nil)
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

// create an element with a simple self defined tag
func ExampleNewElement() {
	t := NewElement(Tag("special"))
	fmt.Println(t)
	// Output: <special></special>
}

// a selfclosing tag can't have content
func ExampleNewElement_selfclosing() {
	t := NewElement(Tag("special"), SelfClosing)
	fmt.Println(t)
	if err := t.Add("will fail"); err != nil {
		fmt.Println("can't add to selfclosing element")
	}
	// Output: <special />
	// can't add to selfclosing element
}

// multiple flags may be passed with bitwise or | and as several parameters
func ExampleNewElement_multipleFlags() {
	t := NewElement(Tag("input"), SelfClosing|Inline, FormField)
	fmt.Println(t)
	// Output: <input />
}

// create an element that does not output it tags when printed
func ExampleNewElement_withoutDecoration() {
	doc := NewElement(Tag("doc"), WithoutDecoration)

	layout := NewTemplate(doc)

	doc.Add(
		Html("<!DOCTYPE html>"),
		Body(
			Id("content")),
		Html("</html>"))

	layout.Assign("content", "in the body")
	fmt.Println(layout)

	// Output: <!DOCTYPE html>
	// <body id="content">in the body</body>
	// </html>
}

// html, text and elementer are added as inner content
func ExampleElement_Add() {
	div := Div()

	div.Add(
		Html("<b>hello</b>"), // is not escaped
		Text("c > d"),        // is escaped
		Span("hiho"))         // objects to tag constructors like Div(), Span(),... gets passed to Add()

	fmt.Println(div)
	// Output: <div><b>hello</b>c &gt; d<span>hiho</span></div>
}

// add / set properties
func ExampleElement_Add_properties() {
	//	css := NewCss(Class("yellow"), Style("background-color", "yellow"))
	d := Div(Class("first")) // objects to tag constructors like Div(), Body(),... gets passed to Add()

	d.Add(
		Id("main"),
		Class("second"),
		//		css, // adds the class of the css to the element, multiple *Css can be given
		Comment("main row"),
		Attr("height", "200")) // multiple attributes at once with Attrs{"height", "200", "width", "300"}
	//Style("width", "500px")) // multiple styles at once with Styles{"height", "200", "width", "300"}

	//fmt.Printf("---CSS---%s---HTML---%s\n", css, d)
	// Output: ---CSS---
	// .yellow {
	// 	background-color: yellow;
	// }
	// ---HTML---
	// <!-- Begin: main row --><div id="main" class="first second yellow " style="width: 500px;" height="200"></div><!-- End: main row -->
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

	if e := div.SetAtPosition(2, h); e == nil {
		err(t, "incorrect result for SetAtPosition error", "out of range", nil)
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
	if found {
		err(t, "incorrect result for PositionOf found", found, false)
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

	if e := div.AddBefore(Body(), h); e == nil {
		err(t, "incorrect result for AddBefore error not found", "not found", nil)
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

	if e := div.AddAfter(Body(), h); e == nil {
		err(t, "incorrect result for AddAfter error not found", "not found", nil)
	}
}

func TestHasClass(t *testing.T) {
	e := A(Class("dipsy"), Class("flopsy"))
	if !e.HasClass(Class("dipsy")) {
		err(t, "incorrect result for HasClass having class", false, true)
	}

	if e.HasClass(Class("lopso")) {
		err(t, "incorrect result for HasClass having no class", true, false)
	}

	e.RemoveClass(Class("dipsy"))

	if e.HasClass(Class("dipsy")) {
		err(t, "incorrect result for RemoveClass", true, false)
	}

	if !e.HasClass(Class("flopsy")) {
		err(t, "incorrect result for RemoveClass keeping the others", false, true)
	}

	e.SetClass(Class("dipsy"), Class("lopso"))

	if e.HasClass(Class("flopsy")) || !e.HasClass(Class("dipsy")) || !e.HasClass(Class("lopso")) {
		err(t, "incorrect result for SetClass", false, true)
	}
}
