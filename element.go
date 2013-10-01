package goh4

import (
	"bytes"
	"fmt"
	"github.com/metakeule/templ"
	"html"
	"io"
	"net/http"
	"strings"
)

type flag int

const (
	_                      = iota
	hasDefaults       flag = 1 << iota // element has default flags
	IdForbidden                        // element should not have an id attribute
	ClassForbidden                     // element should not have a class attribute
	SelfClosing                        // element is selfclosing and contains no content
	Inline                             // element is an inline element (only for visible elements)
	FormField                          // element is a field of a form
	Invisible                          // element doesn't render anything visible
	WithoutEscaping                    // element does not escape inner Text
	WithoutDecoration                  // element just prints the InnerHtml
)

var flagNames = map[flag]string{
	hasDefaults:       "hasDefaults",
	IdForbidden:       "IdForbidden",
	ClassForbidden:    "ClassForbidden",
	SelfClosing:       "SelfClosing",
	Inline:            "Inline",
	FormField:         "FormField",
	Invisible:         "Invisible",
	WithoutEscaping:   "WithoutEscaping",
	WithoutDecoration: "WithoutDecoration",
}

func (ø flag) String() string {
	return flagNames[ø]
}

type Pather interface {
	Path() string
}
type Tager interface {
	Tag() string
}

// a Csser might be applied as Css to an Element
type Csser interface {
	Matcher
	Class() string
}

// an Elementer might be parent of an Element
// by implementing a type that fulfills this interface
// you might peek into the execution.
// when String() method is called, the html of the
// tree is built and when SetParent() it is embedded in another Elementer
// it could be combined with the Pather interface that allows you to modify specific
// css selectors for any children Elements
type Elementer interface {
	Stringer
	Tager
	IsParentAllowed(Tager) bool
	SetParent(Pather)
}

// the base of what becomes a tag when printed
type Element struct {
	attributes map[string]string
	Comment    Comment
	parent     Pather
	classes    []Class
	id         Id
	inner      []Stringer
	ParentTags tags
	style      map[string]string
	flags      flag
	tag        Tag
}

// contruct a new element with some flags.
//
// the tag constructors A(), Body(),... use these method, see tags.go file for examples
//
// use it for your own tags
//
// the following flags are supported
//
// 	IdForbidden                        // element should not have an id attribute
// 	ClassForbidden                     // element should not have a class attribute
// 	SelfClosing                        // element is selfclosing and contains no content
// 	Inline                             // element is an inline element (only for visible elements)
// 	Field                              // element is a field of a form
// 	Invisible                          // element doesn't render anything visible
// 	WithoutEscaping                    // element does not escape inner Text
// 	WithoutDecoration                  // element just prints the InnerHtml
//
// see Add() and Set() methods for how to modify the Element
func NewElement(t Tag, flags ...flag) (ø *Element) {
	ø = &Element{
		attributes: map[string]string{},
		tag:        t,
		flags:      hasDefaults,
		style:      map[string]string{},
		ParentTags: tags{},
	}

	for _, flag := range flags {
		ø.flags = ø.flags | flag
	}
	return
}

func (ø *Element) NotEscape() *Element {
	ø.Add(WithoutEscaping)
	return ø
}

// checks if a given flag is set, e.g.
//
// 	Is(Inline)
//
// checks for the Inline flag
func (ø *Element) Is(f flag) bool {
	return ø.flags&f != 0
}

// returns the formfields
func (ø *Element) Fields() (fields []*Element) {
	return ø.All(FieldMatcher(0))
}

func (ø *Element) Parent() Pather {
	return ø.parent
}

// sets the attribute k to v as long as k is not "id" or "class"
// use SetId() to set the id and AddClass() to add a class
func (ø *Element) SetAttribute(k, v string) {
	if k != "id" && k != "class" {
		ø.attributes[k] = v
	}
}

// removes an attribute
func (ø *Element) RemoveAttribute(k string) {
	delete(ø.attributes, k)
}

func (ø *Element) Attribute(k string) string {
	return ø.attributes[k]
}

func (ø *Element) Attributes() map[string]string {
	return ø.attributes
}

// wraps the children with the given element
func (ø *Element) WrapChildren(wrapper *Element) {
	for _, s := range ø.inner {
		wrapper.Add(s)
	}
	ø.inner = []Stringer{}
	ø.Add(wrapper)
}

func (ø *Element) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ø.String()))
}

