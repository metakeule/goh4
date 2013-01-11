package goh4

import "fmt"

type Template struct {
	*Element
	locals           map[Id]Stringer
	placeholderCache map[Id]*Element
}

// creates a new template with the given element as root
func NewTemplate(t *Element) *Template {
	return &Template{
		Element:          t,
		locals:           map[Id]Stringer{},
		placeholderCache: map[Id]*Element{},
	}
}

// merges the locals to the Templates
func (ø *Template) merge() {
	for k, v := range ø.locals {
		oldParent := ø.placeholderCache[k].Parent()

		elem, isElement := v.(*Element)
		if isElement {
			elem.SetParent(oldParent)
			ø.placeholderCache[k].Set(elem)
		} else {
			ø.placeholderCache[k].Set(v)
		}

	}
}

// caches the Stringer
func (ø *Template) cacheFragement(id Id) (err error) {
	found, h := ø.Element.Any(Id(id))
	if !found {
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
func (ø *Template) Assign(id Id, html Stringer) (err error) {
	if ø.placeholderCache[id] == nil {
		if err = ø.cacheFragement(id); err != nil {
			return err
		}
	}
	ø.locals[id] = html
	ø.merge()
	return
}

// add css to the head of the template
// returns an error if Element is no doc or has no head child
func (ø *Template) AddCss(css Stringer) (err error) {
	if ø.Element.tag != "doc" {
		return fmt.Errorf("can't add Css only to doc pseudotag, not %s", ø.Element.Tag())
	}

	found, style := ø.Element.Any(Tag("style"))
	if !found {
		found, head := ø.Element.Any(Tag("head"))
		if !found {
			return fmt.Errorf("no head element present in %s", ø.Element.Path())
		}
		style = NewElement(Tag("style"), Invisible, WithoutEscaping)
		head.Add(style)
	}
	style.Add(Html(css.String()))
	return
}
