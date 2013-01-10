package goh4

import (
	"fmt"
	"html"
	"strings"
)

type option int

const (
	IdForbidden       option = iota // tag should not have an id attribute
	ClassForbidden                  // tag should not have a class attribute
	SelfClosing                     // tag is selfclosing and contains no content
	Inline                          // tag is an inline element (only for visible tags)
	Field                           // tag is a field of a form
	Invisible                       // tag doesn't render anything visible
	WithoutEscaping                 // for content that is not escaped
	WithoutDecoration               // for tags like doc that won't show up as tags, but show their content
)

// allows the Parent of an Element to be exchanged
type Pather interface {
	Path() string
}
type Tager interface {
	Tag() string
}
type Csser interface {
	Matcher
	Class() string
}

type Elementer interface {
	Stringer
	Tager
	IsParentAllowed(Tager) bool
	SetParent(Pather)
}

type Element struct {
	Attributes        map[string]string
	Comment           Comment
	parent            Pather
	classes           []Class
	id                Id
	inner             []Stringer
	parentTags        []Tag
	style             map[string]string
	inline            bool
	field             bool
	withoutEscaping   bool
	withoutDecoration bool
	tag               Tag
	selfclosing       bool
	idForbidden       bool
	classForbidden    bool
	invisible         bool
}

// contruct a new tag with some options
// use this method inside a func that adds as helper to generate a certain tag, e.g.
// func A() *Element { return MakeElement("a", Inline) }
func NewElement(tag Tag, options ...option) (t *Element) {
	t = &Element{
		Attributes: map[string]string{},
		tag:        tag,
		style:      map[string]string{},
		parentTags: []Tag{},
	}
	for _, o := range options {
		switch o {
		case IdForbidden:
			t.idForbidden = true
		case ClassForbidden:
			t.classForbidden = true
		case SelfClosing:
			t.selfclosing = true
		case Inline:
			t.inline = true
		case Field:
			t.field = true
		case Invisible:
			t.invisible = true
		case WithoutEscaping:
			t.withoutEscaping = true
		case WithoutDecoration:
			t.withoutDecoration = true
		}
	}
	return
}

// returns the fields of a form,
// panics if invoked on something other than a form
func (ø *Element) Fields() (fields []*Element) {
	if ø.tag != "form" {
		panicf("%s is no form", ø.tag)
	}
	return ø.All(FieldMatcher(0))
}

