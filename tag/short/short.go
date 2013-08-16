package short

import (
	ĸ "github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/attr"
	. "github.com/metakeule/goh4/tag"
)

func JsSrc(url string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{Type("text/javascript"), Src(url)}
	return SCRIPT(append(params, objects...)...)
}

func CssHref(url string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{Rel("stylesheet"), Type("text/css"), Href(url)}
	return LINK(append(params, objects...)...)
}

func FormGet(action string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{Method("get"), Action(action)}
	return FORM(append(params, objects...)...)
}

func FormPost(action string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{Method("post"), Action(action)}
	return FORM(append(params, objects...)...)
}

//t.Add(Attr("enctype", "multipart/form-data", "method", "post"))
func FormPostMultipart(action string, objects ...interface{}) *ĸ.Element {
	//t.Add(Attr("enctype", "multipart/form-data", "method", "post"))
	params := []interface{}{Method("post"), Action(action), MultiPart()}
	return FORM(append(params, objects...)...)
}

func FormPut(action string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Method("post"),
		Action(action),
		INPUT(
			Name("_method"),
			Value("PUT"),
			Type("hidden"))}
	return FORM(append(params, objects...)...)
}

func FormDelete(action string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Method("post"),
		Action(action),
		INPUT(
			Name("_method"),
			Value("DELETE"),
			Type("hidden"))}
	return FORM(append(params, objects...)...)
}

func inputType(typ string, name string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Type(typ),
		Name(name),
	}
	return INPUT(append(params, objects...)...)
}

func InputHidden(name string, objects ...interface{}) *ĸ.Element {
	return inputType("hidden", name, objects...)
}

func InputSubmit(name string, objects ...interface{}) *ĸ.Element {
	return inputType("submit", name, objects...)
}

func InputText(name string, objects ...interface{}) *ĸ.Element {
	return inputType("text", name, objects...)
}

func InputButton(name string, objects ...interface{}) *ĸ.Element {
	return inputType("button", name, objects...)
}

func InputPassword(name string, objects ...interface{}) *ĸ.Element {
	return inputType("password", name, objects...)
}

func InputRadio(name string, objects ...interface{}) *ĸ.Element {
	return inputType("radio", name, objects...)
}

func InputCheckbox(name string, objects ...interface{}) *ĸ.Element {
	return inputType("checkbox", name, objects...)
}

func InputFile(name string, objects ...interface{}) *ĸ.Element {
	return inputType("file", name, objects...)
}

func AHref(url string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{Href(url)}
	return A(append(params, objects...)...)
}

func ImgSrc(src string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{Src(src)}
	return IMG(append(params, objects...)...)
}

func LabelFor(for_ string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{For(for_)}
	return LABEL(append(params, objects...)...)
}

func Charset(charset string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{ATTR("charset", charset)}
	return META(append(params, objects...)...)
}

func CharsetUtf8(objects ...interface{}) *ĸ.Element {
	return Charset("utf-8", objects...)
}