func (ø *Element) WriteTo(w io.Writer) {
	w.Write([]byte(ø.String()))
}

func (ø *Element) SetParent(parent Pather) {
	ø.parent = parent
}

func (ø *Element) Tag() string {
	return ø.tag.String()
}

func (ø *Element) Path() string {
	parentPath := ""
	if ø.parent != nil {
		parentPath = ø.parent.Path() + " "
	}
	return fmt.Sprintf(parentPath+"%s%s%s", ø.tag, ø.idPath(), ø.classPath(ø.classes))
}

// creates a Css with Context set to Path()
/*
func (ø *Element) NewCss(objects ...Stringer) *Css {
	objects = append(objects, Context(ø.Path()))
	return NewCss(objects...)
}
*/

func (ø *Element) idPath() (s string) {
	s = ""
	if !ø.Is(IdForbidden) && ø.id != "" {
		s = "#" + ø.id.String()
	}
	return
}

// return false if the given Parent tag is not allowed for Elements tag
func (ø *Element) IsParentAllowed(parent Tager) (allowed bool) {
	if len(ø.ParentTags) == 0 {
		return true
	}
	allowed = false
	for _, p := range ø.ParentTags {
		if p.String() == parent.Tag() {
			allowed = true
		}
	}
	return
}

// adds css properties to the style attribute, same keys are overwritten
func (ø *Element) addStyle(styls ...interface{}) {
	if ø.Is(Invisible) {
		// can't set style for invisible tag
		return
	}
	for _, s := range styls {
		switch v := s.(type) {
		case []Style:
			for _, st := range v {
				ø.style[st.Property] = st.Value
			}
		case Style:
			ø.style[v.Property] = v.Value
			/*
				case styles:
					for _, styl := range v {
						ø.style[styl.Key] = styl.Value
					}
			*/
		case string:
			a := strings.Split(v, ":")
			ø.style[a[0]] = strings.Replace(a[1], ";", "", 1)
		}
	}
}

// apply the css to the element, i.e. add the class of the Css to the Element
//
// returns an error if the Css has no class or if the Css will not match
// the Element because of its tags or its context
// beware that it will always fail if the Css has a context, because the
// matcher can't properly figure out if an element is within a context
// that remains to be done (but requires a proper css selector parser)
/*
func (ø *Element) ApplyCss(rules ...Csser) (err error) {
	for _, r := range rules {
		if r.Class() == "" {
			return fmt.Errorf("can't apply css rule without a class: %s", r)
		}

		if err := ø.AddClass(Class(r.Class())); err != nil {
			return err
		}

		if !r.Matches(ø) {
			return fmt.Errorf("tag %s not affected by css rule %#v", ø.Path(), r)
		}
	}
	return
}
*/

// returns an error if the Element is self closing
func (ø *Element) ensureContentAddIsAllowed() (err error) {
	if ø.Is(SelfClosing) {
		return fmt.Errorf("add not allowed for tag »%s«", ø.tag)
	}
	return
}

// adds Elementer to the inner content at position pos
func (ø *Element) AddAtPosition(pos int, v Elementer) (err error) {
	if pos < 0 || pos > len(ø.inner)-1 {
		return fmt.Errorf("position %v out of range", pos)
	}
	if err := ø.ensureParentIsSetAndContentIsAllowed(v); err != nil {
		return err
	}
	before := ø.inner[0:pos]
	after := ø.inner[pos:len(ø.inner)]

	newSlice := make([]Stringer, (1 + len(ø.inner)))
	newSlice[pos] = v
	for i := 0; i < len(before); i++ {
		newSlice[i] = before[i]
	}

	for i := 0; i < len(after); i++ {
		newSlice[len(before)+i+1] = after[i]
	}
	ø.inner = newSlice
	return
}

// make sure the parent is set and content is allowed
func (ø *Element) ensureParentIsSetAndContentIsAllowed(v Stringer) error {
	e, isElementer := v.(Elementer)
	if isElementer {
		if !e.IsParentAllowed(ø) {
			return fmt.Errorf("tag %s not allowed as parent of %s", ø.Tag(), e.Tag())
		}
		e.SetParent(ø)
	}
	return ø.ensureContentAddIsAllowed()
}

// set the Elementer to the inner content at position pos and overwrite the current content at that position
func (ø *Element) SetAtPosition(pos int, v Elementer) (err error) {
	if pos < 0 || pos > len(ø.inner)-1 {
		return fmt.Errorf("position %v out of range", pos)
	}
	if err := ø.ensureParentIsSetAndContentIsAllowed(v); err != nil {
		return err
	}
	ø.inner[pos] = v
	return
}

