package tags

import (
	. "github.com/metakeule/goh4"
)

// pseudo element for placeholder
func Doc(objects ...interface{}) (t *Element) {
	t = NewElement(Tag("doc"), WithoutDecoration)
	t.Add(objects...)
	return
}

func docType(doctype string, objects ...interface{}) (t *Element) {
	t = NewElement(Tag("doc"), WithoutDecoration)
	t.Add(Html(doctype))
	t.Add(objects...)
	return
}

// the doctypes are taken from http://www.w3.org/QA/2002/04/valid-dtd-list.html#DTD

func HTML4_01Strict(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
   "http://www.w3.org/TR/html4/strict.dtd">`, objects...)
}

func HTML4_01Transitional(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/loose.dtd">`, objects...)
}

func HTML4_01Frameset(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN"
   "http://www.w3.org/TR/html4/frameset.dtd">`, objects...)
}

func XHTML1_0Strict(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`, objects...)
}

func XHTML1_0Transitional(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`, objects...)
}

func XHTML1_0Frameset(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"
   "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`, objects...)
}

func XHTML1_1(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
   "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`, objects...)
}

func XHTML1_1Basic(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN"
    "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`, objects...)
}

func HTML5(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE HTML>`, objects...)
}

func MathML2_0(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE math PUBLIC "-//W3C//DTD MathML 2.0//EN"
  "http://www.w3.org/Math/DTD/mathml2/mathml2.dtd">`, objects...)
}

func MathML1_01(objects ...interface{}) (t *Element) {
	return docType(`<!DOCTYPE math SYSTEM
  "http://www.w3.org/Math/DTD/mathml1/mathml.dtd">`, objects...)
}
