package initializr

import (
	. "github.com/metakeule/goh4"
)

var exampleText string = `Donec id elit non mi porta gravida at eget metus.
													Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh,
													ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna
													mollis euismod. Donec sed odio dui.`

func exampleSpan4(text string) (t *Element) {
	t = Div(Class("span4"))
	t.Add(
		H(2, "heading"),
		P(text),
		P(
			A(
				"View details »",
				Class("btn"),
				Attr("href", "#"))),
	)
	return
}

func SetupExampleMain(ø *Initializr) {
	hero := Div(Class("hero-unit"))
	hero.Add(
		H(1, "Hello, world!"),
		P(`This is a template for a simple marketing or informational
			website. It includes a large callout called the hero unit
			and three supporting pieces of content. Use it as a starting
			point to create something more unique.`),
		P(
			A(
				"Learn more »",
				Class("btn"),
				Class("btn-primary"),
				Class("btn-large"))))

	row := Div(Class("row"))
	row.Add(
		exampleSpan4(exampleText),
		exampleSpan4(exampleText),
		exampleSpan4(exampleText))

	ø.Main.Set(Html(`<!-- Main hero unit for a primary marketing message or call to action -->`))
	ø.Main.Add(
		hero,
		row,
		Hr(),
		Footer(Html("&copy; Company 2012")))
}
