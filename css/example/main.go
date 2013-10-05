package main

import (
	"fmt"
	. "github.com/metakeule/goh4/css"
	. "github.com/metakeule/goh4/styl"
	. "github.com/metakeule/goh4/styl/color"
	. "github.com/metakeule/goh4/tag"
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
