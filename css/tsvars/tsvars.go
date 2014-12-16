package main

import (
	"flag"
	"fmt"
	fr "github.com/metakeule/fastreplace"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	content = `
package main

import (
    "strings"
    "io/ioutil"
    "fmt"
    "gopkg.in/metakeule/goh4.v5"
    "@@package@@"
    "os"
    "time"
)

var preClass = "module Class {\n\t"
var postClass = "}\n\nexport = Class;\n"
var preId = "module Id {\n\t"
var postId = "}\n\nexport = Id;\n"

func CompileClasses(classes []goh4.Class, file string, info string) error {
    ts := []string{"/* " + info + " */", preClass}
    for _, c := range classes {
        name := string(c)
        name = strings.Replace(name, "-", "_", -1)
        ts = append(ts,  fmt.Sprintf("\texport var %s = '%s';", strings.ToUpper(name[0:1])+strings.ToLower(name)[1:], string(c) ))
        ts = append(ts,  fmt.Sprintf("\texport var $%s = '.%s';", strings.ToUpper(name[0:1])+strings.ToLower(name)[1:], string(c) ))
    }
    ts = append(ts, postClass)
    return ioutil.WriteFile(file, []byte(strings.Join(ts, "\n")), os.FileMode(0644))
}

func CompileIds(ids []goh4.Id, file string, info string) error {
    ts := []string{"/* " + info + " */", preId}
    for _, c := range ids {
        name := string(c)
        name = strings.Replace(name, "-", "_", -1)
        ts = append(ts,  fmt.Sprintf("\texport var %s = '%s';", strings.ToUpper(name[0:1])+strings.ToLower(name)[1:], string(c) ))
        ts = append(ts,  fmt.Sprintf("\texport var $%s = '#%s';", strings.ToUpper(name[0:1])+strings.ToLower(name)[1:], string(c) ))
    }
    ts = append(ts, postId)
    return ioutil.WriteFile(file, []byte(strings.Join(ts, "\n")), os.FileMode(0644))
}

func main() {
    target := "@@target@@"
    info := fmt.Sprintf("\n  DO NOT EDIT THIS FILE!\n\n  It was automatically generated from\n\n    @@package@@\n\n  at\n\n    %s\n\n  and will be overwritten. ", time.Now())
    err := @@method@@(@@type@@.@@var@@, target, info)
    if err != nil {
        fmt.Printf("Error while compiling @@package@@ to %s:\n%s", target, err.Error())
        os.Exit(1)
    }
}
`
	target    = flag.String("target", "", "output file to be written to (less or scss file)")
	in        = flag.String("in", "", "input go package that either has the name class and contains a var ALL []goh4.Class with all classes or has the name id and contains a var ALL []goh4.Id with all ids")
	var_      = flag.String("var", "ALL", "name of the variable in package")
	body, _   = fr.NewString("@@", content)
	type_     string
	method    string
	regString = "^[A-Z]([a-zA-Z])+$"
	matcher   = regexp.MustCompile(regString)
)

func Replace() []byte {
	i := body.Instance()
	i.AssignString("package", *in)
	i.AssignString("target", *target)
	i.AssignString("method", method)
	i.AssignString("type", type_)
	i.AssignString("var", *var_)
	return i.Bytes()
}

func Exec(name string, opts ...string) (output string, err error) {
	cmd := exec.Command(name, opts...)
	var out []byte
	out, err = cmd.CombinedOutput()
	output = string(out)
	if err != nil {
		err = fmt.Errorf(output)
	}
	return
}

func Which(cmd string) (path string, err error) {
	path, err = Exec("which", cmd)
	path = strings.TrimRight(path, "\n")
	if path == "" {
		err = fmt.Errorf("not found: %s", cmd)
	}
	return
}

func main() {

	flag.Parse()

	ext := strings.ToLower(filepath.Ext(*target))
	if ext != ".ts" {
		fmt.Printf("invalid target file: %#v, should be of type .ts\n", *target)
		os.Exit(1)
	}

	i := strings.LastIndex(*in, "/")
	if i < 1 {
		fmt.Printf("invalid input package name %#v\n", *in)
		os.Exit(1)
	}

	type_ = (*in)[i+1:]

	switch type_ {
	case "class":
		method = "CompileClasses"
	case "id":
		method = "CompileIds"
	default:
		fmt.Printf("invalid input package %#v, should be class or id\n", *in)
		os.Exit(1)
	}
	if matcher.MatchString(*var_) == false {
		fmt.Printf("invalid var name %#v, must match %s\n", *var_, regString)
		os.Exit(1)
	}

	//func TempDir(dir, prefix string) (name string, err error)
	dir, err := ioutil.TempDir("/tmp", "classidtotscompiler")
	if err != nil {
		fmt.Println("Can't create tempdir: ", err.Error())
		os.Exit(1)
	}
	defer func() {
		os.RemoveAll(dir)
	}()

	file := path.Join(dir, "main.go")
	err = ioutil.WriteFile(file, Replace(), os.FileMode(0644))
	if err != nil {
		fmt.Println("Can't create tempfile: ", err.Error())
		os.Exit(1)
	}

	goBin, err := Which("go")
	if err != nil {
		fmt.Println("can't find go compiler in $PATH")
		os.Exit(1)
	}

	var out string
	out, err = Exec(goBin, "run", file)
	if err != nil {
		fmt.Println("Error while compliling:\n", out)
		os.Exit(1)
	}

}
