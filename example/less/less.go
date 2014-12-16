package main

import (
	"fmt"
	"gopkg.in/metakeule/goh4.v5"
	. "gopkg.in/metakeule/goh4.v5/css"
	. "gopkg.in/metakeule/goh4.v5/less"
	"gopkg.in/metakeule/goh4.v5/less/fn"
	. "gopkg.in/metakeule/goh4.v5/sel/pseudo"
	. "gopkg.in/metakeule/goh4.v5/styl"
	. "gopkg.in/metakeule/goh4.v5/tag"
)

func main() {

	color := Var("color").Val("#4D926F")

	css1 := Css(
		color,
		Css(ID("header"), Color(color.String())),
		Css(H2(), Color(color.String())),
	)

	fmt.Println(css1)

	/*
		@color: #4D926F;

		#header {
		  color: @color;
		}
		h2 {
		  color: @color;
		}
	*/

	radius := Param("radius").Default(Px(5))

	roundedCorners := Mixin(CLASS("rounded-corners"), radius)

	roundedCorners.Style(
		BorderRadius(radius.String()),
		goh4.Style{"-webkit-border-radius", radius.String()},
		goh4.Style{"-moz-border-radius", radius.String()},
	)

	css2 := Css(
		roundedCorners,
		Css(ID("header"), Call(roundedCorners.String())),
		Css(ID("footer"), Call(roundedCorners.String(), Px(10))),
	)

	fmt.Println(css2)

	/*
	   .rounded-corners (@radius: 5px) {
	     border-radius: @radius;
	     -webkit-border-radius: @radius;
	     -moz-border-radius: @radius;
	   }

	   #header {
	     .rounded-corners;
	   }
	   #footer {
	     .rounded-corners(10px);
	   }
	*/

	css3 := Css(
		ID("header"),
		Nest(
			H1(),
			FontSize(Px(26)),
			FontWeight(Bold_),
		),
		Nest(
			P(),
			FontSize(Px(12)),
			Nest(
				A(),
				TextDecoration(None_),
				Nest(
					Super(Hover()),
					BorderWidth(Px(1)),
				),
			),
		),
	)

	fmt.Println(css3)
	/*
		#header {
		  h1 {
		    font-size: 26px;
		    font-weight: bold;
		  }
		  p {
		    font-size: 12px;
		    a {
		      text-decoration: none;
		      &:hover {
		        border-width: 1px;
		      }
		    }
		  }
		}
	*/

	css4 := Css(
		H2(),
		Color(fn.LIGHTEN_("red")),
	)
	fmt.Println(css4)
}
