package goh4

import (
	"fmt"
	"strings"
	"testing"
)

var tagTests = []struct {
	fn   func(...interface{}) *Element
	tag  string
	html string
}{
	{A, "a", "<a></a>"},
	{Body, "body", "<body></body>"},
	{Button, "button", "<button></button>"},
	{B, "b", "<b></b>"},
	{Div, "div", "<div></div>"},
	{Footer, "footer", "<footer></footer>"},
	{P, "p", "<p></p>"},
	// since the order of keys in a map doesn't matter in go, the order of
	// attributes (that does not matter as well), will not always be the same,
	// so this test will fail from time to time, disable it
	//{Form, "form", "<form enctype=\"multipart/form-data\" method=\"post\"></form>"},
	{Head, "head", "<head></head>"},
	{Hr, "hr", "<hr />"},
	{Iframe, "iframe", "<iframe />"},
	{I, "i", "<i></i>"},
	{Img, "img", "<img />"},
	{Input, "input", "<input />"},
	{Label, "label", "<label></label>"},
	{Li, "li", "<li></li>"},
	{Link, "link", "<link />"},
	{Meta, "meta", "<meta />"},
	{Option, "option", "<option></option>"},
	{Pre, "pre", "<pre></pre>"},
	{Script, "script", "<script></script>"},
	{Select, "select", "<select></select>"},
	{Span, "span", "<span></span>"},
	{Strong, "strong", "<strong></strong>"},
	{Table, "table", "<table></table>"},
	{Tr, "tr", "<tr></tr>"},
	{Th, "th", "<th></th>"},
	{Td, "td", "<td></td>"},
	{Textarea, "textarea", "<textarea></textarea>"},
	{Title, "title", "<title></title>"},
	{Ul, "ul", "<ul></ul>"},
	{Ol, "ol", "<ol></ol>"},
	{Doc, "doc", ""},
}

func TestTags(t *testing.T) {
	for _, tt := range tagTests {
		res := tt.fn()
		if res.Tag() != tt.tag {
			err(t, "incorrect tag", res.Tag(), tt.tag)
		}

		if strings.Replace(res.String(), "\n", "", -1) != tt.html {
			err(t, "incorrect html", res.String(), tt.html)
		}
	}
}

func TestHeadings(t *testing.T) {
	for i := 1; i < 7; i++ {
		r := H(i)
		tag := fmt.Sprintf("h%v", i)
		if r.Tag() != tag {
			err(t, "incorrect tag", r.Tag(), tag)
		}

		html := "<" + tag + "></" + tag + ">"
		if strings.Replace(r.String(), "\n", "", -1) != html {
			err(t, "incorrect html", r.String(), html)
		}
	}
}
