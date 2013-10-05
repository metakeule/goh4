package goh4

import (
	"testing"
)

func TestTemplatePlaceholder(t *testing.T) {
	text := Text("text").Placeholder()
	strong := NewElement(Tag("strong"))
	strong.Add(text)
	content := strong.Placeholder("content")

	span := NewElement(Tag("span"))
	span.Add(content)
	layout := span.Placeholder("layout")

	l := layout.MustReplace(content.MustReplace(text.Set("<hu>")))
	exp := `<span><strong>&lt;hu&gt;</strong></span>`

	if l.String() != exp {
		err(t, "incorrect template compilation", l.String(), exp)
	}
}

func TestTemplateAssign(t *testing.T) {
	body := NewElement(Tag("body"))
	div := NewElement(Tag("div"))
	div.Add(Id("content"))
	body.Add(
		Id("contact"),
		div)
	layout := NewTemplate(body)

	content := NewElement(Tag("p"))
	content.Add(Text("hiho"))

	layout.Assign(Id("content"), content)

	if body.Children()[0].Children()[0] != content {
		err(t, "incorrect template assignment", body.Children()[0], content)
	}

	other := NewElement(Tag("span"))
	other.Add(Text("huhu"))

	layout.Assign(Id("content"), other)

	if body.Children()[0].Children()[0] != other {
		err(t, "incorrect template reassignment", body.Children()[0], other)
	}

	if e := layout.Assign(Id("not-there"), other); e == nil {
		err(t, "incorrect template assignment", "not exists", nil)
	}
}

// func TestTemplateAddCss(t *testing.T) {
// 	css := NewCss(Tag("body"), Style("color", "red"))
// 	head := Head()
// 	layout := NewTemplate(Doc(head))
// 	layout.AddCss(css)

// 	if len(head.Children()) == 0 {
// 		err(t, "incorrect template AddCss head got no children", 0, 1)
// 	}

// 	style := head.Children()[0]
// 	if style.Tag() != "style" {
// 		err(t, "incorrect template AddCss style tag", style.Tag(), "style")
// 	}

// 	if style.InnerHtml() != css.String() {
// 		err(t, "incorrect template AddCss style content", style.String(), css.String())
// 	}

// 	layout.AddCss(Text("/* My Comment */"))

// 	if len(head.Children()) == 2 {
// 		err(t, "incorrect template AddCss wrong number of style tags", 2, 1)
// 	}

// }
