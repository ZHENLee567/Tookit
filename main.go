package main

import (
	"os"
	"strings"
	"text/template"
)

//var name string
//
//func main() {
//	err := cmd.Execute()
//	if err != nil {
//		log.Fatalf("cmd.Execute err: %v", err)
//	}
//}

const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl := template.New("go-programming-tour")
	tpl, _ = tpl.Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)
}
