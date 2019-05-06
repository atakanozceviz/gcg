package cmd

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func executeTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	return buf.String(), err
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
