package main

import (
	"flag"
	"fmt"
	fr "github.com/metakeule/fastreplace"
	"github.com/metakeule/goh4"
	"io/ioutil"
	"koelnart/app/backend/class"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	c2 = `
export function Artist()  { return ".artist"; } 
export function ArtistName()  { return ".artist-name"; }
`
)

/*
func createFile(path string) error {
    class.
}
*/