// sets the Elementer to the last position of the inner content and overwrites the current content at that position
//
// If you want to append to the inner content, use Add() instead
func (ø *Element) SetBottom(v Elementer) (err error) {
	if err := ø.ensureParentIsSetAndContentIsAllowed(v); err != nil {
		return err
	}
	ø.inner[len(ø.inner)-1] = v
	return
}

// returns the position of the Element in the inner content. if it could not be found, the last parameter is false
func (ø *Element) PositionOf(v *Element) (pos int, found bool) {
	m := &PositionMatcher{Element: v}
	_ = ø.Any(m)
	pos = m.Pos
	found = m.Found
	return
}

// adds Elementer at the position before the Element in the inner content
// the following elements are moved down
func (ø *Element) AddBefore(v *Element, nu Elementer) (err error) {
	pos, found := ø.PositionOf(v)
	if !found {
		return fmt.Errorf("could not find %#v", v)
	}
	if err := ø.ensureParentIsSetAndContentIsAllowed(nu); err != nil {
		return err
	}
	return ø.AddAtPosition(pos, nu)
}

// adds Elementer at the position before the Element in the inner content
// the following elements are moved down
func (ø *Element) AddAfter(v *Element, nu Elementer) (err error) {
	pos, found := ø.PositionOf(v)
	if !found {
		return fmt.Errorf("could not find %#v", v)
	}
	if err := ø.ensureParentIsSetAndContentIsAllowed(nu); err != nil {
		return err
	}
	if pos > len(ø.inner)-2 {
		return ø.Add(nu)
	} else {
		return ø.AddAtPosition(pos+1, nu)
	}
	return
}

// adds new inner content or properties based on Stringer objects and returns an error if changes could not be applied
//
// the following types are handled in a special way:
//
//  - Comment: sets the comment
//  - Style: set a single style
//  - Styles: sets multiple styles
//  - Attr: set a single attribute   // do not set id or class via Attr(s) directly, use Id() and Class() instead
//  - Attrs: sets multiple attribute
//  - Class: adds a class
//  - Id: sets the id
//  - *Css: applies the css, see ApplyCss()
//
// the following types are added to  the inner content:
//
// 	- Text: ís escaped if the WithoutEscaping flag isn't set
// 	- Html: is never escaped
//
// If the Stringer can be casted to an Elementer (as Element can), it is added to the inner content as well
// otherwise it is handled like Text(), that means any type implementing Stringer can be added as (escaped) text
func (ø *Element) Add(objects ...interface{}) (err error) {
	for _, o := range objects {
		switch v := o.(type) {
		case templ.Placeholder:
			e := ø.Add("@@" + v.Name() + "@@")
			if e != nil {
				return e
			}
			continue
		case *CompiledTemplate:
			e := ø.Add("@@" + v.Name() + "@@")
			if e != nil {
				return e
			}
			continue
		case Placeholder:
			var e error
			switch tp := v.Type().(type) {
			case Comment:
				e = ø.Add(Comment(v.String()))
			case Id:
				e = ø.Add(Id(v.String()))
			case Class:
				e = ø.Add(Class(v.String()))
			case Html:
				e = ø.Add(Html(v.String()))
			case Text:
				e = ø.Add(Text(v.String()))
			case SingleAttr:
				e = ø.Add(SingleAttr{tp.Key, v.String()})
			case Tag:
				ø.tag = Tag(v.String())
			case Style:
				e = ø.Add(Style{tp.Property, v.String()})
			default:
				e = fmt.Errorf("%#v (%T) unsupported placeholder type", v, v)
			}
			if e != nil {
				return e
			}
			continue
			//ø.inner = append(ø.inner, Html(v.Key()))
		case string:
			if err := ø.ensureContentAddIsAllowed(); err != nil {
				return err
			}
			if !ø.Is(WithoutEscaping) {
				v = html.EscapeString(v)
			}
			ø.inner = append(ø.inner, Text(v))
		case Text:
			if err := ø.ensureContentAddIsAllowed(); err != nil {
				return err
			}
			if !ø.Is(WithoutEscaping) {
				v = Text(html.EscapeString(string(v)))
			}
			ø.inner = append(ø.inner, v)
		case Html:
			if err := ø.ensureContentAddIsAllowed(); err != nil {
				return err
			}
			ø.inner = append(ø.inner, v)
		case Comment:
			ø.Comment = v

		case SingleAttr:
			ø.SetAttribute(v.Key, v.Value)

		case Attrs:
			for _, atr := range v {
				ø.SetAttribute(atr.Key, atr.Value)
			}
		case classes:
			for _, cl := range v {
				if err := ø.AddClass(cl); err != nil {
					return err
				}
			}
		case Class:
			if err := ø.AddClass(v); err != nil {
				return err
			}
		case Id:
			if err := ø.SetId(v); err != nil {
				return err
			}
		case Style:
			ø.addStyle(v)
			/*
				case styles:
					ø.addStyle(v)
			*/
			/*
				case *Css:
					if err := ø.ApplyCss(v); err != nil {
						return err
					}
			*/
		case []Style:
			ø.addStyle(v)
		case Stringer:
			if err := ø.ensureContentAddIsAllowed(); err != nil {
				return err
			}
			stringer := v
			/*
				stringer, ok := v.(Stringer)
				if !ok {
					return fmt.Errorf("%#v  is no Stringer", v)
				}
			*/
			if err := ø.ensureParentIsSetAndContentIsAllowed(stringer); err == nil {
				// no error: is an Elementer
				ø.inner = append(ø.inner, stringer)
			} else {
				// handle it like untyped string
				s := Text(stringer.String())

				if !ø.Is(WithoutEscaping) {
					s = Text(html.EscapeString(stringer.String()))
				}
				ø.inner = append(ø.inner, s)
			}
		default:
			return fmt.Errorf("%#v  is no Stringer", v)
		}
	}
	return
}

