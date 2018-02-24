package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func init() {
	tpl = template.Must(template.ParseGlob("/home/abhishek/GolandProjects/Practice/web-programming/template/*"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "passData.html", "Release self-focus; embrace other-focus.")
	check(err)
}