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
	Klantinfo []klant
	Medewerkerinfo []medewerker
}



func IndexBestelling(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM bestelling ORDER BY bestelnummer DESC")
	if err != nil {
		panic(err.Error())
	}
	order := bestelling{}
	res := []bestelling{}

	for selDB.Next() {
		var bestelnummer int
		var status string
		var besteldatum string
		var afbetaling_doorlooptijd int
		var afbetaling_maandbedrag float64
		var klantnummer int
		var verkoper int
		err = selDB.Scan(&bestelnummer, &status, &besteldatum, &afbetaling_doorlooptijd, &afbetaling_maandbedrag, &klantnummer, &verkoper)
		if err != nil {
			panic(err.Error())
		}
		order.Bestelnummer = bestelnummer

		order.Status = status
		order.Besteldatum = besteldatum
		order.Afbetaling_doorlooptijd = afbetaling_doorlooptijd
		order.Afbetaling_maandbedrag = afbetaling_maandbedrag
		order.Klantnummer = klantnummer
		order.Verkoper = verkoper

		res = append(res, order)
	}
	if err := tmpl.ExecuteTemplate(w, "IndexBestelling", res); err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

func BekijkBestelling(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nBestelnummer := r.URL.Query().Get("bestelnummer")
	selDB, err := db.Query("SELECT klantnummer, verkoper FROM bestelling WHERE bestelnummer=?", nBestelnummer)
	if err != nil {
		panic(err.Error())
	}



	klantinfo :=[]klant{}
	medewerkerinfo :=[]medewerker{}

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
