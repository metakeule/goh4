package goh4

type Template struct {
	*Element
	css              []Stringer
	locals           map[string]Stringer
	placeholderCache map[string]*Element
}

func NewTemplate(t *Element) *Template {
	return &Template{
		Element:          t,
		locals:           map[string]Stringer{},
		placeholderCache: map[string]*Element{},
		css:              []Stringer{},
	}
}

// merges the locals to the Templates
func (ø *Template) Merge() {
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

// create a common Css, no context given
func (ø *Template) NewCss(objects ...Stringer) (r *Css) {
	r = NewCss(objects...)
	ø.Add(r)
	return
}

// caches the Stringer and panics if it can't be found
func (ø *Template) cacheFragement(id string) {
	found, h := ø.Element.Any(Id(id))
	if !found {
		panicf("tag with id %v not found in %s", id, ø.Element.String())
	}
	ø.placeholderCache[id] = h
}

// overwrite Element.String()
func (ø *Template) String() string {
	if len(ø.css) > 0 && ø.Element.tag == "doc" {
		sElement := NewElement(Tag("style"), Invisible, WithoutEscaping)
		for _, cr := range ø.css {
			sElement.Add(cr)
		}
		found, head := ø.Element.Any(Tag("head"))
		if !found {
			panicf("no head tag present in %s", ø.Element.Path())
		}
		head.Add(sElement)
	}
	return ø.Element.String()
}

func (ø *Template) Assign(id string, html Stringer) {
	if ø.placeholderCache[id] == nil {
		ø.cacheFragement(id)
	}
	ø.locals[id] = html
	ø.Merge()
}

func (ø *Template) AddCss(rule Stringer) {
	if ø.Element.tag != "doc" {
		panicf("can't add a Css to tag %s", ø.Element.Tag())
	}
	ø.css = append(ø.css, rule)
}
