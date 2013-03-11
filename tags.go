package goh4

import "fmt"

func A(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("a"), Inline)
	t.Add(objects...)
	return
}

func Abbr(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("abbr"))
	t.Add(objects...)
	return
}

func Acronym(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("acronym"))
	t.Add(objects...)
	return
}

func Address(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("address"))
	t.Add(objects...)
	return
}

func Area(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("area"), SelfClosing)
	t.Add(objects...)
	return
}

func Article(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("article"))
	t.Add(objects...)
	return
}

func Aside(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("aside"))
	t.Add(objects...)
	return
}

func Audio(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("audio"))
	t.Add(objects...)
	return
}

func B(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("b"), Inline)
	t.Add(objects...)
	return
}

func Base(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("base"), SelfClosing)
	t.Add(objects...)
	return
}

func Bdo(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("bdo"))
	t.Add(objects...)
	return
}

func Bdi(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("bdi"), Inline)
	t.Add(objects...)
	return
}

func Big(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("big"))
	t.Add(objects...)
	return
}

func Blockquote(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("blockquote"))
	t.Add(objects...)
	return
}

func Body(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("body"))
	t.parentTags = []Tag{Tag("doc")}
	t.Add(objects...)
	return
}

func Br(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("br"), SelfClosing)
	t.Add(objects...)
	return
}

// FormField ??
func Button(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("button"), FormField)
	t.Add(objects...)
	return
}

func Canvas(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("canvas"))
	t.Add(objects...)
	return
}

func Caption(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("caption"))
	t.Add(objects...)
	return
}

func Cite(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("cite"))
	t.Add(objects...)
	return
}

func Code(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("code"))
	t.Add(objects...)
	return
}

func Col(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("col"), SelfClosing)
	t.Add(objects...)
	return
}

func Colgroup(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("colgroup"))
	t.Add(objects...)
	return
}

func Command(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("command"), SelfClosing|Inline)
	t.Add(objects...)
	return
}

func Data(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("data"), Inline)
	t.Add(objects...)
	return
}

func Datalist(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("datalist"), Inline)
	t.Add(objects...)
	return
}

func Dd(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dd"))
	t.Add(objects...)
	return
}

func Del(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("del"))
	t.Add(objects...)
	return
}

func Details(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("details"))
	t.Add(objects...)
	return
}

func Dfn(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dfn"))
	t.Add(objects...)
	return
}

func Div(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("div"))
	t.Add(objects...)
	return
}

func Dialog(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dialog"))
	t.Add(objects...)
	return
}

func Dl(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dl"))
	t.Add(objects...)
	return
}

func Dt(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dt"))
	t.Add(objects...)
	return
}

func Em(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("em"), Inline)
	t.Add(objects...)
	return
}

func Embed(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("embed"), SelfClosing)
	t.Add(objects...)
	return
}

func Fieldset(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("fieldset"))
	t.Add(objects...)
	return
}

func Figcaption(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("figcaption"), Inline)
	t.Add(objects...)
	return
}

func Figure(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("figure"))
	t.Add(objects...)
	return
}

func Footer(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("footer"))
	t.Add(objects...)
	return
}

func Form(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("form"))
	t.Add(Attr("enctype", "multipart/form-data", "method", "post"))
	t.Add(objects...)
	return
}

func Frame(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("frame"), SelfClosing)
	t.Add(objects...)
	return
}

func Frameset(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("frameset"))
	t.Add(objects...)
	return
}

// make a heading of level n (h1, h2, ...)
func H(n int, objects ...interface{}) (t *Element) {
	t = NewElement(Tag(fmt.Sprintf("h%v", n)))
	t.Add(objects...)
	return
}

func H1(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("h1"))
	t.Add(objects...)
	return
}

func H2(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("h2"))
	t.Add(objects...)
	return
}

func H3(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("h3"))
	t.Add(objects...)
	return
}

func H4(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("h4"))
	t.Add(objects...)
	return
}

func H5(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("h5"))
	t.Add(objects...)
	return
}

func H6(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("h6"))
	t.Add(objects...)
	return
}

func Head(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("head"), IdForbidden|ClassForbidden|Invisible)
	t.parentTags = []Tag{Tag("doc")}
	t.Add(objects...)
	return
}

func Header(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("header"))
	t.Add(objects...)
	return
}

func Hgroup(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("hgroup"))
	t.Add(objects...)
	return
}

func Hr(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("hr"), SelfClosing)
	t.Add(objects...)
	return
}

/*
func Html(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("html"))
	t.Add(objects...)
	return
}
*/

func I(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("i"), Inline)
	t.Add(objects...)
	return
}

func Iframe(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("iframe"), SelfClosing, Invisible)
	t.Add(objects...)
	return
}

func Img(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("img"), Inline, SelfClosing)
	t.Add(objects...)
	return
}

func Input(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("input"), FormField, SelfClosing, Inline)
	t.Add(objects...)
	return
}

func Ins(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ins"))
	t.Add(objects...)
	return
}

