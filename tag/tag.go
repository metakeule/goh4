package tag

import (
	"fmt"
	. "gopkg.in/metakeule/goh4.v5"
)

func A(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("a"), Inline)
	t.Add(objects...)
	return
}

func ABBR(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("abbr"))
	t.Add(objects...)
	return
}

func ACRONYM(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("acronym"))
	t.Add(objects...)
	return
}

func ADDRESS(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("address"))
	t.Add(objects...)
	return
}

func AREA(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("area"), SelfClosing)
	t.Add(objects...)
	return
}

func ARTICLE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("article"))
	t.Add(objects...)
	return
}

func ASIDE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("aside"))
	t.Add(objects...)
	return
}

func AUDIO(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("audio"))
	t.Add(objects...)
	return
}

func B(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("b"), Inline)
	t.Add(objects...)
	return
}

func BASE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("base"), SelfClosing)
	t.Add(objects...)
	return
}

func BDO(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("bdo"))
	t.Add(objects...)
	return
}

func BDI(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("bdi"), Inline)
	t.Add(objects...)
	return
}

func BIG(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("big"))
	t.Add(objects...)
	return
}

func BLOCKQUOTE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("blockquote"))
	t.Add(objects...)
	return
}

func BODY(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("body"))
	t.ParentTags = []Tag{Tag("doc")}
	t.Add(objects...)
	return
}

func BR(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("br"), SelfClosing)
	t.Add(objects...)
	return
}

// FormField ??
func BUTTON(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("button"), FormField)
	t.Add(objects...)
	return
}

func CANVAS(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("canvas"))
	t.Add(objects...)
	return
}

func CAPTION(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("caption"))
	t.Add(objects...)
	return
}

func CITE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("cite"))
	t.Add(objects...)
	return
}

func CODE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("code"))
	t.Add(objects...)
	return
}

func COL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("col"), SelfClosing)
	t.Add(objects...)
	return
}

func COLGROUP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("colgroup"))
	t.Add(objects...)
	return
}

func COMMAND(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("command"), SelfClosing|Inline)
	t.Add(objects...)
	return
}

func DATA(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("data"), Inline)
	t.Add(objects...)
	return
}

func DATALIST(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("datalist"), Inline)
	t.Add(objects...)
	return
}

func DD(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dd"))
	t.Add(objects...)
	return
}

func DEL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("del"))
	t.Add(objects...)
	return
}

func DETAILS(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("details"))
	t.Add(objects...)
	return
}

func DFN(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dfn"))
	t.Add(objects...)
	return
}

func DIV(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("div"))
	t.Add(objects...)
	return
}

func DIALOG(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dialog"))
	t.Add(objects...)
	return
}

func DL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dl"))
	t.Add(objects...)
	return
}

func DT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("dt"))
	t.Add(objects...)
	return
}

func EM(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("em"), Inline)
	t.Add(objects...)
	return
}

func EMBED(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("embed"), SelfClosing)
	t.Add(objects...)
	return
}

func FIELDSET(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("fieldset"))
	t.Add(objects...)
	return
}

func FIGCAPTION(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("figcaption"), Inline)
	t.Add(objects...)
	return
}

func FIGURE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("figure"))
	t.Add(objects...)
	return
}

func FOOTER(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("footer"))
	t.Add(objects...)
	return
}

func FORM(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("form"))
	//t.Add(Attr("enctype", "multipart/form-data", "method", "post"))
	t.Add(objects...)
	return
}

func FRAME(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("frame"), SelfClosing)
	t.Add(objects...)
	return
}

func FRAMESET(objects ...interface{}) (t *Element) {
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

func HEAD(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("head"), IdForbidden|ClassForbidden|Invisible)
	t.ParentTags = []Tag{Tag("doc")}
	t.Add(objects...)
	return
}

func HEADER(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("header"))
	t.Add(objects...)
	return
}

func HGROUP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("hgroup"))
	t.Add(objects...)
	return
}

func HR(objects ...interface{}) (t *Element) {
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

func IFRAME(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("iframe"), SelfClosing, Invisible)
	t.Add(objects...)
	return
}

func IMG(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("img"), Inline, SelfClosing)
	t.Add(objects...)
	return
}

func INPUT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("input"), FormField, SelfClosing, Inline)
	t.Add(objects...)
	return
}

func INS(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ins"))
	t.Add(objects...)
	return
}

func KBD(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("kbd"))
	t.Add(objects...)
	return
}

func KEYGEN(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("keygen"), SelfClosing|Inline|FormField)
	t.Add(objects...)
	return
}

func LABEL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("label"))
	t.Add(objects...)
	return
}

func LEGEND(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("legend"))
	t.Add(objects...)
	return
}

