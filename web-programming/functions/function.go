package main

import (
	"html/template"
	"strings"
	"os"
	"fmt"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl =template.Must(template.New("").Funcs(fm).ParseFiles("/home/abhishek/GolandProjects/Practice/web-programming/functions/function.gohtml"))
}

type sage struct {
	Name string
	Motto string
}

func main() {

	data := sage {
		Name:  "Abhishek",
		Motto: "Get job @ Google!",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "function.gohtml", data)
	if err != nil {
		fmt.Println(err)
	}
}