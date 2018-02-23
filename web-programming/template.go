package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/home/abhishek/GolandProjects/Practice/web-programming/template/*"))
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	err := tpl.Execute(os.Stdout, nil)
	check(err)
	err = tpl.ExecuteTemplate(os.Stdout, "a.gohtml", nil)
	check(err)
	err = tpl.ExecuteTemplate(os.Stdout, "b.gohtml", nil)
	check(err)
	err = tpl.ExecuteTemplate(os.Stdout, "c.gohtml", nil)
	check(err)
}