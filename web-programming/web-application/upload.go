package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"os"
	"log"
	"path/filepath"
	"io"
	"strings"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session-id")
	if err!= nil {
		sessionId,_ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: sessionId.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func changeValue(w http.ResponseWriter, cookie *http.Cookie, fname string) *http.Cookie {
	str := cookie.Value
	if !strings.Contains(str, fname) {
		str += "|" + fname
	}
	cookie.Value = str
	http.SetCookie(w, cookie)
	return cookie
}

func upload(w http.ResponseWriter, req *http.Request) {
	if !alreadyLogged(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	cookie := getCookie(w, req)
	// create a new folder with username
	_, err := os.Stat(filepath.Join("pics", dbSessions[cookie.Name]))
	if !os.IsNotExist(err) {
		usr := dbSessions[cookie.Value]
		dirName := filepath.Join("pics", usr)
		err := os.MkdirAll(dirName, 0777)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if req.Method == http.MethodPost {
		mf, mfh, err := req.FormFile("picture")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer mf.Close()

		// create SHA for each uploaded file
		ext := strings.Split(mfh.Filename, ".")[1]
		if ext == "png" || ext == "jpg" || ext == "gif" {
			h := sha1.New()
			io.Copy(h, mf)
			fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

			usr := dbSessions[cookie.Value]
			path := filepath.Join("pics", usr, fname)
			fmt.Println(path, usr, cookie.Value)
			file, err := os.Create(path)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()
			mf.Seek(0, 0)
			io.Copy(file, mf)
			cookie = changeValue(w, cookie, fname)
		}
	}
	//xs := strings.Split(cookie.Value, "|")
	//_, mfh, _ := req.FormFile("picture")
	//fileName := mfh.Filename
	//fmt.Println(fileName)
	tpl.ExecuteTemplate(w, "upload.gohtml", nil)
}

func view (w http.ResponseWriter, req *http.Request) {
	if !alreadyLogged(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	cookie, _ := req.Cookie("session-id")
	usr := dbSessions[cookie.Value]
	dirName := "pics/" + usr
	var s []string
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		a := usr + "/" + f.Name()
		s = append(s, a)
	}

	// process deletion
	if req.Method == http.MethodPost {
		fname := req.FormValue("File")
		usr := dbSessions[cookie.Value]
		path := filepath.Join("pics", usr, fname)
		cmdName := "rm"
		err := exec.Command(cmdName, path).Run()
		if err!=nil{
			fmt.Println(err)
		}
	}
	tpl.ExecuteTemplate(w, "view.gohtml", s)
}

func splitString(s string) string {
	slice :=  strings.Split(s, "/")
	s = string(slice[1])
	return s
}

