package main

import (
	"sync"
	"net/http"
	"log"
	"fmt"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", hander)
	http.HandleFunc("/", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func hander(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
