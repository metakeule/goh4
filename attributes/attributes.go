package attributes

import (
	. "github.com/metakeule/goh4"
)

const (
	Hidden   = "hidden"
	Checkbox = "checkbox"
)

func Rel(relation string) Attrs { return Attr("rel", relation) }
func Href(url string) Attrs     { return Attr("href", url) }
func Type(ty string) Attrs      { return Attr("type", ty) }
func Media(md string) Attrs     { return Attr("media", md) }
func Src(src string) Attrs      { return Attr("src", src) }
func Name(n string) Attrs       { return Attr("name", n) }
func Value(v string) Attrs      { return Attr("value", v) }
func Method(m string) Attrs     { return Attr("method", m) }
func Action(a string) Attrs     { return Attr("action", a) }
func For(f string) Attrs        { return Attr("for", f) }
func Width(f string) Attrs      { return Attr("width", f) }
func Height(f string) Attrs     { return Attr("height", f) }
func Checked() Attrs            { return Attr("checked", "checked") }
func Required() Attrs           { return Attr("required", "required") }
func OnSubmit(js string) Attrs  { return Attr("onsubmit", "javascript:"+js) }
func OnClick(js string) Attrs   { return Attr("onclick", "javascript:"+js) }
