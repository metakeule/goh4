package goh4

import (
	"github.com/metakeule/template"
)

type CompiledTemplate struct {
	*template.Template
	ElementTemplate *Template
}

// returns a *CompiledTemplate that is a template.Template (see github.com/metakeule/template)
// the template can then be initialized with New and merged with placeholders with Replace and Merge
// if you need to change the original template again, you can get it via CompiledTemplate.ElementTemplate
// then call Compile() again to get a new CompiledTemplate
func (ø *Template) Compile() (c *CompiledTemplate, ſ error) {
	c = &CompiledTemplate{ElementTemplate: ø}
	c.Template, ſ = template.New(ø.String())
	if ſ != nil {
		c = nil
		return
	}
	return
}

// panics on error
func (ø *Template) MustCompile() *CompiledTemplate {
	c, e := ø.Compile()
	if e != nil {
		panic(e.Error())
	}
	return c
}
