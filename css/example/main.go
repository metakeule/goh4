package main

import (
	"fmt"
	. "gopkg.in/metakeule/goh4.v5/css"
	. "gopkg.in/metakeule/goh4.v5/styl"
	. "gopkg.in/metakeule/goh4.v5/styl/color"
	. "gopkg.in/metakeule/goh4.v5/tag"
)

var bestClass = CLASS("best")
var mainId = ID("main")
var newsClass = CLASS("news")

func main() {
	css := Css(DIV(mainId, bestClass),
		BackgroundColor(GREEN),

		Css(SECTION(newsClass),
			BorderBottom(Px(2), GREEN, Dotted_),

			Css(ARTICLE(),
				FontSize(Px(12)),

				Css(H2(),
					FontSize(Px(25)),

					Css(P(),
						MarginBottom(Px(10)),
					),
				),
			),
		),
	)

	fmt.Println(css)
}
