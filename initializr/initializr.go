package initializr

import (
	"fmt"
	. "github.com/metakeule/goh4"
)

type Initializr struct {
	*Template
	pre        string
	post       string
	Head       *Element
	Body       *Element
	Navigation *Element
	Main       *Element
	CssPath    string
	JsPath     string
}

var pre string = `<!DOCTYPE html>
<!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]>         <html class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!--> <html class="no-js"> <!--<![endif]-->`

var post string = `</html>`

var infoChromeFrame = `<!--[if lt IE 7]>
            <p class="chromeframe">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> or <a href="http://www.google.com/chromeframe/?redirect=true">activate Google Chrome Frame</a> to improve your experience.</p>
        <![endif]-->`

var infoBootstrap = ` <!-- This code is taken from http://twitter.github.com/bootstrap/examples/hero.html -->`

var style string = `
  body {
      padding-top: 60px;
      padding-bottom: 40px;
  }
`
var lf string = "\n"
var charset = Meta(Attr("charset", "utf-8"))
var chromeFrame = Meta(Attr("http-equiv", "X-UA-Compatible"), Attr("content", "IE=edge,chrome=1"))
var viewport = Meta(Attr("name", "viewport"), Attr("content", "width=device-width"))

func (ø *Initializr) String() string {
	return pre + ø.Template.String() + ø.post
}

// expects a string with placeholder %s
func (ø *Initializr) jsPathed(s string) string {
	return fmt.Sprintf(s, ø.JsPath)
}

// expects a string with placeholder %s
func (ø *Initializr) cssPathed(s string) string {
	return fmt.Sprintf(s, ø.CssPath)
}

func (ø *Initializr) AddCssFile(file string) {
	ø.Head.Add(Link(Attr("rel", "stylesheet"), Attr("href", file)), lf)
}

func (ø *Initializr) AddMetaDescription(descr string) {
	ø.Head.Add(Meta(Attr("name", "description"), Attr("content", descr)), lf)
}

func (ø *Initializr) AddStyle(css string) {
	styleTag := NewElement("style", Invisible)
	styleTag.Add(css)
	ø.Head.Add(styleTag)
}

func (ø *Initializr) AddScriptFile(file string) {
	ø.Body.Add(Script(Attr("src", file)))
}

func (ø *Initializr) AddScript(js string) {
	ø.Body.Add(Script(Html(js)))
}

func (ø *Initializr) AddGoogleAnalytics(account string) {
	ø.AddScript(`var _gaq=[['_setAccount','` + account + `'],['_trackPageview']];
            (function(d,t){var g=d.createElement(t),s=d.getElementsByTagName(t)[0];
            g.src=('https:'==location.protocol?'//ssl':'//www')+'.google-analytics.com/ga.js';
            s.parentNode.insertBefore(g,s)}(document,'script'));`)
}

func (ø *Initializr) NewContainter() *Element {
	return Div(Class("container"))
}

func (ø *Initializr) SetupHead() {
	ø.AddCssFile(ø.cssPathed("%s/bootstrap.min.css"))
	ø.AddStyle(style)
	ø.AddCssFile(ø.cssPathed("%s/bootstrap-responsive.min.css"))
	ø.AddCssFile(ø.cssPathed("%s/main.css"))
	ø.Head.Add(Script(Attr("src", ø.jsPathed("%s/modernizr-2.6.2-respond-1.1.0.min.js"))))
}

func (ø *Initializr) SetupBody() {
	ø.Body.Add(Html(infoChromeFrame))
	ø.Body.Add(Html(infoBootstrap))
	ø.Navigation = ø.NewContainter()
	navBar := Div(
		Class("navbar"),
		Class("navbar-inverse"),
		Class("navbar-fixed-top"),
		Div(
			Class("navbar-inner"),
			ø.Navigation))
	ø.Body.Add(navBar)

	ø.Main = ø.NewContainter()
	ø.Body.Add(ø.Main)

	ø.AddScriptFile("//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js")
	ø.AddScript(ø.jsPathed(`window.jQuery || document.write('<script src="%s/jquery-1.8.3.min.js"><\/script>')`))
	//ø.AddScriptFile(ø.jsPathed(`%s/jquery-1.8.3.min.js`))
	ø.AddScriptFile(ø.jsPathed("%s/bootstrap.min.js"))
	ø.AddScriptFile(ø.jsPathed("%s/main.js"))

}

func Layout() (layout *Initializr) {
	layout = &Initializr{
		Template: &Template{Element: Doc()},
		pre:      pre,
		post:     post,
		Head:     Head(lf, charset, lf, chromeFrame, lf, viewport, lf),
		Body:     Body(),
		CssPath:  "/css",
		JsPath:   "/js",
	}
	//layout.Template.Tag = layout.Body
	layout.Template.Add(layout.Head, layout.Body)
	layout.SetupHead()
	layout.SetupBody()
	return
}
