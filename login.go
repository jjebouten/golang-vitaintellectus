package main

import (
	"log"
	"net/http"
)

type user struct {
	uid      int
	username string
	role     int
}


func Login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Login", nil)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Pusername := r.FormValue("username")
		Ppassword := r.FormValue("password")

		selDB, err := db.Query("SELECT uid, username, password, role FROM user WHERE username=? AND password=?", Pusername, Ppassword)
		if err != nil {
			panic(err.Error())
		}

		user := user{}
		for selDB.Next() {
			var uid, role int
			var username, password string
			err = selDB.Scan(&uid, &username, &password, &role)
			if err != nil {
				panic(err.Error())
			}
			user.uid = uid
			user.username = username
			user.role = role
			log.Println(user)
		}

	}


	defer db.Close()
	http.Redirect(w, r, "/", 302)
}