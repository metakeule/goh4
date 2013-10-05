package tag

import (
	"fmt"
	"github.com/metakeule/goh4"
	"strings"
	"testing"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %v, should be %v\n", is, shouldbe)
}

var tagTests = []struct {
	fn   func(...interface{}) *goh4.Element
	tag  string
	html string
}{
	{A, "a", "<a></a>"},
	{BODY, "body", "<body></body>"},
	{BUTTON, "button", "<button></button>"},
	{B, "b", "<b></b>"},
	{DIV, "div", "<div></div>"},
	{FOOTER, "footer", "<footer></footer>"},
	{P, "p", "<p></p>"},
	// since the order of keys in a map doesn't matter in go, the order of
	// attributes (that does not matter as well), will not always be the same,
	// so this test will fail from time to time, disable it
	//{Form, "form", "<form enctype=\"multipart/form-data\" method=\"post\"></form>"},
	{HEAD, "head", "<head></head>"},
	{HR, "hr", "<hr />"},
	{IFRAME, "iframe", "<iframe />"},
	{I, "i", "<i></i>"},
	{IMG, "img", "<img />"},
	{INPUT, "input", "<input />"},
	{LABEL, "label", "<label></label>"},
	{LI, "li", "<li></li>"},
	{LINK, "link", "<link />"},
	{META, "meta", "<meta />"},
	{OPTION, "option", "<option></option>"},
	{PRE, "pre", "<pre></pre>"},
	{SCRIPT, "script", "<script></script>"},
	{SELECT, "select", "<select></select>"},
	{SPAN, "span", "<span></span>"},
	{STRONG, "strong", "<strong></strong>"},
	{TABLE, "table", "<table></table>"},
	{TR, "tr", "<tr></tr>"},
	{TH, "th", "<th></th>"},
	{TD, "td", "<td></td>"},
	{TEXTAREA, "textarea", "<textarea></textarea>"},
	{TITLE, "title", "<title></title>"},
	{UL, "ul", "<ul></ul>"},
	{OL, "ol", "<ol></ol>"},
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