// clears the inner elements and strings
func (ø *Element) Clear() {
	ø.inner = []Stringer{}
}

func (ø *Element) Selecter(other ...Selecter) Selecter {
	return Selector(SelectorString(ø.Selector()), other...)
}

func (ø *Element) Placeholder() Placeholder {
	return Html(ø.String()).Placeholder()
}

func (ø *Element) Selector() string {
	sele := []string{ø.tag.Selector()}
	if ø.id != Id("") {
		sele = append(sele, ø.id.Selector())
	}
	for _, c := range ø.classes {
		sele = append(sele, c.Selector())
	}
	return strings.Join(sele, "")
}

func (ø *Element) AsTemplate() *Template {
	return NewTemplate(ø)
}

func (ø *Element) Compile(name string) *CompiledTemplate {
	return ø.AsTemplate().MustCompile(name)
}

// clears the inner object array
// and then calles Add() method to add content
//
// see Add() method for more details
func (ø *Element) SetContent(objects ...interface{}) (err error) {
	ø.inner = []Stringer{}
	return ø.Add(objects...)
}

// use this func to set the id of the Element,
// do not set it via Attr directly
// returns error if IdForbidden flag is set
func (ø *Element) SetId(id Id) (err error) {
	if ø.Is(IdForbidden) {
		return fmt.Errorf("id not allowed for tag %s", ø.tag)
	}
	ø.id = id
	return
}

func (ø *Element) classAttrString(classes []Class) (s string) {
	//s = ""
	var buffer bytes.Buffer
	for _, cl := range classes {
		//s += cl.String() + " "
		buffer.WriteString(cl.String() + " ")
	}
	return buffer.String()
}

func (ø *Element) classPath(classes []Class) (s string) {
	//s = ""
	var buffer bytes.Buffer
	for _, cl := range classes {
		//s += "." + cl.String()
		buffer.WriteString("." + cl.String())
	}
	return buffer.String()
}

func (ø *Element) styleAttrString(styles map[string]string) (s string) {
	var buffer bytes.Buffer
	//s = ""
	for k, v := range styles {
		//s += Style{k, v}.String()
		buffer.WriteString(Style{k, v}.String())
	}
	return buffer.String()
	//return
}

func (ø *Element) HasClass(class Class) bool {
	for _, c := range ø.classes {
		if c == class {
			return true
		}
	}
	return false
}

func (ø *Element) RemoveClass(class Class) {
	for i, c := range ø.classes {
		if c == class {
			// remove an element from a slice, according to
			// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/lYz8ftASMQ0
			copy(ø.classes[i:], ø.classes[i+1:])
			ø.classes = ø.classes[:len(ø.classes)-1]
		}
	}
}

// use this func to set the classes of the Element
// do not set them via Attr directly
func (ø *Element) SetClass(classes ...Class) {
	ø.classes = []Class{}
	ø.AddClass(classes...)
}

