package fat

import (
	"fmt"
	"gopkg.in/go-on/lib.v2/internal/fat"
	"gopkg.in/metakeule/goh4.v5"
	"gopkg.in/metakeule/meta.v5"
	"github.com/metakeule/templ"
	"reflect"
	"sync"
)

/*
   goh4 support for fat fields
*/

type registry_ struct {
	*sync.RWMutex
	escapeRegistry map[string]string
}

var (
	registry = &registry_{&sync.RWMutex{}, map[string]string{}}

// maps struct-types to maps of field to escaper
)

// register a struct in the registry
// should be called at initialization time
func Register(østruct interface{}) {
	registry.Lock()
	defer registry.Unlock()
	fn := func(field reflect.StructField, val reflect.Value) {
		f, ok := val.Interface().(*fat.Field)
		if ok {
			if f.StructType() == "" {
				panic(fmt.Sprintf("struct %s has no prototype (not initialized with fat.Proto)", reflect.TypeOf(østruct).String()))
			}
			ty := field.Tag.Get("goh4.type")
			_, found := goh4.Escaper[ty]
			if !found {
				panic(fmt.Sprintf("unknown goh4.type for field %s",
					ty, f.Path(),
				))
			}
			registry.escapeRegistry[f.Path()] = ty
		}
	}
	meta.Struct.EachRaw(østruct, fn)
}

func Placeholder(øfield *fat.Field) templ.Placeholder {
	registry.RLock()
	defer registry.RUnlock()
	ty, ok := registry.escapeRegistry[øfield.Path()]
	if !ok {
		panic(fmt.Sprintf("struct of field %s is not registered with goh4/fat", øfield.Path()))
	}
	return templ.NewPlaceholder(øfield.Path(), goh4.Escaper[ty])
}

func Setter(øfield *fat.Field) templ.Setter {
	return Placeholder(øfield).Set(øfield.String())
}

func Setters(østruct interface{}) (s []templ.Setter) {
	fn := func(field string, val interface{}) {
		if f, ok := val.(*fat.Field); ok {
			s = append(s, Setter(f))
		}
	}
	meta.Struct.Each(østruct, fn)
	return
}

// TODO setters for each field of struct

/*
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
*/
