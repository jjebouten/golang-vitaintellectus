package main

import (
	"log"
	"net/http"
)

type module struct {
	Modulenaam   string
	Omschrijving string
	Stukprijs    string
}


func IndexModules(w http.ResponseWriter, r *http.Request) {

	if v := Loginsession; v == 0 {
		http.Redirect(w, r, "/failed", 302)
	}

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM module ORDER BY modulenaam DESC")
	if err != nil {
		panic(err.Error())
	}
	nModule := module{}
	res := []module{}

	for selDB.Next() {
		err = selDB.Scan(&nModule.Modulenaam, &nModule.Omschrijving, &nModule.Stukprijs)

		if err != nil {
			panic(err.Error())
		}
		res = append(res, nModule)
	}
	if err := tmpl.ExecuteTemplate(w, "Indexmodules", res); err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}
