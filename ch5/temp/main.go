package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name  string
	Names []string
	Age   int
}

func main() {
	tempFunc()
}

func tempFunc() {
	user := User{
		Name:  "Bob",
		Names: []string{"Hello", "World"},
		Age:   10,
	}
	tmpl := `{{.Name}}
	{{range .Names}}
	<p>{{.}}</p>{{end}}
	{{index .Names}}
	{{if gt .Age 10}}
	{{.Name}} is older than 10
	{{else}}
	{{.Name}} is not older than 10
	{{end}}
	`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}

}
