package tag

import (
	. "gopkg.in/metakeule/goh4.v5"
)

func CLASS(c string) Class                          { return Class(c) }
func ID(id string) Id                               { return Id(id) }
func HTML(html string) Html                         { return Html(html) }
func TAG(s string) Tag                              { return Tag(s) }
func ATTR(key1, val1 string, ø ...string) (s Attrs) { return Attr(key1, val1, ø...) }
