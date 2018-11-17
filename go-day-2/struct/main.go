package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	Muhammed := sage{
		Name:  "Muhammed",
		Motto: "Only Allah",
	}
	err := tpl.Execute(os.Stdout, Muhammed)
	if err != nil {
		log.Fatalln(err)
	}
}
