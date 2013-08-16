package attr

import (
	. "github.com/metakeule/goh4"
)

const (
	Hidden   = "hidden"
	Checkbox = "checkbox"
)

func Rel(relation string) SingleAttr { return SingleAttr{"rel", relation} }
func Href(url string) SingleAttr     { return SingleAttr{"href", url} }
func Type(ty string) SingleAttr      { return SingleAttr{"type", ty} }
func Media(md string) SingleAttr     { return SingleAttr{"media", md} }
func Src(src string) SingleAttr      { return SingleAttr{"src", src} }
func Name(n string) SingleAttr       { return SingleAttr{"name", n} }
func Value(v string) SingleAttr      { return SingleAttr{"value", v} }
func Alt(alttext string) SingleAttr  { return SingleAttr{"alt", alttext} }
func Title(text string) SingleAttr   { return SingleAttr{"title", text} }
func Method(m string) SingleAttr     { return SingleAttr{"method", m} }
func Action(a string) SingleAttr     { return SingleAttr{"action", a} }
func For(f string) SingleAttr        { return SingleAttr{"for", f} }
func Width(f string) SingleAttr      { return SingleAttr{"width", f} }
func Height(f string) SingleAttr     { return SingleAttr{"height", f} }
func Checked() SingleAttr            { return SingleAttr{"checked", "checked"} }
func Required() SingleAttr           { return SingleAttr{"required", "required"} }
func OnSubmit(js string) SingleAttr  { return SingleAttr{"onsubmit", "javascript:" + js} }
func OnClick(js string) SingleAttr   { return SingleAttr{"onclick", "javascript:" + js} }
func Enctype(f string) SingleAttr    { return SingleAttr{"enctype", f} }
func Target(f string) SingleAttr     { return SingleAttr{"target", f} }
func MultiPart() SingleAttr          { return Enctype("multipart/form-data") }
func DataToggle(f string) SingleAttr { return SingleAttr{"data-toggle", f} }
func DataTarget(f string) SingleAttr { return SingleAttr{"data-target", f} }
func DataId(f string) SingleAttr     { return SingleAttr{"data-id", f} }

// RDFa
func About(a string) SingleAttr    { return SingleAttr{"about", a} }
func TypeOf(a string) SingleAttr   { return SingleAttr{"typeof", a} }
func Property(a string) SingleAttr { return SingleAttr{"property", a} }
