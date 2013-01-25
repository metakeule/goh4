package goh4

import "fmt"

func A(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("a"), Inline)
	t.Set(objects...)
	return
}

func Abbr(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("abbr"))
	t.Set(objects...)
	return
}

func Acronym(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("acronym"))
	t.Set(objects...)
	return
}

func Address(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("address"))
	t.Set(objects...)
	return
}

func Area(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("area"), SelfClosing)
	t.Set(objects...)
	return
}

func Article(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("article"))
	t.Set(objects...)
	return
}

func Aside(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("aside"))
	t.Set(objects...)
	return
}

func Audio(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("audio"))
	t.Set(objects...)
	return
}

func B(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("b"), Inline)
	t.Set(objects...)
	return
}

func Base(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("base"), SelfClosing)
	t.Set(objects...)
	return
}

func Bdo(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("bdo"))
	t.Set(objects...)
	return
}

func Bdi(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("bdi"), Inline)
	t.Set(objects...)
	return
}

func Big(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("big"))
	t.Set(objects...)
	return
}

func Blockquote(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("blockquote"))
	t.Set(objects...)
	return
}

func Body(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("body"))
	t.parentTags = []Tag{Tag("doc")}
	t.Set(objects...)
	return
}

func Br(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("br"), SelfClosing)
	t.Set(objects...)
	return
}

// FormField ??
func Button(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("button"), FormField)
	t.Set(objects...)
	return
}

func Canvas(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("canvas"))
	t.Set(objects...)
	return
}

func Caption(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("caption"))
	t.Set(objects...)
	return
}

func Cite(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("cite"))
	t.Set(objects...)
	return
}

func Code(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("code"))
	t.Set(objects...)
	return
}

func Col(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("col"), SelfClosing)
	t.Set(objects...)
	return
}

func Colgroup(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("colgroup"))
	t.Set(objects...)
	return
}

func Command(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("command"), SelfClosing|Inline)
	t.Set(objects...)
	return
}

func Data(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("data"), Inline)
	t.Set(objects...)
	return
}

func Datalist(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("datalist"), Inline)
	t.Set(objects...)
	return
}

func Dd(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("dd"))
	t.Set(objects...)
	return
}

func Del(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("del"))
	t.Set(objects...)
	return
}

func Details(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("details"))
	t.Set(objects...)
	return
}

func Dfn(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("dfn"))
	t.Set(objects...)
	return
}

func Div(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("div"))
	t.Set(objects...)
	return
}

func Dialog(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("dialog"))
	t.Set(objects...)
	return
}

func Dl(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("dl"))
	t.Set(objects...)
	return
}

func Dt(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("dt"))
	t.Set(objects...)
	return
}

func Em(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("em"), Inline)
	t.Set(objects...)
	return
}

func Embed(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("embed"), SelfClosing)
	t.Set(objects...)
	return
}

func Fieldset(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("fieldset"))
	t.Set(objects...)
	return
}

func Figcaption(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("figcaption"), Inline)
	t.Set(objects...)
	return
}

func Figure(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("figure"))
	t.Set(objects...)
	return
}

func Footer(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("footer"))
	t.Set(objects...)
	return
}

func Form(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("form"))
	t.Set(objects...)
	return
}

func Frame(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("frame"), SelfClosing)
	t.Set(objects...)
	return
}

func Frameset(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("frameset"))
	t.Set(objects...)
	return
}

// make a heading of level n (h1, h2, ...)
func H(n int, objects ...Stringer) (t *Element) {
	t = NewElement(Tag(fmt.Sprintf("h%v", n)))
	t.Set(objects...)
	return
}

func H1(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("h1"))
	t.Set(objects...)
	return
}

func H2(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("h2"))
	t.Set(objects...)
	return
}

func H3(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("h3"))
	t.Set(objects...)
	return
}

func H4(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("h4"))
	t.Set(objects...)
	return
}

func H5(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("h5"))
	t.Set(objects...)
	return
}

func H6(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("h6"))
	t.Set(objects...)
	return
}

func Head(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("head"), IdForbidden|ClassForbidden|Invisible)
	t.parentTags = []Tag{Tag("doc")}
	t.Set(objects...)
	return
}

func Header(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("header"))
	t.Set(objects...)
	return
}

func Hgroup(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("hgroup"))
	t.Set(objects...)
	return
}

func Hr(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("hr"), SelfClosing)
	t.Set(objects...)
	return
}

/*
func Html(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("html"))
	t.Set(objects...)
	return
}
*/

