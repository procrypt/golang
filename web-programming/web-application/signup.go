package main

import (
	"net/http"
	"github.com/satori/go.uuid"
)

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLogged(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		uname  := req.FormValue("UserName")
		fname  := req.FormValue("FirstName")
		lname  := req.FormValue("LastName")
		passwd := req.FormValue("Password")

		// populate the user struct
		usr := user{uname, fname, lname, passwd}

		// userName already taken
		if _, ok := dbUsers[uname]; ok {
			http.Error(w,"User Name already taken!", http.StatusForbidden)
			//http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}

		// create  sessionID
		sessionID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:     "session-id",
			Value:    sessionID.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		// store the session key with username as value in dbSession Map
		dbSessions[cookie.Value] = uname
		// store the username key with user struct as value in dbUsers Map
		dbUsers[uname] = usr

		// Redirect to the login page after the signup process to login
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Execute the template
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func getUserInfo(w http.ResponseWriter, req *http.Request) user {
	// get Cookie
	cookie, err := req.Cookie("session-id")
	if err != nil {
		sessionID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:     "Session-id",
			Value:    sessionID.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	// if user exists
	var usr user
	if uname, ok := dbSessions[cookie.Value]; ok {
		usr = dbUsers[uname]
	}
	return usr
}

func alreadyLogged(req *http.Request) bool {
	// get cookie
	cookie, err := req.Cookie("session-id")
	if err!= nil {
		return false
	}
	 _, ok := dbSessions[cookie.Value]
	 return ok
}
