package goh4

import "fmt"

func A(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("a"), Inline)
	t.Set(objects...)
	return
}
func Body(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("body"))
	t.parentTags = []Tag{Tag("doc")}
	t.Set(objects...)
	return
}
func Button(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("button"))
	t.Set(objects...)
	return
}
func B(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("b"), Inline)
	t.Set(objects...)
	return
}
func Div(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("div"))
	t.Set(objects...)
	return
}
func Footer(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("footer"))
	t.Set(objects...)
	return
}
func P(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("p"))
	t.Set(objects...)
	return
}
func Form(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("form"), Invisible)
	t.Set(objects...)
	return
}
func Head(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("head"), IdForbidden|ClassForbidden|Invisible)
	t.parentTags = []Tag{Tag("doc")}
	t.Set(objects...)
	return
}

// make a heading of level n (h1, h2, ...)
func H(n int, objects ...Stringer) (t *Element) {
	t = NewElement(Tag(fmt.Sprintf("h%v", n)))
	t.Set(objects...)
	return
}
func Doc(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("doc"), WithoutDecoration)
	t.Set(objects...)
	return
}
func Hr(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("hr"), SelfClosing)
	t.Set(objects...)
	return
}
func Iframe(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("iframe"), SelfClosing, Invisible)
	t.Set(objects...)
	return
}
func I(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("i"), Inline)
	t.Set(objects...)
	return
}
func Img(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("img"), Inline, SelfClosing)
	t.Set(objects...)
	return
}
func Input(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("input"), FormField, SelfClosing, Inline)
	t.Set(objects...)
	return
}
func Label(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("label"))
	t.Set(objects...)
	return
}
func Li(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("li"))
	t.parentTags = []Tag{Tag("ul"), Tag("ol")}
	t.Set(objects...)
	return
}
func Link(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("link"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.parentTags = []Tag{Tag("head"), Tag("body")}
	t.Set(objects...)
	return
}
func Meta(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("meta"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.parentTags = []Tag{Tag("head")}
	t.Set(objects...)
	return
}
func Option(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("option"))
	t.parentTags = []Tag{Tag("select")}
	t.Set(objects...)
	return
}

// TODO: check if it is really not escaped
func Pre(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("pre"), WithoutEscaping)
	t.Set(objects...)
	return
}
func Script(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("script"), Invisible)
	t.parentTags = []Tag{Tag("head"), Tag("body")}
	t.Set(objects...)
	return
}
func Select(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("select"), FormField, Inline)
	t.Set(objects...)
	return
}
func Span(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("span"), Inline)
	t.Set(objects...)
	return
}
func Strong(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("strong"), Inline)
	t.Set(objects...)
	return
}
func Table(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("table"))
	t.Set(objects...)
	return
}
func Tr(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("tr"))
	t.parentTags = []Tag{Tag("tbody"), Tag("table"), Tag("thead")}
	t.Set(objects...)
	return
}
func Th(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("th"))
	t.parentTags = []Tag{Tag("tr")}
	t.Set(objects...)
	return
}
func Td(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("td"))
	t.parentTags = []Tag{Tag("tr")}
	t.Set(objects...)
	return
}
func Textarea(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("textarea"), FormField)
	t.Set(objects...)
	return
}
func Title(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("title"), Invisible)
	t.parentTags = []Tag{Tag("head")}
	t.Set(objects...)
	return
}
func Ul(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("ul"))
	t.Set(objects...)
	return
}
func Ol(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("ol"))
	t.Set(objects...)
	return
}
