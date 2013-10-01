package tag

import (
	. "github.com/metakeule/goh4"
	"io"
	"net/http"
)

// pseudo element for placeholder
func Doc(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("doc"), WithoutDecoration)
	t.Add(objects...)
	return
}

type DocType struct {
	*Element
	doctype        string
	htmlAttributes string
}

func docType(doctypeString string, objects ...interface{}) (d *DocType) {
	e := Doc(objects...)
	return &DocType{e, doctypeString, ""}
}

func docTypeXml(doctypeString string, objects ...interface{}) (d *DocType) {
	e := Doc(objects...)
	return &DocType{e, doctypeString, ` xmlns="http://www.w3.org/1999/xhtml"`}
}

func (ø *DocType) String() string {
	return ø.doctype + "\n<html" + ø.htmlAttributes + ">\n" + ø.Element.String() + "\n</html>\n"
}

func (ø *DocType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ø.String()))
}

func (ø *DocType) WriteTo(w io.Writer) {
	w.Write([]byte(ø.String()))
}

func (ø *DocType) AsTemplate() *Template {
	return NewTemplate(ø)
}

func (ø *DocType) Compile(name string) *CompiledTemplate {
	return ø.AsTemplate().MustCompile(name)
}

/*
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
//  t := NewTemplate(Body(Div(Id("content"))))
//  t.Assign("content", P(Text("here we go")))
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

*/

// the DocTypes are taken from http://www.w3.org/QA/2002/04/valid-dtd-list.html#DTD

func HTML4_01Strict(objects ...interface{}) (t *DocType) {
	return docType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
   "http://www.w3.org/TR/html4/strict.dtd">`, objects...)
}

func HTML4_01Transitional(objects ...interface{}) (t *DocType) {
	return docType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/loose.dtd">`, objects...)
}

func HTML4_01Frameset(objects ...interface{}) (t *DocType) {
	return docType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN"
   "http://www.w3.org/TR/html4/frameset.dtd">`, objects...)
}

func XHTML1_0Strict(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`, objects...)
}

func XHTML1_0Transitional(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`, objects...)
}

func XHTML1_0Frameset(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`, objects...)
}

func XHTML1_1(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
   "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`, objects...)
}

func XHTML1_1Basic(objects ...interface{}) (t *DocType) {
	return docTypeXml(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN"
    "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`, objects...)
}

func HTML5(objects ...interface{}) (t *DocType) {
	return docType(`<!DOCTYPE HTML>`, objects...)
}

func MathML2_0(objects ...interface{}) (t *DocType) {
	return docType(`<!DOCTYPE math PUBLIC "-//W3C//DTD MathML 2.0//EN"
  "http://www.w3.org/Math/DTD/mathml2/mathml2.dtd">`, objects...)
}

func MathML1_01(objects ...interface{}) (t *DocType) {
	return docType(`<!DOCTYPE math SYSTEM
  "http://www.w3.org/Math/DTD/mathml1/mathml.dtd">`, objects...)
}
