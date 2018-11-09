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
		err = selDB.Scan(&nMedewerker.Medewerkersnummer, &nMedewerker.Naam, &nMedewerker.Voorletters, &nMedewerker.Functie,
			&nMedewerker.Fte, &nMedewerker.Datum_in_dienst, &nMedewerker.Postcode, &nMedewerker.Huisnummer,
			&nMedewerker.Huisnummer_toevoeging)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, nMedewerker)

	}

	defer db.Close()
	return res

}
