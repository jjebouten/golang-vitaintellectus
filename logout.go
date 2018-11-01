package main

import (
	"log"
	"net/http"
)


func Logout(w http.ResponseWriter, r *http.Request) {
	Loginsession = 0

	log.Println(Loginsession)
	http.Redirect(w, r, "/Loggedout", 302)
}
