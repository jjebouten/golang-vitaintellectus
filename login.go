package main

import (
	"database/sql"
	"net/http"
)

var Loginsession = 0

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Login", nil)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Pusername := r.FormValue("username")
		Ppassword := r.FormValue("password")

		selDB, err := db.Query("SELECT naam, datum_in_dienst FROM medewerker WHERE naam=? AND datum_in_dienst=?", Pusername, Ppassword)
		if err != nil {
			panic(err.Error())
		}

		medewerker := medewerker{}
		for selDB.Next() {
			var naam, datum_in_dienst sql.NullString
			err = selDB.Scan(&naam, &datum_in_dienst)
			if err != nil {
				panic(err.Error())
			}
			medewerker.Naam = naam
			medewerker.Datum_in_dienst = datum_in_dienst

			Loginsession = 1

		}

	}
	defer db.Close()

	if v := Loginsession; v == 1 {
		http.Redirect(w, r, "/indexbestelling", 302)
	} else {
		http.Redirect(w, r, "/failed", 302)
	}
}

func Failed(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "Failed", nil)
}