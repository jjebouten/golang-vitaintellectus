package main

import (
	"net/http"
)


func Logout(w http.ResponseWriter, r *http.Request) {
	Loginsession = 0
	http.Redirect(w, r, "/Loggedout", 302)
}
