package goh4

func JsSrc(url string, objects ...interface{}) *Element {
	params := []interface{}{Type_("text/javascript"), Src_(url)}
	return Script(append(params, objects...)...)
}

func CssHref(url string, objects ...interface{}) *Element {
	params := []interface{}{Rel_("stylesheet"), Type_("text/css"), Href_(url)}
	return Link(append(params, objects...)...)
}

func FormGet(action string, objects ...interface{}) *Element {
	params := []interface{}{Method_("get"), Action_(action)}
	return Form(append(params, objects...)...)
}

func FormPost(action string, objects ...interface{}) *Element {
	params := []interface{}{Method_("post"), Action_(action)}
	return Form(append(params, objects...)...)
}

func FormPut(action string, objects ...interface{}) *Element {
	params := []interface{}{
		Method_("post"),
		Action_(action),
		Input(
			Name_("_method"),
			Value_("PUT"),
			Type_("hidden"))}
	return Form(append(params, objects...)...)
}

func FormDelete(action string, objects ...interface{}) *Element {
	params := []interface{}{
		Method_("post"),
		Action_(action),
		Input(
			Name_("_method"),
			Value_("DELETE"),
			Type_("hidden"))}
	return Form(append(params, objects...)...)
}

func InputHidden(name string, objects ...interface{}) *Element {
	params := []interface{}{
		Type_("hidden"),
		Name_(name),
	}
	return Input(append(params, objects...)...)
}

func InputSubmit(name string, objects ...interface{}) *Element {
	params := []interface{}{
		Type_("submit"),
		Name_(name),
	}
	return Input(append(params, objects...)...)
}

func InputText(name string, objects ...interface{}) *Element {
	params := []interface{}{
		Type_("text"),
		Name_(name),
	}
	return Input(append(params, objects...)...)
}

func InputButton(name string, objects ...interface{}) *Element {
	params := []interface{}{
		Type_("button"),
		Name_(name),
	}
	return Input(append(params, objects...)...)
}

func AHref(url string, objects ...interface{}) *Element {
	params := []interface{}{Href_(url)}
	return A(append(params, objects...)...)
}

func ImgSrc(src string, objects ...interface{}) *Element {
	params := []interface{}{Src_(src)}
	return Img(append(params, objects...)...)
}
