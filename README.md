goh4
====

[![Build Status](https://secure.travis-ci.org/metakeule/goh4.png)](http://travis-ci.org/metakeule/goh4)

WARNING: This package is in flux. You should not use it in production yet.
API will / might break.

type save html and css generation for golang with functions (inspired by http://godoc.org/github.com/daaku/go.h)

example
-------

```go
package main

import (
	"fmt"
	"github.com/metakeule/goh4"
)

func main(){
	font := NewCss(
		Class("default-font"),
		Style("font-size","20","font-weight","normal"))

	css := NewCss(
		Class("yellow-button"),
		Tags("a","button"),
		Style("background-color","yellow"),
		font)

	a := A(css)

	doc := NewTemplate(
		Doc(
			Head(),
			Body(a)))

	doc.AddCss(css)

	fmt.Println(doc)
}
```
results in

```html
<head>
	<style>
		a.yellow-button,
		button.yellow-button {
			background-color: yellow;
			font-weight: normal;	/* inherited from ».default-font« */
			font-size: 20;	/* inherited from ».default-font« */
		}
	</style>
</head>
<body><a class="yellow-button"></a></body>
```

Documentation
-------------

see http://godoc.org/github.com/metakeule/goh4

TODO
----

- correct handling of css contexts for matching elements, will be done
	with the new exp/html package and http://code.google.com/p/cascadia
