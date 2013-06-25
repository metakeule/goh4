package css

import (
	"fmt"
	. "github.com/metakeule/goh4"
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
	Comment  string
	Selector Selecter
	Styles   []Styler
	nested   []*RuleStruct
}

func Rule(xs ...interface{}) (r *RuleStruct, err error) {
	r = &RuleStruct{}
	for _, x := range xs {
		switch v := x.(type) {
		// we don't want to handle *RuleStruct and *rules here,
		// since there would be different ways to handle them (Next, Embed, Compose )
		// and we don't want to have a implicit default way
		case Selecter:
			r.Selector = v
		case Styler:
			r.Styles = append(r.Styles, v)
		case []Style:
			r.Style(v...)
		case []Styler:
			r.Styles = append(r.Styles, v...)
		case string:
			r.Comment = v
		case Comment:
			r.Comment = string(v)
		default:
			err = fmt.Errorf("%T is not an allowed type", x)
			return
		}
	}
	return
}

// for each selector, my selectors is prefixed and
// my rules are applied
func (ø *RuleStruct) ForEach(c SelecterAdder, sel ...Selecter) (*RuleStruct, error) {
	comb := c.Add(ø.Selector)
	all := Each()
	for _, s := range sel {
		all.Add(comb.Add(s))
	}
	return Rule(all, ø.Styles)
}

func (ø *RuleStruct) String() string {
	styles := []string{}
	for _, st := range ø.Styles {
		styles = append(styles, st.Style())
	}
	comment := ""
	if ø.Comment != "" {
		comment = fmt.Sprintf("/* %s */\n", ø.Comment)
	}
	strs := []string{}
	strs = append(strs, ø.Selector.Selector())
	nested := []string{}
	for _, nr := range ø.nested {
		ns := nr.String()
		nssp := strings.Split(ns, "\n")
		for _, nsi := range nssp {
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

func (ø *RuleStruct) _inner(sa SelecterAdder, sel Selecter, xs ...interface{}) (i *RuleStruct, err error) {
	i, err = Rule(xs...)
	if err != nil {
		return
	}
	i.Selector = sa.Add(ø.Selector).Add(sel)
	return
}

func (ø *RuleStruct) Descendant(sel Selecter, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Descendant(), sel, xs...)
}

func (ø *RuleStruct) Child(sel Selecter, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Child(), sel, xs...)
}

func (ø *RuleStruct) DirectFollows(sel Selecter, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(DirectFollows(), sel, xs...)
}

func (ø *RuleStruct) Follows(sel Selecter, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Follows(), sel, xs...)
}

func (ø *RuleStruct) Each(sel Selecter, xs ...interface{}) (i *RuleStruct, err error) {
	return ø._inner(Each(), sel, xs...)
}

// returns a copy
func (ø *RuleStruct) Copy() (newrule *RuleStruct) {
	newStyles := []Styler{}
	for _, st := range ø.Styles {
		newStyles = append(newStyles, st)
	}
	newrule = &RuleStruct{
		Comment:  ø.Comment,
		Styles:   newStyles,
		Selector: ø.Selector,
	}
	return
}

// returns a copy that is embedded in the selector
func (ø *RuleStruct) Embed(selector Selecter) (newrule *RuleStruct) {
	newrule = ø.Copy()
	newSelector := SelectorString(selector.Selector() + " " + ø.Selector.Selector())
	newrule.Selector = newSelector
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
