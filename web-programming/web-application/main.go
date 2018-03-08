package main

import (
	"html/template"
	"net/http"
	"log"
)

type user struct {
	UserName string
	First string
	Last string
	Password string
}


var (
// variable of Type pointer to template.Template
 tpl *template.Template
// map to store the SessionInfo with SessionId as Key as UserName as Value
 dbSessions  = make(map[string]string)
// map to store the userInfo with UserName as Key and User Struct as Value
 dbUsers = make(map[string]user)
// Function for HTML template
 fm = template.FuncMap{
 	"sp": splitString,
 }
)
// Init function to initialize the template
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("template/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/uploadview", uploadView)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/view", view)
	http.Handle("/pics/", http.FileServer(http.Dir(".")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
