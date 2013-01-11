package goh4

import (
	"fmt"
	"html"
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
	Field                              // element is a field of a form
	Invisible                          // element doesn't render anything visible
	WithoutEscaping                    // for content that is not escaped
	WithoutDecoration                  // for tags like doc that won't show up as tags, but show their content
)

var flagNames = map[flag]string{
	hasDefaults:       "hasDefaults",
	IdForbidden:       "IdForbidden",
	ClassForbidden:    "ClassForbidden",
	SelfClosing:       "SelfClosing",
	Inline:            "Inline",
	Field:             "Field",
	Invisible:         "Invisible",
	WithoutEscaping:   "WithoutEscaping",
	WithoutDecoration: "WithoutDecoration",
}

func (ø flag) String() string {
	return flagNames[ø]
}

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
	Attributes map[string]string
	Comment    Comment
	parent     Pather
	classes    []Class
	id         Id
	inner      []Stringer
	parentTags []Tag
	style      map[string]string
	flags      flag
	tag        Tag
}

// contruct a new element with some flags
// use this method inside a func that adds as helper to generate a certain element, e.g.
// func Input() *Element { return NewElement(Tag("input"), Inline|Field|SelfClosing, Attrs{"name", "myinput"}) }
// or
// func Input() *Element { return NewElement(Tag("input"), Inline,Field,SelfClosing, Attrs{"name", "myinput"}) }
func NewElement(tag Tag, flags ...flag) (ø *Element) {
	ø = &Element{
		Attributes: map[string]string{},
		tag:        tag,
		flags:      hasDefaults,
		style:      map[string]string{},
		parentTags: []Tag{},
	}

	for _, flag := range flags {
		ø.flags = ø.flags | flag
	}
	return
}

func ExampleNewElement() {
	fmt.Println("The output of\nthis example.")
	// Output: The output of
	// this example.
}

func (ø *Element) Is(f flag) bool {
	return ø.flags&f != 0
}

// returns the fields of a form,
func (ø *Element) Fields() (fields []*Element) {
	if ø.tag != "form" {
		return []*Element{}
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
	if !ø.Is(IdForbidden) && ø.id != "" {
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
	if ø.Is(Invisible) {
		// can't set style for invisible tag
		return
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
		}
	}
}

// apply the css to the element
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

func (ø *Element) ensureContentAddIsAllowed() (err error) {
	if ø.Is(SelfClosing) {
		return fmt.Errorf("add not allowed for tag »%s«", ø.tag)
	}
	return
}

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

func (ø *Element) SetBottom(v Elementer) (err error) {
	if err := ø.ensureParentIsSetAndContentIsAllowed(v); err != nil {
		return err
	}
	ø.inner[len(ø.inner)-1] = v
	return
}

// returns -1 if element could not be found
func (ø *Element) PositionOf(v *Element) (pos int, found bool) {
	m := &PositionMatcher{Element: v}
	_, _ = ø.Any(m)
	pos = m.Pos
	found = m.Found
	return
}

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

// add new inner content to a tag
// objects may be Text, Html, *Element, Attr, Attrs, Style, Styles, Css, Id, Class
func (ø *Element) Add(objects ...Stringer) (err error) {
	for _, o := range objects {
		switch v := o.(type) {
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
		case Attr:
			ø.addAttr(v)
		case Attrs:
			ø.addAttr(v)
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
		case Styles:
			ø.addStyle(v)
		case *Css:
			if err := ø.ApplyCss(v); err != nil {
				return err
			}
		default:
			if err := ø.ensureParentIsSetAndContentIsAllowed(v); err != nil {
				return err
			}
			ø.inner = append(ø.inner, v)
		}
	}
	return
}

// sets the inner content of a tag
// objects may be Text, Html, *Element, Attr, Attrs, Style, Styles, Css, Id, Class
func (ø *Element) Set(objects ...Stringer) (err error) {
	ø.inner = []Stringer{}
	ø.classes = []Class{}
	return ø.Add(objects...)
}

// use this func to set the id of the tag,
// do not set it via Attr directly
func (ø *Element) SetId(id Id) (err error) {
	if ø.Is(IdForbidden) {
		return fmt.Errorf("id not allowed for tag %s", ø.tag)
	}
	ø.id = id
	return
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
func (ø *Element) AddClass(classes ...Class) (err error) {
	if ø.Is(ClassForbidden) {
		return fmt.Errorf("class not allowed for tag %s", ø.tag)
	}
	for _, cl := range classes {
		ø.classes = append(ø.classes, cl)
	}
	return
}

// returns the classes
func (ø *Element) Classes() (c []Class) {
	return ø.classes
}
func (ø *Element) Id() Id { return ø.id }

// returns the html with inner content
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
		if !ø.Is(Inline) {
			str = "\n%s<%s%s>%s</%s>%s\n"
		}
		res = fmt.Sprintf(str, commentpre, string(ø.tag), ø.attrsString(), ø.InnerHtml(), string(ø.tag), commentpost)
	}
	return
}

// prepare the id attribute for output
func (ø *Element) attrsString() (res string) {
	res = ""
	if !ø.Is(IdForbidden) && ø.id != "" {
		res += Attr{"id", string(ø.id)}.String()
	}
	if !ø.Is(ClassForbidden) && len(ø.classes) > 0 {
		res += Attr{"class", ø.classAttrString(ø.classes)}.String()
	}
	if !ø.Is(Invisible) && len(ø.style) > 0 {
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
