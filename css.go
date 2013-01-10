package goh4

import (
	"fmt"
	"strings"
)

// interface for a parent of a Css
// is implemented by Css but allows
// your own implementation and interception
// of method calls
type CssParenter interface {
	Styles() map[string]string
	Comment() string
	Selector() string
}

type Css struct {
	class   Class
	Context Context
	comment Comment
	Tags    Tags
	styles  map[string]string
	parents []CssParenter
}

func NewCss(objects ...Stringer) (c *Css) {
	c = &Css{
		Tags:    Tags{},
		styles:  map[string]string{},
		parents: []CssParenter{},
	}
	for _, o := range objects {
		switch v := o.(type) {
		case Context:
			c.Context = v
		case Tag:
			c.Tags = append(c.Tags, v.String())
		case Tags:
			for _, tag := range v {
				c.Tags = append(c.Tags, tag)
			}
		case Styles:
			styles := v.ToStyleArr()
			for _, style := range styles {
				c.Set(style.Key, style.Value)
			}
		case Style:
			c.Set(v.Key, v.Value)
		case Class:
			c.class = v
		case Comment:
			c.comment = v
		default:
			parenter, ok := v.(CssParenter)
			if !ok {
				panicf("unknown type for css construction: %#v", v)
			}
			c.InheritFrom(parenter)
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

func (ø *Css) Matches(t *Element) (m bool) {
	if ø.Context != "" {
		panicf("not supported for matching: css rule contains context %s", ø.Context)
	}
	if ø.class != "" {
		if !t.HasClass(ø.class) {
			return false
		}
	}

	if len(ø.Tags) > 0 {
		for _, tt := range ø.Tags {
			if ø.matchTag(t, Tag(tt)) {
				return true
			}
		}
	} else {
		return true
	}

	return
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

func (ø *Css) Selector() (s string) {
	if len(ø.Tags) > 0 {
		selectors := []string{}

		for _, tt := range ø.Tags {
			selectors = append(selectors, ø.partialSelector(tt))
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
	for k, v := range ø.Styles() {
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
		r += fmt.Sprintf("\t%s%s\n", Style{k, v}, comment)
	}

	return
}

func (ø *Css) String() string {
	descr := ""
	if ø.Comment() != "" {
		descr += fmt.Sprintf("\n\n/* %s */\n", ø.Comment())
	}
	return fmt.Sprintf("\n%s%s {\n%s}\n", descr, ø.Selector(), ø.styleString())
}

// set the styles
func (ø *Css) Set(vals ...string) {
	for i := 0; i < len(vals); i = i + 2 {
		ø.styles[vals[i]] = vals[i+1]
	}
}

func (ø *Css) Styles() map[string]string {
	return ø.styles
}

// let the rule inherit from the given
func (ø *Css) InheritFrom(r CssParenter) {
	ø.parents = append(ø.parents, r)
}
