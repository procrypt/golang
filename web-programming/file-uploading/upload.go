package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"os"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))

}

func foo(w http.ResponseWriter, req *http.Request) {

	//var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		// Open the file
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//FYI
		fmt.Println("\nFile", f, "\nHeader", h, "\nError", err)

		// Read the file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//s = string(bs)

		// Store the file
		file, err := os.Create(h.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		_, err = file.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
			 <input type="file" name="q">
             <input type="submit">
             </form>
             <br>`)
}

