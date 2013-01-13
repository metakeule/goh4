package goh4

import (
	"fmt"
	"testing"
)

// create an element with a simple self defined tag
func ExampleNewCss() {
	css := NewCss(
		Class("active"),
		Tags{"a:hover", "li"}, // single tag with Tag("li")
		Comment("highlight activation"),
		Context("#content"),   // is simply prefixed to the selector
		Style{"color", "red"}) // multiple styles with Styles{"color","red","width","200"}

	fmt.Println(css)
	// Output: /* highlight activation */
	// #content a:hover.active,
	// #content li.active {
	// 	color: red;
	// }
}

func TestCssMatch(t *testing.T) {
	css1 := NewCss(Class("beautyful"))

	if !css1.Matches(A(css1)) {
		err(t, "incorrect css matches", false, true)
	}

	if css1.Matches(A()) {
		err(t, "incorrect css matches", true, false)
	}

	css2 := NewCss(Class("beautyful"), Tag("a"))

	if !css2.Matches(A(css2)) {
		err(t, "incorrect css matches with tag", false, true)
	}

	if css2.Matches(Li(css2)) {
		err(t, "incorrect css matches with tag", true, false)
	}
}

func Example_Css_template() {
	fontsize := NewCss(Class("default-font-size"), Styles{"font-size", "20"})
	css := NewCss(
		Class("yellow-button"),
		Tags{"a", "button"},
		Comment("make ugly buttons"),
		fontsize)
	a := A()
	a.ApplyCss(css)
	a.String()
	doc := NewTemplate(Doc(Head(), Body(a)))
	doc.AddCss(css)
	fmt.Println(doc)
	// Output: <head>
	// <style>
	//
	// /* make ugly buttons */
	// a.yellow-button,
	// button.yellow-button {
	// 	font-size: 20;	/* inherited from ».default-font-size« */
	// }
	// </style>
	// </head>
	//
	// <body><a class="yellow-button "></a></body>
}

// inherit from others
func ExampleNewCss_inherited() {
	red := NewCss(Style{"color", "red"})

	highlight := NewCss(
		Class("highlight"),
		red)

	css := NewCss(
		Class("active"), highlight) // css inherits from red via highlight, multiple inheritance possible

	fmt.Println(css)
	// Output: .active {
	// 	color: red;	/* inherited from ».highlight« */
	// }
}

// overwrite inherited properties
func ExampleNewCss_overwritten() {
	red := NewCss(
		Class("redder"),
		Style{"color", "red"})

	css := NewCss(
		Class("active"),
		red,
		Style{"color", "green"}) // only the same keys are overwritten

	fmt.Println(css)
	// Output: .active {
	// 	color: green;
	// }
}

func TestNewCss(t *testing.T) {
	var exspectedString string
	font := NewCss(Class("default"), Style{"font-size", "12px"})
	bg := NewCss(Class("bg"), Style{"background-color", "black"}, font)

	css := NewCss(
		Context("#myid"),
		Tag("ul"),
		Tags{"ol", "li"},
		Style{"color", "yellow"},
		Styles{"width", "200", "height", "300"},
		Class("special"), Comment("no comment"),
		bg)

	exspectedString = "#myid"
	if css.Context.String() != exspectedString {
		err(t, "incorrect context", css.Context, exspectedString)
	}

	exspectedString = "special"
	if css.Class() != "special" {
		err(t, "incorrect class", css.Class(), exspectedString)
	}

	exspectedString = "#myid ul.special,\n#myid ol.special,\n#myid li.special"
	if css.Selector() != exspectedString {
		err(t, "incorrect selector", css.Selector(), exspectedString)
	}

	var expectedOwnStyles = map[string]string{
		"color":  "yellow",
		"width":  "200",
		"height": "300",
	}

	var expectedInheritedStyles = map[string]string{
		"background-color": "black",
		"font-size":        "12px",
	}

	styles := css.Styles()

	for k, v := range expectedOwnStyles {
		if styles[k] != v {
			err(t, "incorrect style "+k, styles[k], v)
		}
	}

	for k, v := range expectedInheritedStyles {
		if styles[k] != v {
			err(t, "incorrect style (inherited, but should be included)"+k, styles[k], v)
		}
	}

	styles = css.OwnStyles()

	for k, v := range expectedOwnStyles {
		if styles[k] != v {
			err(t, "incorrect own style "+k, styles[k], v)
		}
	}

	styles = css.InheritedStyles()

	for k, v := range expectedInheritedStyles {
		if styles[k] != v {
			err(t, "incorrect inherited style "+k, styles[k], v)
		}
	}

}

func makeCss(outer string, inner string) string {
	return fmt.Sprintf(outer, inner)
}

func TestCssString(t *testing.T) {
	s := Styles{"color", "green"}
	ss := "color:green;"
	cl := Class("t")
	ctx := Context("#my")
	tags := Tags{"ul", "ol"}
	cmt := Comment("hiho")

	var expected = map[*Css]string{
		NewCss(cl, s):                 makeCss(".t{%s}", ss),
		NewCss(ctx, cl, s):            makeCss("#my .t{%s}", ss),
		NewCss(ctx, cl, s, tags):      makeCss("#my ul.t,#my ol.t{%s}", ss),
		NewCss(ctx, cl, s, tags, cmt): makeCss("/* hiho */ #my ul.t,#my ol.t{%s}", ss),
	}

	for k, v := range expected {
		res := removeWhiteSpace(k.String())
		if res != removeWhiteSpace(v) {
			err(t, "incorrect css string ", res, removeWhiteSpace(v))
		}
	}
}
