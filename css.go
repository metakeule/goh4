package goh4

import (
	"fmt"
	"strings"
)

// interface for a parent of a Css -
// is implemented by Css but allows
// your own implementation and interception
// of method calls
type CssParenter interface {
	Styles() map[string]string
	Comment() string
	Selector() string
}

// Describes a CSS Statement, consisting of
// a selector and the styles. Most of the time
// you'll want to give it a Class().
type Css struct {
	class   Class
	Context Context
	comment Comment
	Tags    tags
	styles  map[string]string
	parents []CssParenter
}

// creates a new Css, based on Stringer objects.
// the following types are handled in a special way
//
// 	- Context: sets the context
// 	- Tag: adds a single tag
// 	- Tags: adds multiple tags
// 	- Style: adds a single style
// 	- Styles: adds multiple styles
// 	- Class: sets the class
// 	- Comment: sets the comment
//
// the rest is casted to CssParenter and on success
// the Css inherits from that
func NewCss(objects ...Stringer) (c *Css) {
	c = &Css{
		Tags:    tags{},
		styles:  map[string]string{},
		parents: []CssParenter{},
	}
	for _, o := range objects {
		switch v := o.(type) {
		case Context:
			c.Context = v
		case Tag:
			c.Tags = append(c.Tags, v)
		case tags:
			for _, tg := range v {
				c.Tags = append(c.Tags, tg)
			}
		case styles:
			for _, styl := range v {
				c.Set(styl.Key, styl.Value)
			}
		case style:
			c.Set(v.Key, v.Value)
		case Class:
			c.class = v
		case Comment:
			c.comment = v
		default:
			parenter, ok := v.(CssParenter)
			if ok {
				c.InheritFrom(parenter)
			}
		}
	}
	return
}

func (ø *Css) Class() string {
	return ø.class.String()
}

func (ø *Css) SetClass(c Class) {
	ø.class = c
}

func (ø *Css) Comment() string {
	return ø.comment.String()
}

func (ø *Css) matchTag(t *Element, tt Tag) (r bool) {
	r = tt == t.tag
	return
}

func (ø *Css) partialSelector(t string) string {
	sel := ""
	if t != "" {
		sel += t
	}
	if ø.class != "" {
		sel += "." + ø.class.String()
	}

	if ø.Context != "" {
		sel = ø.Context.String() + " " + sel
	}
	return sel
}

// constructs and returns the selector
func (ø *Css) Selector() (s string) {
	if len(ø.Tags) > 0 {
		selectors := []string{}

		for _, tt := range ø.Tags {
			selectors = append(selectors, ø.partialSelector(tt.String()))
		}
		s = strings.Join(selectors, ",\n")

	} else {
		s = ø.partialSelector("")
	}
	return
}

func (ø *Css) styleString() (r string) {
	r = ""
	consolidated := map[string]string{}
	inherited := map[string]CssParenter{}
	for _, p := range ø.parents {
		for k, v := range p.Styles() {
			inherited[k] = p
			consolidated[k] = v
		}
	}
	for k, v := range ø.styles {
		if inherited[k] != nil {
			delete(inherited, k)
		}
		consolidated[k] = v
	}

	for k, v := range consolidated {
		comment := ""
		if p := inherited[k]; p != nil {
			cmt := ""
			if p.Comment() != "" {
				cmt = `: "` + p.Comment() + `"`
			}
			comment = fmt.Sprintf("\t/* inherited from »%s«%s */", p.Selector(), cmt)
		}
		r += fmt.Sprintf("\t%s%s\n", style{k, v}, comment)
	}

	return
}

// return the Css as stylesheet string
func (ø *Css) String() string {
	descr := ""
	if ø.Comment() != "" {
		descr += fmt.Sprintf("\n/* %s */\n", ø.Comment())
	}
	return fmt.Sprintf("\n%s%s {\n%s}\n", descr, ø.Selector(), ø.styleString())
}

// set the styles, they are given in pairs, e.g.
//
// 	Set("color","green","width","200")
func (ø *Css) Set(vals ...string) {
	for i := 0; i < len(vals); i = i + 2 {
		ø.styles[vals[i]] = vals[i+1]
	}
}

// returns the styles that are inherited (own overwrites are not respected)
func (ø *Css) InheritedStyles() map[string]string {
	inherited := map[string]string{}
	for _, p := range ø.parents {
		for k, v := range p.Styles() {
			inherited[k] = v
		}
	}
	return inherited
}

// returns the styles that are not inherited
func (ø *Css) OwnStyles() map[string]string {
	return ø.styles
}

// returns the styles with the inherited, overwritten by the own
func (ø *Css) Styles() map[string]string {
	consolidated := map[string]string{}
	for _, p := range ø.parents {
		for k, v := range p.Styles() {
			consolidated[k] = v
		}
	}
	for k, v := range ø.styles {
		consolidated[k] = v
	}
	return consolidated
}

// lets the Css inherit from the given CssParenter
func (ø *Css) InheritFrom(ps ...CssParenter) {
	for _, p := range ps {
		ø.parents = append(ø.parents, p)
	}
}
