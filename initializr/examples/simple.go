package main

import (
	ŧ "fmt"
	. "github.com/metakeule/goh4/initializr"
)

func main() {
	ø := Layout()
	ø.AddGoogleAnalytics("mystring")
	SetupExampleNavigation(ø)
	SetupExampleMain(ø)
	ŧ.Println(ø)
}
