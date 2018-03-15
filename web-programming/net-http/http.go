package main

import (
	"net/http"
	"io"
)

type hotdog int

func(m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {

	var d hotdog
	http.ListenAndServe(":8080", d)

}