// use this func to add the classes of the tag,
// do not set it via Attr directly
func (ø *Element) AddClass(classes ...Class) (err error) {
	if ø.Is(ClassForbidden) {
		return fmt.Errorf("class not allowed for tag %s", ø.tag)
	}
	for _, cl := range classes {
		ø.classes = append(ø.classes, cl)
	}
	return
}

// use this method to get the classes since they won't show up in attributes
func (ø *Element) Classes() (c []Class) {
	return ø.classes
}

// use this method to get the id since it won't show up in attributes
func (ø *Element) Id() Id { return ø.id }

// returns the html with inner content (and the own tags if WithoutDecoration is not set)
func (ø *Element) String() (res string) {
	if ø.Is(WithoutDecoration) {
		return ø.InnerHtml()
	}

	commentpre := ""
	commentpost := ""
	if ø.Comment != "" {
		commentpre = fmt.Sprintf("<!-- Begin: %s -->", ø.Comment)
		commentpost = fmt.Sprintf("<!-- End: %s -->", ø.Comment)
	}
	if ø.Is(SelfClosing) {
		res = fmt.Sprintf("%s<%s%s />%s", commentpre, string(ø.tag), ø.attrsString(), commentpost)
	} else {
		str := "%s<%s%s>%s</%s>%s"
		//if !ø.Is(Inline) {
		//	str = "\n%s<%s%s>%s</%s>%s\n"
		//}
		res = fmt.Sprintf(str, commentpre, string(ø.tag), ø.attrsString(), ø.InnerHtml(), string(ø.tag), commentpost)
	}
	return
}

// prepare the id attribute for output
func (ø *Element) attrsString() (res string) {
	var buffer bytes.Buffer
	//res = ""
	if !ø.Is(IdForbidden) && ø.id != "" {
		//res += SingleAttr{"id", string(ø.id)}.String()
		buffer.WriteString(SingleAttr{"id", string(ø.id)}.String())
	}
	if !ø.Is(ClassForbidden) && len(ø.classes) > 0 {
		//res += SingleAttr{"class", ø.classAttrString(ø.classes)}.String()
		buffer.WriteString(SingleAttr{"class", ø.classAttrString(ø.classes)}.String())
	}
	if !ø.Is(Invisible) && len(ø.style) > 0 {
		//res += SingleAttr{"style", ø.styleAttrString(ø.style)}.String()
		buffer.WriteString(SingleAttr{"style", ø.styleAttrString(ø.style)}.String())
	}

	for k, v := range ø.attributes {
		//res += SingleAttr{k, v}.String()
		buffer.WriteString(SingleAttr{k, v}.String())
	}
	return buffer.String()
}

func (ø *Element) InnerHtml() (res string) {
	var buffer bytes.Buffer
	// res = ""
	for _, in := range ø.inner {
		//res += in.String()
		buffer.WriteString(in.String())
	}
	return buffer.String()
}

// returns only children that are Elements, no Text or Html
func (ø *Element) Children() (c []*Element) {
	c = []*Element{}
	if len(ø.inner) == 0 {
		return
	}
	for _, in := range ø.inner {
		switch t := in.(type) {
		case *Element:
			c = append(c, t)
		}
	}
	return
}

// filter by anything that fullfills the matcher interface,
// e.g. Class, Id, Attr, Attrs, Css, Tag, Style, Styles
// recursive finds all tags from the children
func (ø *Element) All(m Matcher) (r []*Element) {
	r = []*Element{}
	if len(ø.inner) == 0 {
		return
	}
	for _, in := range ø.inner {
		switch t := in.(type) {
		case *Element:
			if m.Matches(t) {
				r = append(r, t)
			}
			innerFound := t.All(m)
			for _, innerT := range innerFound {
				r = append(r, innerT)
			}
		}
	}
	return
}

// filter by anything that fullfills the matcher interface,
// e.g. Class, Id, Attr, Attrs, Css, Tag, Style, Styles
// returns the first tag in the children and the subchildren that matches
func (ø *Element) Any(m Matcher) (r *Element) {
	if len(ø.inner) == 0 {
		return nil
	}
	for _, in := range ø.inner {
		switch t := in.(type) {
		case *Element:
			if m.Matches(t) {
				r = t
				return
			}
			r = t.Any(m)
			if r != nil {
				return
			}
		}
	}
	return nil
}
