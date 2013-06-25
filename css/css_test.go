package css

import (
	// "fmt"
	"testing"
)

func TestRule(t *testing.T) {
	var expectedString string
	font, _ := Rule(Class("default"), Style{"font-size", "12px"})
	bg, _ := Rule(Class("bg"), Style{"background-color", "black"})
	bg = bg.Compose(font)
	special := Class("special")

	css, _ := Rule(
		Context(Id("myid"), Ul(special), Ol(special), Li(special)),
		Style{"color", "yellow"},
		Styles(Style{"width", "200"}, Style{"height", "300"}),
		Comment("no comment"))

	css = css.Compose(bg)

	expectedString = "#myid ul.special,\n#myid ol.special,\n#myid li.special"
	if css.Selectors[0].Selector() != expectedString {
		err(t, "incorrect selector1", css.Selectors[0].Selector(), expectedString)
	}

	styles := css.Styles

	expectedString = "color: yellow"
	if styles[0].Style() != expectedString {
		err(t, "incorrect style", styles[0].Style(), expectedString)
	}

	expectedString = "width: 200"
	if styles[1].Style() != expectedString {
		err(t, "incorrect style", styles[1].Style(), expectedString)
	}

	expectedString = "height: 300"
	if styles[2].Style() != expectedString {
		err(t, "incorrect style", styles[2].Style(), expectedString)
	}

	expectedString = "background-color: black"
	if styles[3].Style() != expectedString {
		err(t, "incorrect style", styles[3].Style(), expectedString)
	}
	expectedString = "font-size: 12px"
	if styles[4].Style() != expectedString {
		err(t, "incorrect style", styles[4].Style(), expectedString)
	}
}
