package goh4

import (
	"fmt"
	"strings"
)

type Styler interface {
	Style() string
}

type Import string

func (ø Import) String() string {
	return fmt.Sprintf("@import %#v;\n", string(ø))
}

type RuleStruct struct {
	Comment   string
	Selectors []Selecter
	Styles    []Styler
	nested    []*RuleStruct
}

func Rule(xs ...interface{}) (r *RuleStruct, err error) {
	r = &RuleStruct{}
	for _, x := range xs {
		switch v := x.(type) {
		case Selecter:
			r.Selectors = append(r.Selectors, v)
		case *RuleStruct:
			r.nested = append(r.nested, v)
		case Styler:
			r.Styles = append(r.Styles, v)
		case []Style:
			r.Style(v...)
		case string:
			r.Comment = v
		default:
			err = fmt.Errorf("%T is no stringer", x)
			return
		}
	}
	return
}

func (ø *RuleStruct) String() string {
	styles := []string{}
	for _, st := range ø.Styles {
		styles = append(styles, st.Style()+";")
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
func (ø *RuleStruct) Style(styles ...Style) *RuleStruct {
	for _, st := range styles {
		ø.Styles = append(ø.Styles, st)
	}
	return ø
}

func (ø *RuleStruct) Nest(xs ...interface{}) (i *RuleStruct, err error) {
	i, err = Rule(xs...)
	if err != nil {
		return
	}
	ø.nested = append(ø.nested, i)
	return
}

// returns a copy
func (ø *RuleStruct) Copy() (newrule *RuleStruct) {
	newStyles := []Styler{}
	for _, st := range ø.Styles {
		newStyles = append(newStyles, st)
	}
	newSelectors := []Selecter{}
	for _, st := range ø.Selectors {
		newSelectors = append(newSelectors, st)
	}
	newrule = &RuleStruct{
		Comment:   ø.Comment,
		Styles:    newStyles,
		Selectors: newSelectors,
	}
	return
}

// returns a copy that is embedded in the selector
func (ø *RuleStruct) Embed(selector Selecter) (newrule *RuleStruct) {
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
func (ø *RuleStruct) Compose(parents ...*RuleStruct) (newrule *RuleStruct) {
	newrule = ø.Copy()
	for _, parent := range parents {
		for _, st := range parent.Styles {
			newrule.Styles = append(newrule.Styles, st)
		}
	}
	return
}

type rules struct {
	rules []*RuleStruct
}

func Rules() *rules {
	r := []*RuleStruct{}
	return &rules{r}
}

func (ø *rules) Add(r *RuleStruct) {
	ø.rules = append(ø.rules, r)
}

func (ø *rules) New(xs ...interface{}) (r *RuleStruct, err error) {
	r, err = Rule(xs...)
	if err != nil {
		return
	}
	ø.Add(r)
	return
}

func (ø *rules) String() string {
	rules := []string{}

	for _, r := range ø.rules {
		rules = append(rules, r.String())
	}

	return strings.Join(rules, "\n\n")
}

// returns a copy with all rules embedded with selector
func (ø *rules) Embed(selector Selecter) *rules {
	newRules := []*RuleStruct{}
	for _, r := range ø.rules {
		newRules = append(newRules, r.Embed(selector))
	}
	return &rules{newRules}
}

type Css []Stringer

func (ø Css) String() string {
	rules := []string{}

	for _, r := range ø {
		rules = append(rules, r.String())
	}

	return strings.Join(rules, "\n\n")
}
