goh4
====

[![Build Status](https://secure.travis-ci.org/metakeule/goh4.png)](http://travis-ci.org/metakeule/goh4)

type save html and css generation for golang with functions, e.g.

	font := NewCss(
		Class("default-font"),
		Style("font-size","20","font-weight","normal"))

	css := NewCss(
		Class("yellow-button"),
		Tags("a","button"),
		Style("background-color","yellow"),
		font)

	a := A(css)

	doc := NewTemplate(
		Doc(
			Head(),
			Body(a)))

	doc.AddCss(css)

	fmt.Ptrintln(doc)

results in

	<head>
		<style>
			a.yellow-button,
			button.yellow-button {
				background-color: yellow;
				font-weight: normal;	\/* inherited from ».default-font« *\/
				font-size: 20;	\/* inherited from ».default-font« *\/
			}
		</style>
	</head>
	<body><a class="yellow-button"></a></body>

see the Documentation at http://godoc.org/github.com/metakeule/goh4

TODO
====

- more element contruction helpers
- check element contruction helpers against w3c specs
- contruction helpers for the various doctypes
- support for html tidy
- correct handling of css contexts for matching elements, needs a css selector parser
- test all the matchers
- add tools with
	- initializr / html boilerplate
	- jquery initialization tool (takes the path of the affected element and executes a js function on it, with and optional json config, writes it in the layout on document ready)
