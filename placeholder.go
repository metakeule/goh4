package goh4

import (
	"fmt"
	"github.com/metakeule/templ"
	"html"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var Escaper = templ.Escaper{
	"text":     handleStrings(html.EscapeString, true),
	"":         handleStrings(html.EscapeString, true),
	"html":     handleStrings(idem, true),
	"px":       units("%vpx"),
	"%":        units("%v%%"),
	"em":       units("%vem"),
	"pt":       units("%vpt"),
	"urlparam": handleStrings(url.QueryEscape, false),
}

type view struct {
	*templ.View
}

type placeholder struct {
	templ.Placeholder
}

func (p placeholder) String() string {
	return "@@" + p.Name() + "@@"
}

func (v *view) Placeholder(field string) placeholder {
	return placeholder{v.View.Placeholder(field)}
}

func View(stru interface{}, tag string) *view {
	return &view{Escaper.View(stru, tag)}
}

//Placeholder("Link")

func units(format string) func(interface{}) string {
	fn := func(in interface{}) (out string) {
		switch v := in.(type) {
		case int, int8, int16, int32, int64, float32, float64:
			return fmt.Sprintf(format, v)
		default:
			panic("unsupported type: " + fmt.Sprintf("%v (%T)", v, v))
		}
	}
	return fn
}

// takes different types and outputs a string
func Str(in interface{}) string {
	switch v := in.(type) {
	case *templ.Placeholder:
		return "@@" + v.Name() + "@@"
	case templ.Placeholder:
		return "@@" + v.Name() + "@@"
	case Stringer:
		return v.String()
	case string:
		return v
	}
	panic("unsupported type: " + fmt.Sprintf("%v (%T)", in, in))
}

func idem(in string) (out string) { return in }

// is  used by FillStruct, see github.com/metakeule/template
/*
var Transformer = map[string]func(interface{}) string{
	"text": handleStrings(html.EscapeString, true),
	"html": handleStrings(idem, false),
}
*/

/*
// shortcut for template.FillStruct with transformer
func FillStruct(ptrToStruct interface{}) map[string]string {
	return template.FillStruct("goh4", Transformer, ptrToStruct)
}
*/

func handleStrings(trafo func(string) string, allowAll bool) func(interface{}) string {
	return func(in interface{}) (out string) {
		if in == nil {
			return ""
		}
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
	templ.Setter
	Set(val interface{}) templ.Setter
	Setf(format string, val ...interface{}) templ.Setter
	String() string
	Type() interface{}
}

type typedPlaceholder struct {
	templ.Placeholder
	typ interface{}
}

func (ø typedPlaceholder) String() string {
	return "@@" + ø.Name() + "@@"
}

func (ø typedPlaceholder) Type() interface{} {
	return ø.typ
}

func newTPh(ph templ.Placeholder, i interface{}) typedPlaceholder {
	return typedPlaceholder{ph, i}
}

func stripGoPath(path string) {
	gopath := strings.Split(os.Getenv("GOPATH"), ":")[0]
	if gopath == "" {
		panic("GOPATH not set")
	}
}

func caller(skip int) string {
	_, file, num, ok := runtime.Caller(skip)
	if !ok {
		panic("can't get caller")
	}
	return fmt.Sprintf("%s:%v", file, num)
}

func (ø Comment) Placeholder() Placeholder {
	return newTPh(templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}

func (ø Class) Placeholder() Placeholder {
	return newTPh(templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}
func (ø Id) Placeholder() Placeholder {
	return newTPh(templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}
func (ø Html) Placeholder() Placeholder {
	return newTPh(templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}
func (ø Text) Placeholder() Placeholder {
	t := templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø), handleStrings(html.EscapeString, true))
	return newTPh(t, ø)
}

func (ø SingleAttr) Placeholder() Placeholder {
	t := templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+ø.Value, handleStrings(html.EscapeString, true))
	return newTPh(t, ø)
}

func (ø Tag) Placeholder() Placeholder {
	return newTPh(templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+string(ø)), ø)
}

func (ø Style) Placeholder() Placeholder {
	t := templ.NewPlaceholder(reflect.TypeOf(ø).Name()+"."+ø.Value, handleStrings(html.EscapeString, true))
	return newTPh(t, ø)
}
