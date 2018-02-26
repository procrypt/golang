package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}


func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name: "abhishek",
		Value: "Abhishek",
	})
	fmt.Fprintln(w, "Cookie Written.")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("abhishek")
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "Cookie: #1", c1)
	}

	c2, err := req.Cookie("nyc")
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "Cookie: #2", c2)
	}

	c3, err := req.Cookie("nyu")
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "Cookie: #3", c3)
	}

	c4, err := req.Cookie("bobby")
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "Cookie: #4", c4)
	}

}

func abundance(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name: "nyc",
		Value: "New York City",
	})

	http.SetCookie(w, &http.Cookie{
		Name: "nyu",
		Value: "New York University",
	})

	http.SetCookie(w, &http.Cookie{
		Name: "bobby",
		Value: "Jaipur",
	})

}