func (ø *Element) Parent() Pather {
	return ø.parent
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

// create a css with context set to Path()
func (ø *Element) NewCss(objects ...Stringer) *Css {
	objects = append(objects, Context(ø.Path()))
	return NewCss(objects...)
}

func (ø *Element) idPath() (s string) {
	s = ""
	if !ø.idForbidden && ø.id != "" {
		s = "#" + ø.id.String()
	}
	return
}

// return false if the given parent tag is not allowed
func (ø *Element) IsParentAllowed(parent Tager) (allowed bool) {
	if len(ø.parentTags) == 0 {
		return true
	}
	allowed = false
	for _, p := range ø.parentTags {
		if p.String() == parent.Tag() {
			allowed = true
		}
	}
	return
}

// adds css properties to the style attribute, same keys are overwritten
func (ø *Element) addStyle(styles ...interface{}) {
	if ø.invisible {
		panicf("can't set style for invisible tag %s", ø.tag)
	}
	for _, s := range styles {
		switch v := s.(type) {
		case Style:
			ø.style[v.Key] = v.Value
		case Styles:
			st := v.ToStyleArr()
			for _, style := range st {
				ø.style[style.Key] = style.Value
			}
		case string:
			a := strings.Split(v, ":")
			ø.style[a[0]] = strings.Replace(a[1], ";", "", 1)
		default:
			panicf("unsupported style: %#v", v)
		}
	}
}

// adds attributes, same keys are overwritten
func (ø *Element) addAttr(attrs ...interface{}) {
	for _, a := range attrs {
		switch v := a.(type) {
		case Attr:
			ø.Attributes[v.Key] = v.Value
		case Attrs:
			st := v.ToAttrArr()
			for _, attr := range st {
				ø.Attributes[attr.Key] = attr.Value
			}
		case string:
			ar := strings.Split(v, "=")
			ø.Attributes[ar[0]] = strings.Replace(ar[1], `"`, ``, 2)
		default:
			panicf("unsupported attribute: %#v", v)
		}
	}
}

// apply the css to the element
func (ø *Element) ApplyCss(rules ...Csser) {
	for _, r := range rules {
		if r.Class() == "" {
			panicf("can't apply css rule without a class: %s", r)
		}

		ø.AddClass(Class(r.Class()))

		if !r.Matches(ø) {
			panicf("tag %s not affected by css rule %#v", ø.Path(), r)
		}
	}
}

func (ø *Element) ensureContentAddIsAllowed() {
	if ø.selfclosing {
		panicf("add not allowed for tag »%s«", ø.tag)
	}
}

func (ø *Element) AddAtPosition(pos int, v Stringer) {
	ø.ensureParentIsSetAndContentIsAllowed(v)
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

}

func (ø *Element) ensureParentIsSetAndContentIsAllowed(v Stringer) {
	e, isElementer := v.(Elementer)
	if isElementer {
		if !e.IsParentAllowed(ø) {
			panicf("tag %s not allowed as parent of %s", ø.Tag(), e.Tag())
		}
		e.SetParent(ø)
	}
	ø.ensureContentAddIsAllowed()
}

func (ø *Element) SetAtPosition(pos int, v Stringer) {
	ø.ensureParentIsSetAndContentIsAllowed(v)
	ø.inner[pos] = v
}

func (ø *Element) SetBottom(v Stringer) {
	ø.ensureParentIsSetAndContentIsAllowed(v)
	ø.inner[len(ø.inner)-1] = v
}

type PositionMatcher struct {
	Element *Element
	Pos     int
}

func (ø *PositionMatcher) Matches(e *Element) (f bool) {
	// no recursive findings
	if e.Parent() != ø.Element.Parent() {
		return
	}

	ø.Pos += 1
	if ø.Element == e {
		return true
	}
	return
}

// returns -1 if element could not be found
func (ø *Element) PositionOf(v *Element) (pos int, found bool) {
	m := &PositionMatcher{v, -1}
	_, _ = ø.Any(m)
	pos = m.Pos
	if pos != -1 {
		found = true
	}
	return
}

func (ø *Element) AddBefore(v *Element, nu *Element) {
	pos, _ := ø.PositionOf(v)
	if pos == -1 {
		panicf("could not find %#v", v)
	}
	ø.ensureParentIsSetAndContentIsAllowed(nu)
	ø.AddAtPosition(pos, nu)
}

func (ø *Element) AddAfter(v *Element, nu *Element) {
	pos, _ := ø.PositionOf(v)
	if pos == -1 {
		panicf("could not find %#v", v)
	}
	ø.ensureParentIsSetAndContentIsAllowed(nu)
	if pos > len(ø.inner)-2 {
		ø.Add(nu)
	} else {
		ø.AddAtPosition(pos+1, nu)
	}
}

// add new inner content to a tag
// objects may be Text, Html, *Element, Attr, Attrs, Style, Styles, Css, Id, Class
func (ø *Element) Add(objects ...Stringer) {
	for _, o := range objects {
		switch v := o.(type) {
		case Text:
			ø.ensureContentAddIsAllowed()
			if !ø.withoutEscaping {
				v = Text(html.EscapeString(string(v)))
			}
			ø.inner = append(ø.inner, v)
		case Html:
			ø.ensureContentAddIsAllowed()
			ø.inner = append(ø.inner, v)
		case Comment:
			ø.Comment = v
		case Attr:
			ø.addAttr(v)
		case Attrs:
			ø.addAttr(v)
		case Class:
			ø.AddClass(v)
		case Id:
			ø.SetId(v)
		case Style:
			ø.addStyle(v)
		case Styles:
			ø.addStyle(v)
		case *Css:
			ø.ApplyCss(v)
		default:
			ø.ensureParentIsSetAndContentIsAllowed(v)
			ø.inner = append(ø.inner, v)
		}
	}
}

// sets the inner content of a tag
// objects may be Text, Html, *Element, Attr, Attrs, Style, Styles, Css, Id, Class
func (ø *Element) Set(objects ...Stringer) {
	ø.inner = []Stringer{}
	ø.classes = []Class{}
	ø.Add(objects...)
}

// use this func to set the id of the tag,
// do not set it via Attr directly
func (ø *Element) SetId(id Id) {
	if ø.idForbidden {
		panicf("id not allowed for tag %s", ø.tag)
	}
	ø.id = id
}

func (ø *Element) classAttrString(classes []Class) (s string) {
	s = ""
	for _, cl := range classes {
		s += cl.String() + " "
	}
	return
}

func (ø *Element) classPath(classes []Class) (s string) {
	s = ""
	for _, cl := range classes {
		s += "." + cl.String()
	}
	return
}

func (ø *Element) styleAttrString(styles map[string]string) (s string) {
	s = ""
	for k, v := range styles {
		s += Style{k, v}.String()
	}
	return
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

// use this func to set the classes of the tag,
// do not set it via Attr directly
func (ø *Element) SetClass(classes ...Class) {
	ø.classes = []Class{}
	ø.AddClass(classes...)
}

// use this func to add the classes of the tag,
// do not set it via Attr directly
func (ø *Element) AddClass(classes ...Class) {
	if ø.classForbidden {
		panicf("class not allowed for tag %s", ø.tag)
	}
	for _, cl := range classes {
		ø.classes = append(ø.classes, cl)
	}
}

// returns the classes
func (ø *Element) Classes() (c []Class) {
	return ø.classes
}
func (ø *Element) Id() Id { return ø.id }

// returns the html with inner content
func (ø *Element) String() (res string) {
	if ø.withoutDecoration {
		return ø.InnerHtml()
	}
	commentpre := ""
	commentpost := ""
	if ø.Comment != "" {
		commentpre = fmt.Sprintf("<!-- Begin: %s -->", ø.Comment)
		commentpost = fmt.Sprintf("<!-- End: %s -->", ø.Comment)
	}
	if ø.selfclosing {
		res = fmt.Sprintf("%s<%s%s />%s", commentpre, string(ø.tag), ø.attrsString(), commentpost)
	} else {
		str := "%s<%s%s>%s</%s>%s"
		if !ø.inline {
			str = "\n%s<%s%s>%s</%s>%s\n"
		}
		res = fmt.Sprintf(str, commentpre, string(ø.tag), ø.attrsString(), ø.InnerHtml(), string(ø.tag), commentpost)
	}
	return
}

// prepare the id attribute for output
func (ø *Element) attrsString() (res string) {
	res = ""
	if !ø.idForbidden && ø.id != "" {
		res += Attr{"id", string(ø.id)}.String()
	}
	if !ø.classForbidden && len(ø.classes) > 0 {
		res += Attr{"class", ø.classAttrString(ø.classes)}.String()
	}
	if !ø.invisible && len(ø.style) > 0 {
		res += Attr{"style", ø.styleAttrString(ø.style)}.String()
	}

	for k, v := range ø.Attributes {
		res += Attr{k, v}.String()
	}
	return
}

func (ø *Element) InnerHtml() (res string) {
	res = ""
	for _, in := range ø.inner {
		res += in.String()
	}
	return
}

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
func (ø *Element) Any(m Matcher) (found bool, r *Element) {
	if len(ø.inner) == 0 {
		return
	}
	for _, in := range ø.inner {
		switch t := in.(type) {
		case *Element:
			if m.Matches(t) {
				r = t
				found = true
				return
			}
			found, r = t.Any(m)
			if found {
				return
			}
		}
	}
	return
}
