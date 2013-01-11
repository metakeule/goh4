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

type Tags []string

func (ø Tags) ToTagArr() (a []Tag) {
	a = []Tag{}
	for _, s := range ø {
		a = append(a, Tag(s))
	}
	return
}

func removeWhiteSpace(in string) string {
	reg := regexp.MustCompile(`\s`)
	return reg.ReplaceAllString(in, "")
}

func (ø Tags) String() string {
	return strings.Join(ø, ", ")
}

type Style struct {
	Key   string
	Value string
}

func (ø Style) String() string { return ø.Key + ": " + ø.Value + ";" }

// helper to easily create multiple styles
// use is like this
// Styles{"width","200px","height","30px","color","green"}.ToStyleArr()
type Styles []string

func (ø Styles) ToStyleArr() (s []Style) {
	s = []Style{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, Style{ø[i], ø[i+1]})
	}
	return
}

func (ø Styles) String() (s string) {
	ss := []string{}
	styles := ø.ToStyleArr()
	for _, style := range styles {
		ss = append(ss, style.String())
	}
	return strings.Join(ss, " ")
}

type Attr struct {
	Key   string
	Value string
}

func (ø Attr) String() string { return " " + ø.Key + `="` + html.EscapeString(ø.Value) + `"` }

// helper to easily create multiple attrs
// use is like this
// Attrs{"width","200px","height","30px","value","hiho"}.ToAttrArr()
type Attrs []string

func (ø Attrs) ToAttrArr() (s []Attr) {
	s = []Attr{}
	for i := 0; i < len(ø); i = i + 2 {
		s = append(s, Attr{ø[i], ø[i+1]})
	}
	return
}

func (ø Attrs) String() (s string) {
	ss := []string{}
	attrs := ø.ToAttrArr()
	for _, attr := range attrs {
		ss = append(ss, attr.String())
	}
	return strings.Join(ss, " ")
}

/*
func fmt.Errorf(message string, obj ...interface{}) error {
	return errors.New(fmt.Sprintf(message+"\n", obj...))
}
*/
