package tag

var (
	inlineElements      = map[string]bool{}
	selfclosingElements = map[string]bool{}
	formElements        = map[string]bool{}
)

func AddInline(elements ...string) {
	for _, el := range elements {
		inlineElements[el] = true
	}
}

func IsInline(element string) bool {
	return inlineElements[element]
}

func AddSelfclosing(elements ...string) {
	for _, el := range elements {
		selfclosingElements[el] = true
	}
}

func IsSelfclosing(element string) bool {
	return selfclosingElements[element]
}

func AddFormElement(elements ...string) {
	for _, el := range elements {
		formElements[el] = true
	}
}

func IsFormElement(element string) bool {
	return formElements[element]
}

func init() {
	AddInline(
		"a", "b", "bdi", "command", "data", "datalist", "em", "figcaption", "i", "img", "input",
		"keygen", "mark", "meter", "output", "progress", "rp", "rt", "ruby", "select", "span",
		"strong", "summary", "time", "wbr",
	)

	AddSelfclosing(
		"area", "base", "br", "col", "command", "embed", "frame", "hr", "iframe", "img",
		"input", "keygen", "link", "meta", "param", "track", "wbr",
	)

	AddFormElement(
		"button", "input", "keygen", "object", "select", "textarea",
	)
}
