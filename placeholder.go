package goh4

import (
	"fmt"
	"github.com/metakeule/template"
	"html"
	"reflect"
)

func idem(in string) (out string) { return in }

// is  used by FillStruct, see github.com/metakeule/template
var Transformer = map[string]func(interface{}) string{
	"text": handleStrings(html.EscapeString, true),
	"html": handleStrings(idem, false),
}

// shortcut for template.FillStruct with transformer
func FillStruct(ptrToStruct interface{}) map[string]string {
	return template.FillStruct("goh4", Transformer, ptrToStruct)
}

func handleStrings(trafo func(string) string, allowAll bool) func(interface{}) string {
	return func(in interface{}) (out string) {
		var s string
		switch v := in.(type) {
		case Stringer:
			s = v.String()
		case string:
			s = v
		default:
			if allowAll {
				s = fmt.Sprintf("%v", v)
			} else {
				panic("unsupported type: " + fmt.Sprintf("%v (%T)", v, v))
			}
		}
		return trafo(s)
	}
}

type Placeholder interface {
	template.Replacer
	Set(val interface{}) template.Placeholder
	Setf(format string, val ...interface{}) template.Placeholder
	String() string
	Type() interface{}
}

type typedPlaceholder struct {
	template.Placeholder
	typ interface{}
}

func (ø typedPlaceholder) Type() interface{} {
	return ø.typ
}

func newTPh(ph template.Placeholder, i interface{}) typedPlaceholder {
	return typedPlaceholder{ph, i}
}

func (ø Comment) Placeholder() Placeholder {
	return newTPh(template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}

func (ø Class) Placeholder() Placeholder {
	return newTPh(template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}
func (ø Id) Placeholder() Placeholder {
	return newTPh(template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}
func (ø Html) Placeholder() Placeholder {
	return newTPh(template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}
func (ø Text) Placeholder() Placeholder {
	t := template.NewPlaceholder(reflect.TypeOf(ø).Name() + "." + string(ø))
	t.Transformer = handleStrings(html.EscapeString, true)
	return newTPh(t, ø)
}

func (ø SingleAttr) Placeholder() Placeholder {
	t := template.NewPlaceholder(reflect.TypeOf(ø).Name() + "." + ø.Value)
	t.Transformer = handleStrings(html.EscapeString, true)
	return newTPh(t, ø)
}

func (ø Tag) Placeholder() Placeholder {
	return newTPh(template.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}

func (ø Style) Placeholder() Placeholder {
	t := template.NewPlaceholder(reflect.TypeOf(ø).Name() + "." + ø.Value)
	t.Transformer = handleStrings(html.EscapeString, true)
	return newTPh(t, ø)
}
