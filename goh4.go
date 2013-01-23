package goh4

import (
	"html"
	"regexp"
	"strings"
)

type Stringer interface {
	String() string
}

type Context string

func (ø Context) String() string { return string(ø) }

type Comment string

func (ø Comment) String() string { return string(ø) }

type Class string

func (ø Class) String() string { return string(ø) }

type Id string

func (ø Id) String() string { return string(ø) }

type Text string

func (ø Text) String() string { return string(ø) }

type Html string

func (ø Html) String() string { return string(ø) }

type Tag string

func (ø Tag) String() string { return string(ø) }

type tags []Tag

func Tags(ø ...string) (a tags) {
	a = tags{}
	for _, s := range ø {
		a = append(a, Tag(s))
	}
	return
}

func removeWhiteSpace(in string) string {
	reg := regexp.MustCompile(`\s`)
	return reg.ReplaceAllString(in, "")
}

func (ø tags) String() string {
	str := []string{}
	for _, t := range ø {
		str = append(str, t.String())
	}
	return strings.Join(str, ", ")
}

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
func Style(ø ...string) (s styles) {
	s = styles{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, style{ø[i], ø[i+1]})
	}
	return
}

type attr struct {
	Key   string
	Value string
}

func (ø attr) String() string { return " " + ø.Key + `="` + html.EscapeString(ø.Value) + `"` }

type attrs []attr

func (ø attrs) String() (s string) {
	ss := []string{}
	for _, atr := range ø {
		ss = append(ss, atr.String())
	}
	return strings.Join(ss, " ")
}

// helper to easily create multiple attrs
// use is like this
// Attr("width","200px","height","30px","value","hiho")
func Attr(ø ...string) (s attrs) {
	s = []attr{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, attr{ø[i], ø[i+1]})
	}
	return
}
