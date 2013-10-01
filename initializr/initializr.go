package initializr

import (
	ŧ "fmt"
	. "github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	"io/ioutil"
	ħ "net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

type Loadable interface {
	Load() string
}

type directText string

func (ø directText) Load() string {
	// ŧ.Printf("loading: %#v\n", ø)
	return string(ø)
}

var isUrl = regexp.MustCompile("https?://")

type file string

// load the file and return as string
func (ø file) Load() string {
	if isUrl.MatchString(string(ø)) {
		// ŧ.Printf("loading: %#v\n", ø)
		resp, ſ := ħ.Get(string(ø))
		if ſ != nil {
			panic(ſ.Error())
		}
		defer resp.Body.Close()
		body, ſ := ioutil.ReadAll(resp.Body)
		if ſ != nil {
			panic(ſ.Error())
		}
		return string(body)
	} else {
		// ŧ.Printf("loading: %#v\n", ø)
		wd, _ := os.Getwd()
		// fmt.Printf("working dir: %s\n", wd)
		f := path.Join(wd, string(ø))
		// fmt.Printf("file: %s\n", f)
		b, ſ := ioutil.ReadFile(f)
		if ſ != nil {
			panic(ſ.Error())
		}
		return string(b)
	}
	return ""
}

type Initializr struct {
	*Template
	pre            string
	post           string
	Head           *Element
	Body           *Element
	Navigation     *Element
	Main           *Element
	CssPath        string
	JsPath         string
	UseCachedFiles bool
	cachedJs       []Loadable
	cachedJsHead   []Loadable
	cachedCss      []Loadable
}

var cachedJsPath = "%s/scripts.js"
var cachedJsHeadPath = "%s/scripts-head.js"
var cachedCssPath = "%s/style.css"

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
var charset = META(Attr("charset", "utf-8"))
var chromeFrame = META(Attr("http-equiv", "X-UA-Compatible"), Attr("content", "IE=edge,chrome=1"))
var viewport = META(Attr("name", "viewport"), Attr("content", "width=device-width"))

func (ø *Initializr) Compile(name string) (*CompiledTemplate, error) {
	ø.Template.Add(Html(post))
	return ø.Template.Compile(name)
}

func (ø *Initializr) MustCompile(name string) *CompiledTemplate {
	t, e := ø.Compile(name)
	if e != nil {
		panic(e.Error())
	}
	return t
}

func (ø *Initializr) CachedJsFile() (path string, content string) {
	js := []string{}
	// ŧ.Println("CachedJsFile called")
	// ŧ.Printf("cachedJs: %#v\n", ø.cachedJs)
	for _, j := range ø.cachedJs {
		// ŧ.Println("cachedjs")
		js = append(js, j.Load())
	}
	path = ø.jsPathed(cachedJsPath)
	content = strings.Join(js, "\n")
	return
}

func (ø *Initializr) CachedJsHeadFile() (path string, content string) {
	js := []string{}
	// ŧ.Println("CachedJsHeadFile called")
	for _, j := range ø.cachedJsHead {
		// ŧ.Println("cachedjsHead")
		js = append(js, j.Load())
	}
	path = ø.jsPathed(cachedJsHeadPath)
	content = strings.Join(js, "\n")
	return
}

func (ø *Initializr) CachedCssFile() (path string, content string) {
	css := []string{}
	// ŧ.Println("CachedCssFile called")
	for _, j := range ø.cachedCss {
		// ŧ.Println("cachedCss")
		css = append(css, j.Load())
	}
	path = ø.cssPathed(cachedCssPath)
	content = strings.Join(css, "\n")
	return
}

// expects a string with placeholder %s
func (ø *Initializr) jsPathed(s string) string {
	return ŧ.Sprintf(s, ø.JsPath)
}

// expects a string with placeholder %s
func (ø *Initializr) cssPathed(s string) string {
	return ŧ.Sprintf(s, ø.CssPath)
}

func (ø *Initializr) AddCssFile(f string) {
	ø.cachedCss = append(ø.cachedCss, file(f))
	// ŧ.Printf("adding css file %v, cachedcss: %#v\n", f, ø.cachedCss)
	if !ø.UseCachedFiles {
		ø.Head.Add(LINK(Attr("rel", "stylesheet"), Attr("href", f)), lf)
	}
}

func (ø *Initializr) AddMetaDescription(descr string) {
	ø.Head.Add(META(Attr("name", "description"), Attr("content", descr)), lf)
}

func (ø *Initializr) AddStyle(css string) {
	// ŧ.Printf("adding css %v cachedcss: %#v\n", css, ø.cachedCss)
	ø.cachedCss = append(ø.cachedCss, directText(css))
	if !ø.UseCachedFiles {
		styleTag := NewElement("style", Invisible)
		styleTag.Add(Html("\n" + css + "\n"))
		ø.Head.Add(styleTag)
	}
}

func (ø *Initializr) AddScriptFile(f string) {
	// ŧ.Printf("adding js file %v cachedJs: %#v\n", f, ø.cachedJs)
	ø.cachedJs = append(ø.cachedJs, file(f))
	if !ø.UseCachedFiles {
		ø.Body.Add(SCRIPT(Attr("src", f)))
	}
}

func (ø *Initializr) AddScriptFileHead(f string) {
	// ŧ.Printf("adding js file head %v cachedJsHead: %#v\n", f, ø.cachedJsHead)
	ø.cachedJsHead = append(ø.cachedJsHead, file(f))
	if !ø.UseCachedFiles {
		ø.Head.Add(SCRIPT(Attr("src", f)))
	}
}

func (ø *Initializr) AddScriptHead(js string) {
	// ŧ.Printf("adding js head %v cachedJsHead: %#v\n", js, ø.cachedJsHead)
	ø.cachedJsHead = append(ø.cachedJsHead, directText(js))
	if !ø.UseCachedFiles {
		ø.Head.Add(SCRIPT(Html(js)))
	}
}

func (ø *Initializr) AddScript(js string) {
	// ŧ.Printf("adding js %v cachedJs: %#v\n", js, ø.cachedJs)
	ø.cachedJs = append(ø.cachedJs, directText(js))
	if !ø.UseCachedFiles {
		ø.Body.Add(SCRIPT(Html(js)))
	}
}

func (ø *Initializr) AddGoogleAnalytics(account string) {
	ø.AddScript(`var _gaq=[['_setAccount','` + account + `'],['_trackPageview']];
            (function(d,t){var g=d.createElement(t),s=d.getElementsByTagName(t)[0];
            g.src=('https:'==location.protocol?'//ssl':'//www')+'.google-analytics.com/ga.js';
            s.parentNode.insertBefore(g,s)}(document,'script'));`)
}

func (ø *Initializr) NewContainter() *Element {
	return DIV(Class("container"))
}

func (ø *Initializr) SetupHead() {
	if ø.UseCachedFiles {
		ø.Head.Add(LINK(Attr("rel", "stylesheet"), Attr("href", ø.cssPathed(cachedCssPath))), lf)
		ø.Head.Add(SCRIPT(Attr("src", ø.jsPathed(cachedJsHeadPath))))
	} else {
		ø.AddCssFile(ø.cssPathed("%s/bootstrap.min.css"))
		ø.AddStyle(style)
		ø.AddCssFile(ø.cssPathed("%s/bootstrap-responsive.min.css"))
		//ø.AddCssFile(ø.cssPathed("%s/main.css"))
		ø.Head.Add(SCRIPT(Attr("src", ø.jsPathed("%s/modernizr-2.6.2-respond-1.1.0.min.js"))))
	}
}

func (ø *Initializr) SetupBody() {
	ø.Body.Add(Html(infoChromeFrame))
	ø.Body.Add(Html(infoBootstrap))
	ø.Navigation = ø.NewContainter()
	navBar := DIV(
		Class("navbar"),
		Class("navbar-inverse"),
		Class("navbar-fixed-top"),
		DIV(
			Class("navbar-inner"),
			ø.Navigation))
	ø.Body.Add(navBar)

	ø.Main = ø.NewContainter()
	ø.Body.Add(ø.Main)
	if ø.UseCachedFiles {
		ø.Body.Add(SCRIPT(Attr("src", ø.jsPathed(cachedJsPath))))
	} else {
		//ø.AddScriptFileHead("http://ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js")
		//ø.AddScriptHead(ø.jsPathed(`window.jQuery || document.write('<script src="%s/jquery-1.8.3.min.js"><\/script>')`))
		ø.AddScriptFileHead(ø.jsPathed("%s/jquery-1.8.3.min.js"))
		//ø.AddScriptFile(ø.jsPathed(`%s/jquery-1.8.3.min.js`))
		//ø.AddScriptFile(ø.jsPathed("%s/bootstrap.min.js"))
		ø.AddScriptFile(ø.jsPathed("%s/bootstrap.js"))
		ø.AddScriptFile(ø.jsPathed("%s/main.js"))
	}
}

func Layout() (layout *Initializr) {
	layout = &Initializr{
		Template:     &Template{Element: Doc(Html(pre))},
		pre:          pre,
		post:         post,
		Head:         HEAD(lf, charset, lf, chromeFrame, lf, viewport, lf),
		Body:         BODY(),
		CssPath:      "/css",
		JsPath:       "/js",
		cachedJs:     []Loadable{},
		cachedJsHead: []Loadable{},
		cachedCss:    []Loadable{},
	}
	//layout.Template.Tag = layout.Body
	layout.Template.Add(layout.Head, layout.Body)
	return
}
