// Copyright 2013 Marc René Arns. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package goh4 constructs html element trees (something like DOM), html templates and integrates them with css rules.
All of them are done by using types, not strings. This makes reuse and shortcuts easy.

It is easy to extend goh4 and build upon it and you may bring in some html and css strings as well.

Elements


There are some helper functions to generate elements for you. For instance

	element := A()

gives you an anchor element. Here is how A() is defined and how you may define your own shortcuts.

	func A(objects ...Stringer) (t *Element) {
		t = NewElement(Tag("a"), Inline)
		t.Set(objects...)
		return
	}

A() calls NewElement() with a typed Tag (created from a string) and a flag that indicates that it's an
inline element. All flags are defined as constants and may be combined with the bitwise or | or given as
seperate parameters.

Then A() sets up the element by passing the optional Stringer to the Set() method. This allows you do
some neat things, for instance

	a := A(
		Class("button"),
		Id("start"),
		Text("begin"))

	a.String() // <a class="button" id="start">begin</a>

	a = A(
		Attrs{"href","#","_target","_blank"},
		Text("something"))

	a.String() // <a href="#" target="_blank">something</a>

	a = A(
		Styles{"color","red"},
		Strong(
			Html("not <escaped>!")))
	a.String() // <a style="color:red;"><strong>not <escaped></strong></a>

Then you may get all the nice methods of element

	children := a.Children() // only the Elements
	inner := a.InnerHtml()   // everything inside
	buttons := a.All(Class("button"))
	start := a.Any(Id("start"))
	...

You may define your own Matcher and pass it to All() and Any().

Templates and Css


Template is an element that can Assign() anything that can be inside
an element to any element within the tree, that has the given Id().

	content := Div(Id("content"))
	body := NewTemplate(Body(content))
	body.Assign(Id("content"), P(Text("here I am")))
	body.String() // <body><div id="content"><p>here I am</p></div></body>

Template may also add css to your head element. In this case you need to
use the Doc() pseudo element, that holds the root of the document.
And you need a head obviously.

	doc := NewTemplate(Doc(Head())
	doc.AddCss("body{ color: red; }")
	doc.String() // <head><style>body { color: red; }</style></head>

Element and Template are structs you can inherit from them, e.g.

	type Layout struct {
		*Template
	}

	func (ø *Layout) String() string {
		// do your thing here
	}

If you want type save reusable css, you can use Css.

	fontsize := NewCss(Class("default-font-size"),Styles{"font-size","20"})
	css := NewCss(
		Class("yellow-button"),
		Tags{"a","button"},
		Styles{"background-color","yellow"},
		fontsize)

and then you might apply it to your elements.

	a := A()
	a.ApplyCss(css)
	a.String() // <a class="yellow-button"></a>

don't forget to put your css into the template
(we don't put fontsize into it, since it is only a building block)

	doc := NewTemplate(Doc(Head(),Body(a)))
	doc.AddCss(css)
	doc.String()

this results in the following (no auto indentation at the moment, sorry):

	<head>
		<style>
			a.yellow-button,
			button.yellow-button {
				background-color: yellow;
				font-size: 20;	\/* inherited from ».default-font-size« *\/
			}
		</style>
	</head>
	<body><a class="yellow-button"></a></body>
*/
package goh4
