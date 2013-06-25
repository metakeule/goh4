package shortcuts

import (
	ĸ "github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/attributes"
	. "github.com/metakeule/goh4/tags"
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

func InputHidden(name string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Type("hidden"),
		Name(name),
	}
	return INPUT(append(params, objects...)...)
}

func InputSubmit(name string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Type("submit"),
		Name(name),
	}
	return INPUT(append(params, objects...)...)
}

func InputText(name string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Type("text"),
		Name(name),
	}
	return INPUT(append(params, objects...)...)
}

func InputButton(name string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Type("button"),
		Name(name),
	}
	return INPUT(append(params, objects...)...)
}

func InputPassword(name string, objects ...interface{}) *ĸ.Element {
	params := []interface{}{
		Type("password"),
		Name(name),
	}
	return INPUT(append(params, objects...)...)
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
