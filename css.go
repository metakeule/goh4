package goh4

import (
	"fmt"
	"strings"
)

type Styler interface {
	StyleCmd() string
}

type Import string

func (ø Import) String() string {
	return fmt.Sprintf("@import %#v;\n", string(ø))
}

type rule struct {
	Comment   string
	Selectors []Selecter
	Styles    []Styler
	nested    []*rule
}

func Rule(xs ...interface{}) (r *rule, err error) {
	r = &rule{}
	for _, x := range xs {
		switch v := x.(type) {
		case Selecter:
			r.Selectors = append(r.Selectors, v)
		case *rule:
			r.nested = append(r.nested, v)
		case Styler:
			r.Styles = append(r.Styles, v)
		case string:
			r.Comment = v
		default:
			err = fmt.Errorf("%T is no stringer", x)
			return
		}
	}
	return
}

func (ø *rule) String() string {
	styles := []string{}
	for _, st := range ø.Styles {
		styles = append(styles, st.StyleCmd()+";")
	}
	comment := ""
	if ø.Comment != "" {
		comment = fmt.Sprintf("/* %s */\n", ø.Comment)
	}
	strs := []string{}
	for _, s := range ø.Selectors {
		strs = append(strs, s.Selector())
	}
	nested := []string{}
	for _, nr := range ø.nested {
		ns := nr.String()
		for _, nsi := range strings.Split(ns, "\n") {
			nested = append(nested, nsi)
		}
	}
	return fmt.Sprintf("%s%s {\n\t%s\n\t%s\n}", comment, strings.Join(strs, ",\n"), strings.Join(styles, "\n\t"), strings.Join(nested, "\n\t"))
}

// adds given styles
func (ø *rule) Style(styles ...Style) *rule {
	for _, st := range styles {
		ø.Styles = append(ø.Styles, st)
	}
	return ø
}

func (ø *rule) Nest(xs ...interface{}) (i *rule, err error) {
	i, err = Rule(xs...)
	if err != nil {
		return
	}
	ø.nested = append(ø.nested, i)
	return
}

// returns a copy
func (ø *rule) Copy() (newrule *rule) {
	newStyles := []Styler{}
	for _, st := range ø.Styles {
		newStyles = append(newStyles, st)
	}
	newSelectors := []Selecter{}
	for _, st := range ø.Selectors {
		newSelectors = append(newSelectors, st)
	}
	newrule = &rule{
		Comment:   ø.Comment,
		Styles:    newStyles,
		Selectors: newSelectors,
	}
	return
}

// returns a copy that is embedded in the selector
func (ø *rule) Embed(selector Selecter) (newrule *rule) {
	newrule = ø.Copy()
	newSelectors := []Selecter{}
	for _, s := range ø.Selectors {
		newSelectors = append(newSelectors, SelectorString(selector.Selector()+" "+s.Selector()))
	}
	newrule.Selectors = newSelectors
	return
}

// returns a copy that is a composition of this rule with the styles
// of other rules
func (ø *rule) Compose(parents ...*rule) (newrule *rule) {
	newrule = ø.Copy()
	for _, parent := range parents {
		for _, st := range parent.Styles {
			newrule.Styles = append(newrule.Styles, st)
		}
	}
	return
}

//type Css []*rule
type Css []Stringer

// returns a copy with all rules embedded with selector
/*
func (ø Css) Embed(selector Selecter) (newCss Css) {
	newCss = []*rule{}
	for _, r := range ø {
		newCss = append(newCss, r.Embed(selector))
	}
	return
}
*/

func (ø Css) String() string {
	rules := []string{}

	for _, r := range ø {
		rules = append(rules, r.String())
	}

	return strings.Join(rules, "\n\n")
}
