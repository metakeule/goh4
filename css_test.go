package goh4

import (
	"fmt"
	"testing"
)

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
