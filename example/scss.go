package main

import (
	"fmt"
	. "github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/css"
	"github.com/metakeule/goh4/scss"
	"github.com/metakeule/goh4/sel/attmatch"
	"github.com/metakeule/goh4/sel/pseudo"
	. "github.com/metakeule/goh4/styl"
	. "github.com/metakeule/goh4/tag"
)

func main() {
	topId := ID("top")
	activeClass := CLASS("active")
	blueBorder := BorderColor("blue")

	a := A(activeClass, topId)

	rule1, _ := Rule(a, blueBorder)
	// fmt.Println(rule1)

	rule2, _ := Rule(Selector(a, pseudo.Hover()), blueBorder)
	// fmt.Println(rule2)

	rule3, _ := Rule(
		Child(UL(topId), LI(activeClass)),
		blueBorder,
	)
	// fmt.Println(rule3)

	yellow := scss.Var("yellow").Val("yellow")

	rule4, _ := Rule(
		Selector(topId, activeClass, pseudo.Link()),
		BackgroundColor(yellow.Name()),
		blueBorder,
	)
	rule4.Nest(
		Selector(
			TEXTAREA(),
			activeClass,
			attmatch.BeginsWith("name", "comment"),
		),
		blueBorder,
	)

	rule4.Nest(
		Selector(
			scss.Parent,
			pseudo.Hover(),
		),
		blueBorder,
	)

	// Import("colors"),
	/*
		fmt.Print(
			yellow.Definition(), ";\n",
			rule4,
			"\n")
	*/
	vert := scss.Param("vert")
	horz := scss.Param("horz")
	radius := scss.Param("radius").Default("10px")
	rounded := scss.Mixin("rounded", vert, horz, radius)
	rounded.Style(
		Style{fmt.Sprintf("border-#{%s}-#{%s}-radius", vert, horz), radius.Name()},
		Style{fmt.Sprintf("-moz-border-radius-#{%s}-#{%s}", vert, horz), radius.Name()},
		Style{fmt.Sprintf("-webkit-border-#{%s}-#{%s}-radius", vert, horz), radius.Name()},
	)
	// fmt.Println(rounded)

	rule5, _ := Rule(LI(ID("navbar")), scss.Include(rounded.Name, "top", "left"))
	// fmt.Println(rule5)

	rule6, _ := Rule(ID("footer"), scss.Include(rounded.Name, "top", "left", "5px"))
	rule7, _ := Rule(ID("sidebar"), scss.Include(rounded.Name, "top", "left", "8px"))
	// fmt.Println(rule6)
	// fmt.Println(rule7)

	rule8, _ := Rule(
		ID("footer"),
		BackgroundColor(scss.Call("lighten", yellow.Name(), Percent(20)).String()),
	)
	// fmt.Println(rule8)

	rule9, _ := Rule(Descendant(activeClass, a, topId), blueBorder)

	css := Css(rule1, rule2, rule3, rule4, rule5, rule6, rule7, rule8, rule9)
	fmt.Println(yellow, "\n", rounded, "\n", css)

	// fmt.Println(blueBorder.Value)

	fmt.Println(DIV(Styles(Display(None_)), "hiho"))
	fmt.Println(DIV(Display(None_), FontWeight(Bold_), "hiho"))
}
