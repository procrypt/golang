package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("/home/abhishek/GolandProjects/Practice/web-programming/template/passCompositeData.gohtml"))
}

func main() {

	data := sage{
		Name: "Abhishek",
		Motto: "Get a job @ Google!",
	}

	tpl.Execute(os.Stdout, data)
}
