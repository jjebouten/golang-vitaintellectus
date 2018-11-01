package main

import (
	"database/sql"
	"log"
	"net/http"
)

type klant struct {
	Klantnummer           sql.NullInt64
	Naam                  sql.NullString
	Voornaam              sql.NullString
	Postcode              sql.NullString
	Huisnummer            sql.NullInt64
	Huisnummer_toevoeging sql.NullString
	Geboortedatum         sql.NullString
	Geslacht              sql.NullString
	Bloedgroep            sql.NullString
	Rhesusfactor          sql.NullString
	Beroepsrisicofactor   sql.NullString
	Inkomen               sql.NullInt64
	Kredietregistratie    sql.NullString
	Opleiding             sql.NullString
	Opmerking             sql.NullString
}

func getKlant(nKlantnummer int) []klant{

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM klant WHERE klantnummer=?", nKlantnummer)
	if err != nil {
		panic(err.Error())
	}
	nKlant := klant{}
	res := []klant{}
	for selDB.Next() {
		var klantnummer sql.NullInt64
		var naam sql.NullString
		var voornaam sql.NullString
		var postcode sql.NullString
		var huisnummer sql.NullInt64
		var huisnummer_toevoeging sql.NullString
		var geboortedatum sql.NullString
		var geslacht sql.NullString
		var bloedgroep sql.NullString
		var rhesusfactor sql.NullString
		var beroepsrisicofactor sql.NullString
		var inkomen sql.NullInt64
		var kredietregistratie sql.NullString
		var opleiding sql.NullString
		var opmerking sql.NullString
		err = selDB.Scan(&klantnummer, &naam, &voornaam, &postcode, &huisnummer, &huisnummer_toevoeging, &geboortedatum, &geslacht, &bloedgroep, &rhesusfactor, &beroepsrisicofactor, &inkomen, &kredietregistratie, &opleiding, &opmerking)
		if err != nil {
			panic(err.Error())
		}
		nKlant.Klantnummer = klantnummer
		nKlant.Naam = naam
		nKlant.Voornaam = voornaam
		nKlant.Postcode = postcode
		nKlant.Huisnummer = huisnummer
		nKlant.Huisnummer_toevoeging = huisnummer_toevoeging
		nKlant.Geboortedatum = geboortedatum
		nKlant.Geslacht = geslacht
		nKlant.Bloedgroep = bloedgroep
		nKlant.Rhesusfactor = rhesusfactor
		nKlant.Beroepsrisicofactor = beroepsrisicofactor
		nKlant.Inkomen = inkomen
		nKlant.Kredietregistratie = kredietregistratie
		nKlant.Opleiding = opleiding
		nKlant.Opmerking = opmerking

		res = append(res, nKlant)

	}

	defer db.Close()
	return res

}

func IndexKlanten(w http.ResponseWriter, r *http.Request) {

	if v := Loginsession; v == 0 {
		http.Redirect(w, r, "/failed", 302)
	}

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM klant ORDER BY klantnummer DESC")
	if err != nil {
		panic(err.Error())
	}
	nKlant := klant{}
	res := []klant{}

	for selDB.Next() {
		var klantnummer sql.NullInt64
		var naam sql.NullString
		var voornaam sql.NullString
		var postcode sql.NullString
		var huisnummer sql.NullInt64
		var huisnummer_toevoeging sql.NullString
		var geboortedatum sql.NullString
		var geslacht sql.NullString
		var bloedgroep sql.NullString
		var rhesusfactor sql.NullString
		var beroepsrisicofactor sql.NullString
		var inkomen sql.NullInt64
		var kredietregistratie sql.NullString
		var opleiding sql.NullString
		var opmerking sql.NullString
		err = selDB.Scan(&klantnummer, &naam, &voornaam, &postcode, &huisnummer, &huisnummer_toevoeging, &geboortedatum, &geslacht, &bloedgroep, &rhesusfactor, &beroepsrisicofactor, &inkomen, &kredietregistratie, &opleiding, &opmerking)
		if err != nil {
			panic(err.Error())
		}
		nKlant.Klantnummer = klantnummer
		nKlant.Naam = naam
		nKlant.Voornaam = voornaam
		nKlant.Postcode = postcode
		nKlant.Huisnummer = huisnummer
		nKlant.Huisnummer_toevoeging = huisnummer_toevoeging
		nKlant.Geboortedatum = geboortedatum
		nKlant.Geslacht = geslacht
		nKlant.Bloedgroep = bloedgroep
		nKlant.Rhesusfactor = rhesusfactor
		nKlant.Beroepsrisicofactor = beroepsrisicofactor
		nKlant.Inkomen = inkomen
		nKlant.Kredietregistratie = kredietregistratie
		nKlant.Opleiding = opleiding
		nKlant.Opmerking = opmerking

		res = append(res, nKlant)
	}
	if err := tmpl.ExecuteTemplate(w, "Indexklanten", res); err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}
