package initializr

import (
	. "gopkg.in/metakeule/goh4.v5"
	. "gopkg.in/metakeule/goh4.v5/tag"
)

func exampleLiA(href string, text string) *Element {
	return LI(A(text, Attr("href", href)))
}

func SetupExampleNavigation(ø *Initializr) {
	btn := A(
		Class("btn"),
		Class("btn-navbar"),
		Attr("data-toggle", "collapse"),
		Attr("data-target", ".nav-collapse"))

	btn.Add(SPAN(Class("icon-bar")))
	btn.Add(SPAN(Class("icon-bar")))
	btn.Add(SPAN(Class("icon-bar")))

	dropDownMenu := UL(Class("dropdown-menu"))
	dropDownMenu.Add(exampleLiA("#", "Action"))
	dropDownMenu.Add(exampleLiA("#", "Another action"))
	dropDownMenu.Add(exampleLiA("#", "Something else here"))
	dropDownMenu.Add(LI(Class("divider")))
	dropDownMenu.Add(LI(Class("nav-header"), "Nav header"))
	dropDownMenu.Add(exampleLiA("#", "Separated link"))
	dropDownMenu.Add(exampleLiA("#", "One more separated link"))

	dropDown := LI(Class("dropdown"))

	dropdownToggle := A(
		Class("dropdown-toggle"),
		Attr("data-toggle", "dropdown"))

	dropdownToggle.Add("Dropdown", B(Class("caret")))
	dropDown.Add(dropdownToggle)
	dropDown.Add(dropDownMenu)

	navUl := UL(Class("nav"))
	navUl.Add(LI(Class("active"), A(Attr("href", "#"), "Home")))
	navUl.Add(exampleLiA("#about", "About"))
	navUl.Add(exampleLiA("#contact", "Contact"))
	navUl.Add(dropDown)

	form := FORM(Class("navbar-form"), Class("pull-right"))
	form.Add(INPUT(Class("span2"), Attr("type", "text"), Attr("placeholder", "Email")))
	form.Add(INPUT(Class("span2"), Attr("type", "password"), Attr("placeholder", "Password")))
	form.Add(BUTTON(Class("btn"), Attr("type", "submit"), "Sign in"))

	navCollapse := DIV(Class("nav-collapse"), Class("collapse"))
	navCollapse.Add(navUl)
	navCollapse.Add(form)

	ø.Navigation.SetContent() // reset the navigation
	ø.Navigation.Add(btn)
	ø.Navigation.Add(A(Class("brand"), Attr("href", "#"), "Project name"))
	ø.Navigation.Add(navCollapse)
}
