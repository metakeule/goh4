package main

import (
	"fmt"
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"net/http"
)

var (
	content__ = HTML("content").Placeholder()
	title__   = goh4.Text("title").Placeholder()
	layout    = HTML5(
		HEAD(TITLE(title__), CharsetUtf8()),
		BODY(
			DIV(CLASS("header"), "header"),
			DIV(CLASS("content"), content__),
		),
	).Compile()

	static = XHTML1_0Transitional(
		HEAD(
			TITLE("Hallo"),
			HttpEquivUtf8(),
		),
		BODY(
			DIV(CLASS("header"), "header"),
			DIV(CLASS("content"),
				AHref("/a", "A"),
				BR(),
				AHref("/b", "B"),
			),
		),
	)
)

type Layout struct {
	Content *goh4.Element
	Title   string
}

func (ø Layout) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	layout.Replace(
		content__.Set(ø.Content),
		title__.Set(ø.Title),
	).WriteTo(rw)
}

var (
	_ = fmt.Println
)

func init() {

}

func main() {
	http.Handle("/a", &Layout{H1("a"), "A - Titel"})
	http.Handle("/b", &Layout{H1("b"), "B - Titel"})
	http.Handle("/", static)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
