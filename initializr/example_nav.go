package initializr

import (
	. "github.com/metakeule/goh4"
)

func exampleLiA(href string, text string) *Element {
	return Li(A(text, Attr("href", href)))
}

func SetupExampleNavigation(ø *Initializr) {
	btn := A(
		Class("btn"),
		Class("btn-navbar"),
		Attr("data-toggle", "collapse"),
		Attr("data-target", ".nav-collapse"))

	btn.Add(Span(Class("icon-bar")))
	btn.Add(Span(Class("icon-bar")))
	btn.Add(Span(Class("icon-bar")))

	dropDownMenu := Ul(Class("dropdown-menu"))
	dropDownMenu.Add(exampleLiA("#", "Action"))
	dropDownMenu.Add(exampleLiA("#", "Another action"))
	dropDownMenu.Add(exampleLiA("#", "Something else here"))
	dropDownMenu.Add(Li(Class("divider")))
	dropDownMenu.Add(Li(Class("nav-header"), "Nav header"))
	dropDownMenu.Add(exampleLiA("#", "Separated link"))
	dropDownMenu.Add(exampleLiA("#", "One more separated link"))

	dropDown := Li(Class("dropdown"))

	dropdownToggle := A(
		Class("dropdown-toggle"),
		Attr("data-toggle", "dropdown"))

	dropdownToggle.Add("Dropdown", B(Class("caret")))
	dropDown.Add(dropdownToggle)
	dropDown.Add(dropDownMenu)

	navUl := Ul(Class("nav"))
	navUl.Add(Li(Class("active"), A(Attr("href", "#"), "Home")))
	navUl.Add(exampleLiA("#about", "About"))
	navUl.Add(exampleLiA("#contact", "Contact"))
	navUl.Add(dropDown)

	form := Form(Class("navbar-form"), Class("pull-right"))
	form.Add(Input(Class("span2"), Attr("type", "text"), Attr("placeholder", "Email")))
	form.Add(Input(Class("span2"), Attr("type", "password"), Attr("placeholder", "Password")))
	form.Add(Button(Class("btn"), Attr("type", "submit"), "Sign in"))

	navCollapse := Div(Class("nav-collapse"), Class("collapse"))
	navCollapse.Add(navUl)
	navCollapse.Add(form)

	ø.Navigation.Set() // reset the navigation
	ø.Navigation.Add(btn)
	ø.Navigation.Add(A(Class("brand"), Attr("href", "#"), "Project name"))
	ø.Navigation.Add(navCollapse)
}
