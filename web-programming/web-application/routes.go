package main

import (
	"net/http"
)
func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func uploadView(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "uploadView.gohtml", nil)
}
