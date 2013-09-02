package goh4

import (
	"fmt"
)

type Templatable interface {
	String() string
	Tag() string
	Path() string
	Any(m Matcher) (r *Element)
	Add(objects ...interface{}) (err error)
}

type Template struct {
	Element          Templatable //*Element
	locals           map[Id]Stringer
	placeholderCache map[Id]*Element
}

// creates a new template with the given element as root
func NewTemplate(t Templatable) *Template {
	return &Template{
		Element:          t,
		locals:           map[Id]Stringer{},
		placeholderCache: map[Id]*Element{},
	}
}

func (ø *Template) String() string {
	return ø.Element.String()
}

func (ø *Template) Add(objects ...interface{}) (err error) {
	return ø.Element.Add(objects...)
}

// merges the locals to the Templates
func (ø *Template) merge() {
	for k, v := range ø.locals {
		oldParent := ø.placeholderCache[k].Parent()

		elem, isElement := v.(*Element)
		if isElement {
			elem.SetParent(oldParent)
			ø.placeholderCache[k].SetContent(elem)
		} else {
			ø.placeholderCache[k].SetContent(v)
		}

	}
}

// caches the Stringer
func (ø *Template) cacheFragment(id Id) (err error) {
	h := ø.Element.Any(Id(id))
	if h == nil {
		return fmt.Errorf("element with id %v not found in %s", id, ø.Element.String())
	}
	ø.placeholderCache[id] = h
	return
}

// replaces the content of an child Element with the given id with Stringer e.g.
//
// 	t := NewTemplate(Body(Div(Id("content"))))
// 	t.Assign("content", P(Text("here we go")))
//
// results in <body><div id="content"><p>here we go</p></div></body>
func (ø *Template) Assign(id Id, html interface{}) (err error) {
	if ø.placeholderCache[id] == nil {
		if err = ø.cacheFragment(id); err != nil {
			return err
		}
	}
	switch v := html.(type) {
	case string:
		ø.locals[id] = Text(v)
	default:
		ø.locals[id] = v.(Stringer)
	}

	ø.merge()
	return
}

// add css to the head of the template
// returns an error if Element is no doc or has no head child
func (ø *Template) AddCss(css ...Stringer) (err error) {
	if ø.Element.Tag() != "doc" {
		return fmt.Errorf("can't add Css only to doc pseudotag, not %s", ø.Element.Tag())
	}

	style := ø.Element.Any(Tag("style"))
	if style == nil {
		head := ø.Element.Any(Tag("head"))
		if head == nil {
			return fmt.Errorf("no head element present in %s", ø.Element.Path())
		}
		style = NewElement(Tag("style"), Invisible, WithoutEscaping)
		head.Add(style)
	}

	for _, cs := range css {
		style.Add(Html(cs.String()))
	}
	return
}
