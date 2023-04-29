package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t := template.Must(template.New("").ParseGlob("tmp/*.tmpl"))
	if err:= t.ExecuteTemplate(os.Stdout, "index", "これは本文です"); err!=nil{
		log.Println(err)
	}
}
