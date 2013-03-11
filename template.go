package goh4

import (
	"fmt"
	"github.com/metakeule/fastreplace"
)

type Template struct {
	*Element
	locals           map[Id]Stringer
	placeholderCache map[Id]*Element
	Delimiter        string
}

// creates a new template with the given element as root
func NewTemplate(t *Element) *Template {
	return &Template{
		Element:          t,
		locals:           map[Id]Stringer{},
		placeholderCache: map[Id]*Element{},
		Delimiter:        "@@",
	}
}

type CompiledTemplate struct {
	*fastreplace.FReplace
	Template *Template
}

// returns a *CompiledTemplate that is a FReplace (see github.com/metakeule/fastreplace)
// the template will be build as string and then searched for placeholders that
// start and end with the ø.Delimiter (defaults to "@@")
// then the various methods of fastreplace.FReplace can be used to merge
// the placeholders of the template with the values
// if you need to change the original template again, you can get it via CompiledTemplate.Template
// then call Compile() again to get a new CompiledTemplate
func (ø *Template) Compile() (c *CompiledTemplate) {
	return &CompiledTemplate{
		FReplace: fastreplace.NewString(ø.Delimiter, ø.String()),
		Template: ø,
	}
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
func (ø *Template) cacheFragement(id Id) (err error) {
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
		if err = ø.cacheFragement(id); err != nil {
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
	if ø.Element.tag != "doc" {
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
