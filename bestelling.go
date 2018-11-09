package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type bestelling struct {
	Bestelnummer            int
	Status                  string
	Besteldatum             string
	Afbetaling_doorlooptijd int
	Afbetaling_maandbedrag  float64
	Klantnummer             int
	Verkoper                int
}

type data struct {
	Klantinfo      []klant
	Medewerkerinfo []medewerker
}

func IndexBestelling(w http.ResponseWriter, r *http.Request) {

	if v := Loginsession; v == 0 {
		http.Redirect(w, r, "/failed", 302)
	}

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM bestelling ORDER BY bestelnummer DESC")

	if err != nil {
		panic(err.Error())
	}
	nOrder := bestelling{}
	res := []bestelling{}

	for selDB.Next() {
		err = selDB.Scan(&nOrder.Bestelnummer, &nOrder.Status, &nOrder.Besteldatum, &nOrder.Afbetaling_doorlooptijd,
			&nOrder.Afbetaling_maandbedrag, &nOrder.Klantnummer, &nOrder.Verkoper)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, nOrder)
	}
	if err := tmpl.ExecuteTemplate(w, "IndexBestelling", res); err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

func BekijkBestelling(w http.ResponseWriter, r *http.Request) {

	if v := Loginsession; v == 0 {
		http.Redirect(w, r, "/failed", 302)
	}

	db := dbConn()
	nBestelnummer := r.URL.Query().Get("bestelnummer")
	selDB, err := db.Query("SELECT klantnummer, verkoper FROM bestelling WHERE bestelnummer=?", nBestelnummer)
	if err != nil {
		panic(err.Error())
	}

	klantinfo := []klant{}
	medewerkerinfo := []medewerker{}
	data := data{}

	for selDB.Next() {
		var klantnummer, verkoper int
		err = selDB.Scan(&klantnummer, &verkoper)
		if err != nil {
			panic(err.Error())
		}

		klantinfo = getKlant(klantnummer)
		medewerkerinfo = getMedewerker(verkoper)

		data.Klantinfo = klantinfo
		data.Medewerkerinfo = medewerkerinfo
	}

	if err := tmpl.ExecuteTemplate(w, "Bekijkbestelling", data); err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
}

func ServeNewBestelling(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewBestelling", nil)
}
