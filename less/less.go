package less

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

func Var(name string) var_           { return var_{Name_: name} }
func (ø var_) Val(value string) var_ { return var_{ø.Name_, value} }
func (ø var_) Name() string          { return "@" + ø.Name_ }
func (ø var_) Css() string           { return fmt.Sprintf("%s: %s;", ø.Name(), ø.Value) }
func (ø var_) String() string        { return ø.Name() }

func CompileClasses(classes []goh4.Class, file string, info string) error {
	css := []string{"/* " + info + " */"}
	for _, c := range classes {
		name := string(c)
		name = strings.Replace(name, "-", "_", -1)
		css = append(css, Var("class_"+name).Val(string(c)).Css())
	}
	return ioutil.WriteFile(file, []byte(strings.Join(css, "\n")), os.FileMode(0644))
}

func CompileIds(ids []goh4.Id, file string, info string) error {
	css := []string{"/* " + info + " */"}
	for _, c := range ids {
		name := string(c)
		name = strings.Replace(name, "-", "_", -1)
		css = append(css, Var("id_"+name).Val(string(c)).Css())
	}
	return ioutil.WriteFile(file, []byte(strings.Join(css, "\n")), os.FileMode(0644))
}

//func (ø var_) Selector() string      { return fmt.Sprintf("#{%s}", ø.String()) }

type param struct {
	Var      var_
	default_ string
}

func Param(name string) param { return param{Var: var_{Name_: name}} }

func (ø param) Default(s string) param {
	return param{var_{Name_: ø.Var.Name_, Value: ø.Var.Value}, s}
}

func (ø param) AsArg() (s string) {
	if ø.default_ == "" {
		s = fmt.Sprintf("%s", ø.Var.Name())
	} else {
		s = fmt.Sprintf("%s: %s", ø.Var.Name(), ø.default_)
	}
	return
}

func Exp(format string, i ...interface{}) string {
	return fmt.Sprintf("("+format+")", i...)
}

func (ø param) String() (s string)   { return ø.Var.Name() }
func (ø param) Css() (s string)      { return ø.Var.Name() }
func (ø param) Interpol() (s string) { return fmt.Sprintf(`@{%s}`, ø.Var.Name_) }

func Escape(s string) string {
	return fmt.Sprintf(`~%#v;`, s)
}

type mixin struct {
	*css.RuleStruct
	name string
	Args []string
}

func Mixin(name goh4.Class, params ...param) (ø *mixin) {
	args := []string{}
	for _, param := range params {
		args = append(args, param.AsArg())
	}
	ø = &mixin{&css.RuleStruct{}, string(name), args}
	return
}

func Arguments() string {
	return "@arguments"
}

func NameSpace(id goh4.Id, class goh4.Class) string {
	return id.String() + " > " + class.String() + ";"
}

func (ø *mixin) String() string {
	return "." + ø.name
}

func (ø *mixin) Css() string {
	ø.RuleStruct.Selector = ø
	return ø.RuleStruct.String()
}

func (ø *mixin) argString() (s string) { return strings.Join(ø.Args, "; ") }

func (ø *mixin) Selector() (s string) {
	if len(ø.Args) == 0 {
		s = "." + ø.name
	} else {
		s = fmt.Sprintf(".%s (%s)", ø.name, ø.argString())
	}
	return
}

type CallStruct mixin

func Call(fn string, args ...string) (ø *CallStruct) {
	ø = &CallStruct{}
	ø.name = fn
	ø.Args = args
	return
}

func (ø *CallStruct) Style() (s string) {
	return ø.String() + ";"
}

func (ø *CallStruct) String() (s string) {
	m := mixin(*ø)
	s = fmt.Sprintf("%s(%s)", m.name, m.argString())
	return
}
