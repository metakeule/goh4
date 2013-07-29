package goh4

import (
	"fmt"
	// "github.com/metakeule/template"
	"html"
	"regexp"
	"strings"
)

// type Context string

// func (ø Context) String() string { return string(ø) }

type Comment string

func (ø Comment) String() string { return string(ø) }

type Class string

func (ø Class) String() string { return string(ø) }

func (ø Class) Selector() string { return fmt.Sprintf(".%s", ø) }

func Classf(format string, i ...interface{}) Class {
	return Class(fmt.Sprintf(format, i...))
}

type classes []Class

func Classes(c ...string) classes {
	cl := classes{}
	for _, cc := range c {
		cl = append(cl, Class(cc))
	}
	return cl
}

/*
func (ø Class) Rule() (r *RuleStruct) {
    return &RuleStruct{Selector: ø}
}
*/

type Id string

func (ø Id) String() string { return string(ø) }

func (ø Id) Selector() string { return fmt.Sprintf("#%s", ø) }

/*
func (ø Id) Rule() (r *RuleStruct) {
    return &RuleStruct{Selector: ø}
}
*/

type Text string

func (ø Text) String() string { return string(ø) }

func Textf(format string, i ...interface{}) Text {
	return Text(fmt.Sprintf(format, i...))
}

type Html string

func (ø Html) String() string { return string(ø) }

func Htmlf(format string, i ...interface{}) Html {
	return Html(fmt.Sprintf(format, i...))
}

type Tag string

func (ø Tag) String() string { return string(ø) }

func (ø Tag) Selector() string { return ø.String() }

type SingleAttr struct {
	Key   string
	Value string
}

func (ø SingleAttr) String() string { return " " + ø.Key + `="` + html.EscapeString(ø.Value) + `"` }

type Attrs []SingleAttr

func (ø Attrs) String() (s string) {
	ss := []string{}
	for _, atr := range ø {
		ss = append(ss, atr.String())
	}
	return strings.Join(ss, " ")
}

type tags []Tag

func Tags(tag string, ø ...string) (a tags) {
	a = tags{Tag(tag)}
	for _, s := range ø {
		a = append(a, Tag(s))
	}
	return
}

func (ø tags) String() string {
	str := []string{}
	for _, t := range ø {
		str = append(str, t.String())
	}
	return strings.Join(str, ", ")
}

// helper to easily create multiple SingleAttrs
// use is like this
// Attr("width","200px","height","30px","value","hiho")
func Attr(key1, val1 string, ø ...string) (s Attrs) {
	s = []SingleAttr{SingleAttr{key1, val1}}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, SingleAttr{ø[i], ø[i+1]})
	}
	return
}

/*
type Placeholder string

//func (ø Placeholder) String() string { return "@@" + string(ø) + "@@" }
func (ø Placeholder) Name() string   { return string(ø) }
func (ø Placeholder) String() string { return string(ø) }

func (ø Placeholder) Key() string { return "@@" + string(ø) + "@@" }

//func (ø Placeholder) Html() Html     { return Html(ø.String()) }
*/

type Stringer interface {
	String() string
}

type Selecter interface {
	Selector() string
}

func Context(ctx Selecter, inner1 Selecter, inner ...Selecter) (r Selecter) {
	sel := []string{}
	inner = append(inner, inner1)
	for _, i := range inner {
		sel = append(sel, ctx.Selector()+" "+i.Selector())
	}
	return SelectorString(strings.Join(sel, ",\n"))
}

func ContextString(ctx string, inner1 Selecter, inner ...Selecter) (r Selecter) {
	return Context(SelectorString(ctx), inner1, inner...)
}

type SelecterAdder interface {
	Selector() string
	Add(Selecter) SelecterAdder
}

type SelectorString string

func (ø SelectorString) Selector() string { return string(ø) }

// combine several selectors to one
func Selector(sel1 Selecter, selects ...Selecter) Selecter {
	s := []string{sel1.Selector()}
	for _, sel := range selects {
		s = append(s, sel.Selector())
	}
	return SelectorString(strings.Join(s, ""))
}

/*
func (ø Tag) Rule() (r *RuleStruct) {
	return &RuleStruct{Selector: ø}
}
*/

type Scss string

func (ø Scss) String() string { return string(ø) }

func removeWhiteSpace(in string) string {
	reg := regexp.MustCompile(`\s`)
	return reg.ReplaceAllString(in, "")
}

/*
type style struct {
	Key   string
	Value string
}

func (ø style) String() string { return ø.Key + ": " + ø.Value + ";" }

type styles []style

func (ø styles) String() (s string) {
	ss := []string{}
	for _, st := range ø {
		ss = append(ss, st.String())
	}
	return strings.Join(ss, " ")
}

// helper to easily create multiple styles
// use is like this
// Style("width","200px","height","30px","color","green")
func Styles(ø ...string) (s styles) {
	s = styles{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, style{ø[i], ø[i+1]})
	}
	return
}
*/