func I(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("i"), Inline)
	t.Set(objects...)
	return
}

func Iframe(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("iframe"), SelfClosing, Invisible)
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

func Ins(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("ins"))
	t.Set(objects...)
	return
}

func Kbd(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("kbd"))
	t.Set(objects...)
	return
}

func Keygen(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("keygen"), SelfClosing|Inline|FormField)
	t.Set(objects...)
	return
}

func Label(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("label"))
	t.Set(objects...)
	return
}

func Legend(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("legend"))
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

func Map(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("map"))
	t.Set(objects...)
	return
}

func Mark(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("mark"), Inline)
	t.Set(objects...)
	return
}

func Menu(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("menu"))
	t.Set(objects...)
	return
}

func Meta(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("meta"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.parentTags = []Tag{Tag("head")}
	t.Set(objects...)
	return
}

func Meter(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("meter"), Inline)
	t.Set(objects...)
	return
}

func Nav(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("nav"))
	t.Set(objects...)
	return
}

func Noframes(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("noframes"))
	t.Set(objects...)
	return
}

func Noscript(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("noscript"))
	t.Set(objects...)
	return
}

func Object(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("object"), FormField)
	t.Set(objects...)
	return
}

func Ol(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("ol"))
	t.Set(objects...)
	return
}

func Optgroup(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("optgroup"))
	t.Set(objects...)
	return
}

func Option(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("option"))
	t.parentTags = []Tag{Tag("select")}
	t.Set(objects...)
	return
}

func Output(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("output"), Inline)
	t.Set(objects...)
	return
}

func P(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("p"))
	t.Set(objects...)
	return
}

func Param(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("param"), SelfClosing)
	t.Set(objects...)
	return
}

// TODO: check if it is really not escaped
func Pre(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("pre"), WithoutEscaping)
	t.Set(objects...)
	return
}

func Progress(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("progress"), Inline)
	t.Set(objects...)
	return
}

func Q(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("q"))
	t.Set(objects...)
	return
}

func Rp(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("rp"), Inline)
	t.Set(objects...)
	return
}

func Rt(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("rt"), Inline)
	t.Set(objects...)
	return
}

func Ruby(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("ruby"), Inline)
	t.Set(objects...)
	return
}

func Samp(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("samp"))
	t.Set(objects...)
	return
}

func Script(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("script"), Invisible)
	t.parentTags = []Tag{Tag("head"), Tag("body")}
	t.Set(objects...)
	return
}

func Section(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("section"))
	t.Set(objects...)
	return
}

func Select(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("select"), FormField, Inline)
	t.Set(objects...)
	return
}

func Small(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("small"))
	t.Set(objects...)
	return
}

func Source(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("source"))
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

/*
func Style(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("style"))
	t.Set(objects...)
	return
}
*/

func Sub(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("sub"))
	t.Set(objects...)
	return
}

func Summary(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("summary"), Inline)
	t.Set(objects...)
	return
}

func Sup(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("sup"))
	t.Set(objects...)
	return
}

func Svg(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("svg"))
	t.Set(objects...)
	return
}

func Table(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("table"))
	t.Set(objects...)
	return
}

func Tbody(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("tbody"))
	t.parentTags = []Tag{Tag("table")}
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

func Tfoot(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("tfoot"))
	t.Set(objects...)
	return
}

func Th(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("th"))
	t.parentTags = []Tag{Tag("tr")}
	t.Set(objects...)
	return
}

func Thead(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("thead"))
	t.parentTags = []Tag{Tag("table")}
	t.Set(objects...)
	return
}

func Time(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("time"), Inline)
	t.Set(objects...)
	return
}

func Title(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("title"), Invisible)
	t.parentTags = []Tag{Tag("head")}
	t.Set(objects...)
	return
}

func Tr(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("tr"))
	t.parentTags = []Tag{Tag("tbody"), Tag("table"), Tag("thead")}
	t.Set(objects...)
	return
}

func Track(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("track"), SelfClosing)
	t.parentTags = []Tag{Tag("audio"), Tag("video")}
	t.Set(objects...)
	return
}

func Tt(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("tt"))
	t.Set(objects...)
	return
}

func Ul(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("ul"))
	t.Set(objects...)
	return
}

func Var(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("var"))
	t.Set(objects...)
	return
}

func Video(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("video"))
	t.Set(objects...)
	return
}

func Wbr(objects ...Stringer) (t *Element) {
	t = NewElement(Tag("wbr"), SelfClosing, Inline)
	t.Set(objects...)
	return
}