func Kbd(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("kbd"))
	t.Add(objects...)
	return
}

func Keygen(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("keygen"), SelfClosing|Inline|FormField)
	t.Add(objects...)
	return
}

func Label(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("label"))
	t.Add(objects...)
	return
}

func Legend(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("legend"))
	t.Add(objects...)
	return
}

func Li(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("li"))
	t.parentTags = []Tag{Tag("ul"), Tag("ol")}
	t.Add(objects...)
	return
}

func Link(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("link"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.parentTags = []Tag{Tag("head"), Tag("body")}
	t.Add(objects...)
	return
}

func Map(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("map"))
	t.Add(objects...)
	return
}

func Mark(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("mark"), Inline)
	t.Add(objects...)
	return
}

func Menu(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("menu"))
	t.Add(objects...)
	return
}

func Meta(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("meta"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.parentTags = []Tag{Tag("head")}
	t.Add(objects...)
	return
}

func Meter(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("meter"), Inline)
	t.Add(objects...)
	return
}

func Nav(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("nav"))
	t.Add(objects...)
	return
}

func Noframes(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("noframes"))
	t.Add(objects...)
	return
}

func Noscript(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("noscript"))
	t.Add(objects...)
	return
}

func Object(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("object"), FormField)
	t.Add(objects...)
	return
}

func Ol(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ol"))
	t.Add(objects...)
	return
}

func Optgroup(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("optgroup"))
	t.Add(objects...)
	return
}

func Option(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("option"))
	t.parentTags = []Tag{Tag("select")}
	t.Add(objects...)
	return
}

func Output(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("output"), Inline)
	t.Add(objects...)
	return
}

func P(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("p"))
	t.Add(objects...)
	return
}

func Param(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("param"), SelfClosing)
	t.Add(objects...)
	return
}

// TODO: check if it is really not escaped
func Pre(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("pre"), WithoutEscaping)
	t.Add(objects...)
	return
}

func Progress(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("progress"), Inline)
	t.Add(objects...)
	return
}

func Q(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("q"))
	t.Add(objects...)
	return
}

func Rp(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("rp"), Inline)
	t.Add(objects...)
	return
}

func Rt(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("rt"), Inline)
	t.Add(objects...)
	return
}

func Ruby(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ruby"), Inline)
	t.Add(objects...)
	return
}

func Samp(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("samp"))
	t.Add(objects...)
	return
}

func Script(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("script"), Invisible)
	t.parentTags = []Tag{Tag("head"), Tag("body")}
	t.Add(objects...)
	return
}

func Section(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("section"))
	t.Add(objects...)
	return
}

func Select(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("select"), FormField, Inline)
	t.Add(objects...)
	return
}

func Small(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("small"))
	t.Add(objects...)
	return
}

func Source(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("source"))
	t.Add(objects...)
	return
}

func Span(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("span"), Inline)
	t.Add(objects...)
	return
}

func Strong(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("strong"), Inline)
	t.Add(objects...)
	return
}

/*
func Style(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("style"))
	t.Add(objects...)
	return
}
*/

func Sub(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("sub"))
	t.Add(objects...)
	return
}

func Summary(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("summary"), Inline)
	t.Add(objects...)
	return
}

func Sup(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("sup"))
	t.Add(objects...)
	return
}

func Svg(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("svg"))
	t.Add(objects...)
	return
}

func Table(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("table"))
	t.Add(objects...)
	return
}

func Tbody(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tbody"))
	t.parentTags = []Tag{Tag("table")}
	t.Add(objects...)
	return
}

func Td(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("td"))
	t.parentTags = []Tag{Tag("tr")}
	t.Add(objects...)
	return
}

func Textarea(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("textarea"), FormField)
	t.Add(objects...)
	return
}

func Tfoot(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tfoot"))
	t.Add(objects...)
	return
}

func Th(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("th"))
	t.parentTags = []Tag{Tag("tr")}
	t.Add(objects...)
	return
}

func Thead(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("thead"))
	t.parentTags = []Tag{Tag("table")}
	t.Add(objects...)
	return
}

func Time(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("time"), Inline)
	t.Add(objects...)
	return
}

func Title(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("title"), Invisible)
	t.parentTags = []Tag{Tag("head")}
	t.Add(objects...)
	return
}

func Tr(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tr"))
	t.parentTags = []Tag{Tag("tbody"), Tag("table"), Tag("thead")}
	t.Add(objects...)
	return
}

func Track(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("track"), SelfClosing)
	t.parentTags = []Tag{Tag("audio"), Tag("video")}
	t.Add(objects...)
	return
}

func Tt(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tt"))
	t.Add(objects...)
	return
}

func Ul(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ul"))
	t.Add(objects...)
	return
}

func Var(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("var"))
	t.Add(objects...)
	return
}

func Video(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("video"))
	t.Add(objects...)
	return
}

func Wbr(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("wbr"), SelfClosing, Inline)
	t.Add(objects...)
	return
}
