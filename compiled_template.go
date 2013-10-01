package goh4

import (
	"github.com/metakeule/templ"
)

type CompiledTemplate struct {
	*templ.Template
	ElementTemplate *Template
}

// returns a *CompiledTemplate that is a template.Template (see github.com/metakeule/template)
// the template can then be initialized with New and merged with placeholders with Replace and Merge
// if you need to change the original template again, you can get it via CompiledTemplate.ElementTemplate
// then call Compile() again to get a new CompiledTemplate
func (ø *Template) Compile(name string) (c *CompiledTemplate, ſ error) {
	c = &CompiledTemplate{ElementTemplate: ø}
	c.Template = templ.New(name).MustAdd(ø.String())
	ſ = c.Template.Parse()
	if ſ != nil {
		c = nil
		return
	}
	return
}

// panics on error
func (ø *Template) MustCompile(name string) *CompiledTemplate {
	c, e := ø.Compile(name)
	if e != nil {
		panic(e.Error())
	}
	return c
}
