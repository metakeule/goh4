package goh4

import (
	"fmt"
	"strings"
)

type scssVar struct {
	Name_ string
	Value string
}

func ScssVar(name string) scssVar {
	return scssVar{Name_: name}
}

func (ø scssVar) Val(value string) scssVar {
	return scssVar{ø.Name_, value}
}

func (ø scssVar) Name() string {
	return "$" + ø.Name_
}

func (ø scssVar) StyleCmd() string {
	return fmt.Sprintf("%s: %s", ø.Name(), ø.Value)
}

func (ø scssVar) String() string {
	return ø.StyleCmd() + ";"
}

func (ø scssVar) Selector() string {
	return fmt.Sprintf("#{%s}", ø.String())
}

type scssParam struct {
	Var      scssVar
	default_ string
}

func ScssParam(name string) scssParam {
	return scssParam{Var: scssVar{Name_: name}}
}

func (ø scssParam) Default(s string) scssParam {
	return scssParam{scssVar{Name_: ø.Var.Name_, Value: ø.Var.Value}, s}
}

func (ø scssParam) String() (s string) {
	if ø.default_ == "" {
		s = fmt.Sprintf("%s", ø.Var.Name())
	} else {
		s = fmt.Sprintf("%s: %s", ø.Var.Name(), ø.default_)
	}
	return
}

func (ø scssParam) Name() (s string) {
	return ø.Var.Name()
}

type scssInclude mixin

type scssCall mixin

func ScssCall(fn string, args ...string) (ø *scssCall) {
	ø = &scssCall{}
	ø.Name = fn
	ø.Args = args
	return
}

func (ø *scssCall) StyleCmd() (s string) {
	m := mixin(*ø)
	s = fmt.Sprintf("%s(%s)", m.Name, m.argString())
	return
}

func (ø *scssCall) String() (s string) {
	return ø.StyleCmd()
}

func ScssInclude(mixin string, args ...string) (ø *scssInclude) {
	ø = &scssInclude{}
	ø.Name = mixin
	ø.Args = args
	return
}

func (ø *scssInclude) StyleCmd() (s string) {
	m := mixin(*ø)
	if len(m.Args) == 0 {
		s = fmt.Sprintf("@include %s", m.Name)
	} else {
		s = fmt.Sprintf("@include %s(%s)", m.Name, m.argString())
	}
	return
}

var ScssParent = SelectorString("&")

type mixin struct {
	*rule
	Name string
	Args []string
}

func ScssMixin(name string, params ...scssParam) (ø *mixin) {
	args := []string{}
	for _, param := range params {
		args = append(args, param.String())
	}
	ø = &mixin{&rule{}, name, args}
	return
}

func (ø *mixin) String() string {
	ø.rule.Selectors = []Selecter{ø}
	return ø.rule.String()
}

func (ø *mixin) argString() (s string) {
	return strings.Join(ø.Args, ", ")
}

func (ø *mixin) Selector() (s string) {
	if len(ø.Args) == 0 {
		s = "@mixin " + ø.Name
	} else {
		s = fmt.Sprintf("@mixin %s(%s)", ø.Name, ø.argString())
	}
	return
}
