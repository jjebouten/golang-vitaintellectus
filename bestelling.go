package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"net/http"
	"strconv"
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
	Moduleinfo     []module
	Bestellinginfo []bestelling
	Besteldata     []besteldata
}

type besteldata struct {
	Totalekosten    float64
	Totaleopbrengst float64
	Betaalbaar      float64
	Doorlooptijd    int
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

	data := data{}

	data.Klantinfo = getExistingKlanten()
	//medewerkerid = setidbijinloggen?()
	data.Moduleinfo = getExistingModules()

	if err := tmpl.ExecuteTemplate(w, "NewBestelling", data); err != nil {
		log.Fatalln(err)
	}
}

func getExistingKlanten() []klant {

	db := dbConn()
	selDB, err := db.Query("SELECT klantnummer, naam FROM klant ORDER BY klantnummer DESC")
	if err != nil {
		panic(err.Error())
	}
	nKlant := klant{}
	res := []klant{}

	for selDB.Next() {
		err = selDB.Scan(&nKlant.Klantnummer, &nKlant.Naam)

		if err != nil {
			panic(err.Error())
		}
		res = append(res, nKlant)
	}

	defer db.Close()
	return res
}

func getExistingModules() []module {

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM module WHERE modulenaam not like 'basis'")
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
	defer db.Close()
	return res
}

func NieuweBestelling(w http.ResponseWriter, r *http.Request) {

	var price []string

	price = append(price, r.FormValue("basis"))
	price = append(price, r.FormValue("Cor"))
	price = append(price, r.FormValue("Dermal"))
	price = append(price, r.FormValue("Memoria"))
	price = append(price, r.FormValue("Oculus"))
	price = append(price, r.FormValue("Oricula"))
	price = append(price, r.FormValue("Pes"))
	price = append(price, r.FormValue("Sanguis"))
	price = append(price, r.FormValue("Somnus"))

	var numbers []float64

	for _, elem := range price {
		i, err := strconv.ParseFloat(elem, 64)
		if err == nil {
			numbers = append(numbers, i)
		}
	}

	//Count all number of selected modules
	totalprice := 0.00
	for i := 0; i < len(numbers); i++ {
		totalprice = totalprice + numbers[i]
	}

	var klantinfo = []klant{}

	klantnummer, err := strconv.Atoi(r.FormValue("Klantnummer"))
	if err == nil {
		klantinfo = getKlant(klantnummer)
	}

	//AGE hebben we

	//Totaal bedrag hebben we

	//basis doorlooptijd =

	doorlooptijd := Doorlooptijd(klantinfo)

	//Formule

	//maandbedrag Z = inkomen per jaar (x0.2) X / doorlooptijd in maanden Y

	//totalprice = basismodule A + (optionele toevoegingen B t/m H)

	//als totaalbedrag Q < (maandbedrag Z x doorlooptijd in maanden Y) = verkoop mag doorgaan.

	Voldoetaaneisen(klantinfo, totalprice, doorlooptijd, w, r)
}

func Doorlooptijd(klantinfo []klant) int {

	var doorlooptijd = 120

	age := getAge(klantinfo)

	if age >= 45 && age <= 55 {
		doorlooptijd = 90
	}
	if age > 55 {
		doorlooptijd = 55
	}

	//aftrek factor is 3 maanden x het beroepsrisico factor
	var aftrekfactor int

	beroepsrisico, err := strconv.Atoi((klantinfo[0].Beroepsrisicofactor.String))
	if err == nil {
		aftrekfactor = beroepsrisico * 3
	}

	doorlooptijd = doorlooptijd - aftrekfactor

	//if kredietregistratie is true dan is doorlooptijd x 1.5 na de aftrek factor
	if (klantinfo[0].Kredietregistratie.String == "J") {
		doorlooptijd = doorlooptijd * 3
		doorlooptijd = doorlooptijd / 2
	}

	return doorlooptijd

}

func Voldoetaaneisen(klantinfo []klant, totalekosten float64, doorlooptijd int, w http.ResponseWriter, r *http.Request) {

	if doorlooptijd < 3 {
		if err := tmpl.ExecuteTemplate(w, "Doorlooptijd", doorlooptijd); err != nil {
			log.Fatalln(err)
		}
	}

	inkomen := klantinfo[0].Inkomen.Int64

	inkomenx := float64(inkomen)
	doorlooptijdx := float64(doorlooptijd)

	inkomenx = inkomenx * 0.2

	maandbedrag := inkomenx / doorlooptijdx

	totaleopbreng := maandbedrag * doorlooptijdx

	betaalbaar := totalekosten - totaleopbreng
	betaalbaar = (math.Floor(betaalbaar*100) / 100) //(round down 2 digs)
	totaleopbreng = (math.Floor(totaleopbreng*100) / 100) //(round down 2 digs)

	Besteldatax := besteldata{}

	data := data{}

	Besteldatax.Betaalbaar = betaalbaar
	Besteldatax.Totaleopbrengst = totaleopbreng
	Besteldatax.Totalekosten = totalekosten
	Besteldatax.Doorlooptijd = doorlooptijd

	data.Besteldata = append(data.Besteldata, Besteldatax)

	data.Klantinfo = klantinfo

	//20% van inkomen
	if totalekosten > totaleopbreng {
		if err := tmpl.ExecuteTemplate(w, "Totaleopbreng", Besteldatax); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := tmpl.ExecuteTemplate(w, "Bestellingpreview", data); err != nil {
			log.Fatalln(err)
		}
	}

	//TODO als het mag de bestelling opslaan en nog doen incombinatie met nieuwe klant

}

func Besteldatapreview(data data) {

}
