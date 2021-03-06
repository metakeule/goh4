package scss

import (
	"fmt"
	"gopkg.in/metakeule/goh4.v5"
	"gopkg.in/metakeule/goh4.v5/css"
	"io/ioutil"
	"os"
	"strings"
)

type var_ struct {
	Name_ string
	Value string
}

var Parent = goh4.SelectorString("&")

func Var(name string) var_           { return var_{Name_: name} }
func (ø var_) Val(value string) var_ { return var_{ø.Name_, value} }
func (ø var_) Name() string          { return "$" + ø.Name_ }
func (ø var_) Css() string           { return ø.Style() }
func (ø var_) Style() string         { return fmt.Sprintf("%s: %s;", ø.Name(), ø.Value) }
func (ø var_) String() string        { return ø.Style() }
func (ø var_) Selector() string      { return fmt.Sprintf("#{%s}", ø.String()) }

func CompileClasses(classes []goh4.Class, file string) error {
	css := []string{}
	for _, c := range classes {
		name := string(c)
		name = strings.Replace(name, "-", "_", -1)
		css = append(css, Var("class_"+name).Val(c.Selector()).Css())
	}
	return ioutil.WriteFile(file, []byte(strings.Join(css, "\n")), os.FileMode(0644))
}

func CompileIds(ids []goh4.Id, file string) error {
	css := []string{}
	for _, c := range ids {
		name := string(c)
		name = strings.Replace(name, "-", "_", -1)
		css = append(css, Var("id_"+name).Val(c.Selector()).Css())
	}
	return ioutil.WriteFile(file, []byte(strings.Join(css, "\n")), os.FileMode(0644))
}

type param struct {
	Var      var_
	default_ string
}

func Param(name string) param { return param{Var: var_{Name_: name}} }

func (ø param) Default(s string) param {
	return param{var_{Name_: ø.Var.Name_, Value: ø.Var.Value}, s}
}

func (ø param) String() (s string) {
	if ø.default_ == "" {
		s = fmt.Sprintf("%s", ø.Var.Name())
	} else {
		s = fmt.Sprintf("%s: %s", ø.Var.Name(), ø.default_)
	}
	return
}

func (ø param) Name() (s string) { return ø.Var.Name() }

type mixin struct {
	*css.RuleStruct
	Name string
	Args []string
}

func Mixin(name string, params ...param) (ø *mixin) {
	args := []string{}
	for _, param := range params {
		args = append(args, param.String())
	}
	ø = &mixin{&css.RuleStruct{}, name, args}
	return
}

func (ø *mixin) String() string {
	ø.RuleStruct.Selector = ø
	return ø.RuleStruct.String()
}

func (ø *mixin) argString() (s string) { return strings.Join(ø.Args, ", ") }

func (ø *mixin) Selector() (s string) {
	if len(ø.Args) == 0 {
		s = "@mixin " + ø.Name
	} else {
		s = fmt.Sprintf("@mixin %s(%s)", ø.Name, ø.argString())
	}
	return
}

type call mixin
type include mixin

func Call(fn string, args ...string) (ø *call) {
	ø = &call{}
	ø.Name = fn
	ø.Args = args
	return
}

func (ø *call) Style() (s string) {
	return ø.String() + ";"
}

func (ø *call) String() (s string) {
	m := mixin(*ø)
	s = fmt.Sprintf("%s(%s)", m.Name, m.argString())
	return
}

func Include(mixin string, args ...string) (ø *include) {
	ø = &include{}
	ø.Name = mixin
	ø.Args = args
	return
}

func (ø *include) Style() (s string) {
	m := mixin(*ø)
	if len(m.Args) == 0 {
		s = fmt.Sprintf("@include %s;", m.Name)
	} else {
		s = fmt.Sprintf("@include %s(%s);", m.Name, m.argString())
	}
	return
}
