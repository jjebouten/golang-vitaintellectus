package main

import (

	"net/http"
)

var Loginsession = 0
var Medewerkersnummer = int64(0)

func Login(w http.ResponseWriter, r *http.Request) {

	if Loginsession == 1 {
		http.Redirect(w, r, "/indexbestelling", 302)
	} else {
		tmpl.ExecuteTemplate(w, "Login", nil)
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Pusername := r.FormValue("username")
		Ppassword := r.FormValue("password")


		selDB, err := db.Query("SELECT medewerkernummer, naam, datum_in_dienst FROM medewerker WHERE naam=? AND datum_in_dienst=?", Pusername, Ppassword)
		if err != nil {
			panic(err.Error())
		}

		nMedewerker := medewerker{}
		for selDB.Next() {
			err = selDB.Scan(&nMedewerker.Medewerkersnummer, &nMedewerker.Naam, &nMedewerker.Datum_in_dienst)
			if err != nil {
				panic(err.Error())
			}

			Loginsession = 1
			Medewerkersnummer = nMedewerker.Medewerkersnummer.Int64
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