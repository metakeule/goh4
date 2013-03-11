package goh4

import (
	"testing"
)

func TestNewTemplate(t *testing.T) {
	body := Body(Id("contact"))
	layout := NewTemplate(body)

	if layout.Id() != body.Id() {
		err(t, "incorrect template element id", layout.Id(), body.Id())
	}

}

func TestTemplateCompile(t *testing.T) {
	layout := NewTemplate(Span(Html("@@content@@")))
	content := Strong(Text("hiho"))
	compiled := layout.Compile()
	i := compiled.Instance()
	i.AssignString("content", content.String())
	exp := `<span><strong>hiho</strong></span>`

	if i.String() != exp {
		err(t, "incorrect template compilation", i.String(), exp)
	}
}

func TestTemplateAssign(t *testing.T) {
	body := Body(
		Id("contact"),
		Div(
			Id("content")))
	layout := NewTemplate(body)

	content := P(Text("hiho"))

	layout.Assign(Id("content"), content)

	if body.Children()[0].Children()[0] != content {
		err(t, "incorrect template assignment", body.Children()[0], content)
	}

	other := Span(Text("huhu"))

	layout.Assign(Id("content"), other)

	if body.Children()[0].Children()[0] != other {
		err(t, "incorrect template reassignment", body.Children()[0], other)
	}

	if e := layout.Assign(Id("not-there"), other); e == nil {
		err(t, "incorrect template assignment", "not exists", nil)
	}
}

func TestTemplateAddCss(t *testing.T) {
	css := NewCss(Tag("body"), Style("color", "red"))
	head := Head()
	layout := NewTemplate(Doc(head))
	layout.AddCss(css)

	if len(head.Children()) == 0 {
		err(t, "incorrect template AddCss head got no children", 0, 1)
	}

	style := head.Children()[0]
	if style.Tag() != "style" {
		err(t, "incorrect template AddCss style tag", style.Tag(), "style")
	}

	if style.InnerHtml() != css.String() {
		err(t, "incorrect template AddCss style content", style.String(), css.String())
	}

	layout.AddCss(Text("/* My Comment */"))

	if len(head.Children()) == 2 {
		err(t, "incorrect template AddCss wrong number of style tags", 2, 1)
	}

}
