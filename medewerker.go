package main

import (
	"database/sql"
)

type medewerker struct {
	Medewerkersnummer     sql.NullInt64
	Naam                  sql.NullString
	Voorletters           sql.NullString
	Functie               sql.NullString
	Fte                   sql.NullString
	Datum_in_dienst       sql.NullString
	Postcode              sql.NullString
	Huisnummer            sql.NullString
	Huisnummer_toevoeging sql.NullString
}

func getMedewerker(nMedewerkersnummer int) []medewerker {

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM medewerker WHERE medewerkernummer=?", nMedewerkersnummer)
	if err != nil {
		panic(err.Error())
	}
	nMedewerker := medewerker{}
	res := []medewerker{}
	for selDB.Next() {
		var medewerkersnummer sql.NullInt64
		var naam sql.NullString
		var voorletters sql.NullString
		var functie sql.NullString
		var fte sql.NullString
		var datum_in_dienst sql.NullString
		var postcode sql.NullString
		var huisnummer sql.NullString
		var huisnummer_toevoeging sql.NullString
		err = selDB.Scan(&medewerkersnummer ,&naam ,&voorletters ,&functie ,&fte ,&datum_in_dienst ,&postcode ,&huisnummer ,&huisnummer_toevoeging)
		if err != nil {
			panic(err.Error())
		}
		nMedewerker.Medewerkersnummer = medewerkersnummer
		nMedewerker.Naam = naam
		nMedewerker.Voorletters = voorletters
		nMedewerker.Functie = functie
		nMedewerker.Fte = fte
		nMedewerker.Datum_in_dienst = datum_in_dienst
		nMedewerker.Postcode = postcode
		nMedewerker.Huisnummer = huisnummer
		nMedewerker.Huisnummer_toevoeging = huisnummer_toevoeging

		res = append(res, nMedewerker)

	}

	defer db.Close()
	return res

}
