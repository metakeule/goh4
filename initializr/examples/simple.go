package main

import (
	ŧ "fmt"
	. "gopkg.in/metakeule/goh4.v5/initializr"
)

func main() {
	ø := Layout()
	ø.AddGoogleAnalytics("mystring")
	SetupExampleNavigation(ø)
	SetupExampleMain(ø)
	ŧ.Println(ø)
}
