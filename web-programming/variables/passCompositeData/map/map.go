package main

import (
	"html/template"
	"os"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("/home/abhishek/GolandProjects/Practice/web-programming/variables/passCompositeData/map/map.gohtml"))
}


func main() {

	sage := map[string]string{
		"Abhsihek": "NYU",
		"Bobby":    "NYC",
	}

	tpl.Execute(os.Stdout, sage)

}
