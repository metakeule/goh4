package main

import (
	"fmt"
	. "github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/attr"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	// "github.com/metakeule/templ"
)

var BodyClass = Class("body-class").Placeholder()
var BodyId = Id("body-id").Placeholder()
var Content = Html("content").Placeholder()
var ContentComment = Comment("content-comment").Placeholder()
var Greeting = Text("greeting-text").Placeholder()
var AltText = Alt("img-alt").Placeholder()

type Figure struct {
	FirstName string   `greet:"-"`     // no defaults
	LastName  string   `greet:"-"`     // no defaults
	Greeting  Stringer `greet:"-"`     // no defaults
	Width     int      `greet:"400px"` // defaults to 400px
}

/*
type figurePlaceholders struct {
	FirstName template.Placeholder `goh4:"text"` // goh4 is the key for the goh4 types, the value must be a key of goh4.Transformer
	LastName  template.Placeholder `goh4:"text"`
	Width     template.Placeholder `goh4:"text"`
	Greeting  template.Placeholder `goh4:"html"`
}
*/

//var FigurePlaceholders = &figurePlaceholders{}
//var _ = FillStruct(FigurePlaceholders)

func main() {
	t :=
		BODY(
			BodyClass, BodyId,
			H1(Greeting),
			P(
				Class("content"),
				ContentComment,
				Content,
				ImgSrc(
					"/testimg.png",
					AltText,
					//					Width(FigurePlaceholders.Width.String()),
				),
			),
			//			H2(FigurePlaceholders.LastName),
			//			FigurePlaceholders.Greeting,
		)
	fmt.Println(t)

	template := t.Compile("template")

	/*
		donald := Figure{
			FirstName: "Donald",
			LastName:  "<Duck>",
			Greeting:  P("Are you fine?"),
		}
	*/
	//nu := template.New()
	//	nu.Merge(donald, "greet", FigurePlaceholders)
	//fmt.Printf("%T\n", nu)

	nu := template.MustReplace(
		BodyClass.Set("my-body-class"),
		BodyId.Set("my-body-id"),
		Greeting.Set("<my greeting>"),
		ContentComment.Set("my content comment"),
		Content.Set(SPAN("my content")),
		AltText.Set("my <alttext>"),
		// uncomment this line to get the defaults of 400px
		//		FigurePlaceholders.Width.Set("800px"),
	)

	fmt.Println(nu)
}
