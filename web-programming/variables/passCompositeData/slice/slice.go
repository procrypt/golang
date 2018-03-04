package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("/home/abhishek/GolandProjects/Practice/web-programming/variables/passCompositeData/slice/slice.gohtml"))

}

func main() {

	sage := []string{"Abhishek","Pratap","Singh"}

	tpl.Execute(os.Stdout, sage)
}
