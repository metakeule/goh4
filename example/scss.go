package main

import (
	"fmt"
	. "github.com/metakeule/goh4"
	"github.com/metakeule/goh4/scss"
	attr "github.com/metakeule/goh4/selectors/attributematcher"
	comb "github.com/metakeule/goh4/selectors/combinators"
	pseudo "github.com/metakeule/goh4/selectors/pseudoclasses"
	. "github.com/metakeule/goh4/styles"
)

func main() {
	topId := Id("top")
	activeClass := Class("active")
	blueBorder := BorderColor("blue")

	a := A(activeClass, topId)

	css := Css{}

	rule1, _ := Rule(a, blueBorder)
	// fmt.Println(rule1)

	rule2, _ := Rule(Selector(a, pseudo.Hover()), blueBorder)
	// fmt.Println(rule2)

	rule3, _ := Rule(
		comb.Child(Ul(topId), Li(activeClass)),
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
			Textarea(),
			activeClass,
			attr.BeginsWith("name", "comment"),
		),
		blueBorder,
	)

	rule4.Nest(
		comb.Descendant(
			Body(Class("firefox")),
			scss.Parent,
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

	rule5, _ := Rule(Li(Id("navbar")), scss.Include(rounded.Name, "top", "left"))
	// fmt.Println(rule5)

	rule6, _ := Rule(Id("footer"), scss.Include(rounded.Name, "top", "left", "5px"))
	rule7, _ := Rule(Id("sidebar"), scss.Include(rounded.Name, "top", "left", "8px"))
	// fmt.Println(rule6)
	// fmt.Println(rule7)

	rule8, _ := Rule(
		Id("footer"),
		BackgroundColor(scss.Call("lighten", yellow.Name(), PERCENT(20)).String()),
	)
	// fmt.Println(rule8)

	rule9, _ := Rule(Context(activeClass, a, topId), blueBorder)

	css = append(css, rule1, rule2, rule3, yellow, rule4, rounded, rule5, rule6, rule7, rule8, rule9)
	fmt.Println(css)

	fmt.Println(blueBorder.Value)

	fmt.Println(Div(Styles(Display(NONE)), "hiho"))
	fmt.Println(Div(Display(NONE), FontWeight(BOLD), "hiho"))
}