func LI(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("li"))
	t.ParentTags = []Tag{Tag("ul"), Tag("ol")}
	t.Add(objects...)
	return
}

func LINK(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("link"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.ParentTags = []Tag{Tag("head"), Tag("body")}
	t.Add(objects...)
	return
}

func MAP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("map"))
	t.Add(objects...)
	return
}

func MARK(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("mark"), Inline)
	t.Add(objects...)
	return
}

func MENU(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("menu"))
	t.Add(objects...)
	return
}

func META(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("meta"), SelfClosing, IdForbidden, ClassForbidden, Invisible)
	t.ParentTags = []Tag{Tag("head")}
	t.Add(objects...)
	return
}

func METER(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("meter"), Inline)
	t.Add(objects...)
	return
}

func NAV(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("nav"))
	t.Add(objects...)
	return
}

func NOFRAMES(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("noframes"))
	t.Add(objects...)
	return
}

func NOSCRIPT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("noscript"))
	t.Add(objects...)
	return
}

func OBJECT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("object"), FormField)
	t.Add(objects...)
	return
}

func OL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ol"))
	t.Add(objects...)
	return
}

func OPTGROUP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("optgroup"))
	t.Add(objects...)
	return
}

func OPTION(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("option"))
	t.ParentTags = []Tag{Tag("select")}
	t.Add(objects...)
	return
}

func OUTPUT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("output"), Inline)
	t.Add(objects...)
	return
}

func P(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("p"))
	t.Add(objects...)
	return
}

func PARAM(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("param"), SelfClosing)
	t.Add(objects...)
	return
}

// TODO: check if it is really not escaped
func PRE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("pre"), WithoutEscaping)
	t.Add(objects...)
	return
}

func PROGRESS(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("progress"), Inline)
	t.Add(objects...)
	return
}

func Q(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("q"))
	t.Add(objects...)
	return
}

func RP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("rp"), Inline)
	t.Add(objects...)
	return
}

func RT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("rt"), Inline)
	t.Add(objects...)
	return
}

func RUBY(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ruby"), Inline)
	t.Add(objects...)
	return
}

func SAMP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("samp"))
	t.Add(objects...)
	return
}

func SCRIPT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("script"), Invisible)
	t.ParentTags = []Tag{Tag("head"), Tag("body")}
	t.Add(objects...)
	return
}

func SECTION(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("section"))
	t.Add(objects...)
	return
}

func SELECT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("select"), FormField, Inline)
	t.Add(objects...)
	return
}

func SMALL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("small"))
	t.Add(objects...)
	return
}

func SOURCE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("source"))
	t.Add(objects...)
	return
}

func SPAN(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("span"), Inline)
	t.Add(objects...)
	return
}

func STRONG(objects ...interface{}) (t *Element) {
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

func SUB(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("sub"))
	t.Add(objects...)
	return
}

func SUMMARY(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("summary"), Inline)
	t.Add(objects...)
	return
}

func SUP(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("sup"))
	t.Add(objects...)
	return
}

func SVG(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("svg"))
	t.Add(objects...)
	return
}

func TABLE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("table"))
	t.Add(objects...)
	return
}

func TBODY(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tbody"))
	t.ParentTags = []Tag{Tag("table")}
	t.Add(objects...)
	return
}

func TD(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("td"))
	t.ParentTags = []Tag{Tag("tr")}
	t.Add(objects...)
	return
}

func TEXTAREA(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("textarea"), FormField)
	t.Add(objects...)
	return
}

func TFOOT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tfoot"))
	t.Add(objects...)
	return
}

func TH(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("th"))
	t.ParentTags = []Tag{Tag("tr")}
	t.Add(objects...)
	return
}

func THEAD(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("thead"))
	t.ParentTags = []Tag{Tag("table")}
	t.Add(objects...)
	return
}

func TIME(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("time"), Inline)
	t.Add(objects...)
	return
}

func TITLE(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("title"), Invisible)
	t.ParentTags = []Tag{Tag("head")}
	t.Add(objects...)
	return
}

func TR(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tr"))
	t.ParentTags = []Tag{Tag("tbody"), Tag("table"), Tag("thead")}
	t.Add(objects...)
	return
}

func TRACK(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("track"), SelfClosing)
	t.ParentTags = []Tag{Tag("audio"), Tag("video")}
	t.Add(objects...)
	return
}

func TT(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("tt"))
	t.Add(objects...)
	return
}

func UL(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("ul"))
	t.Add(objects...)
	return
}

func VAR(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("var"))
	t.Add(objects...)
	return
}

func VIDEO(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("video"))
	t.Add(objects...)
	return
}

func WBR(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("wbr"), SelfClosing, Inline)
	t.Add(objects...)
	return
}
