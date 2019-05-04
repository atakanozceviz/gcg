package cmd

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"text/template"
)

func executeTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(template.FuncMap{"isPR": isPR}).Parse(tmplStr)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	return buf.String(), err
}

var re = regexp.MustCompile("https://github.com/\\w+/\\w+/pull/\\d+")

func isPR(url string) bool {
	if re.MatchString(url) {
		return true
	}
	return false
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
