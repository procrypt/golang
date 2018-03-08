package main

import "net/http"

func logout(w http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("session-id")

	//user associated with the cookie
	user := dbSessions[cookie.Value]

	// delete the session
	delete(dbSessions, user)

	//create a new cookie with negative MaxAge
	newCookie := &http.Cookie{
		Name: "session-id",
		Value: "random",
		MaxAge: -1,
	}
	http.SetCookie(w, newCookie)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
