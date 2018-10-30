package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type bestelling struct {
	bestelnummer int
	status  string
	besteldatum string
	afbetaling_doorlooptijd int
	afbetaling_maandbedrag float64
	klantnummer int
	verkoper int
}

var tmpl = template.Must(template.ParseGlob("form/*"))

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
		var status  string
		var besteldatum string
		var afbetaling_doorlooptijd int
		var afbetaling_maandbedrag float64
		var klantnummer int
		var verkoper int
		err = selDB.Scan(&bestelnummer,&status,&besteldatum,&afbetaling_doorlooptijd,&afbetaling_maandbedrag,&klantnummer,&verkoper)
		if err != nil {
			panic(err.Error())
		}
		order.bestelnummer = bestelnummer

		order.status = status
		order.besteldatum = besteldatum
		order.afbetaling_doorlooptijd = afbetaling_doorlooptijd
		order.afbetaling_maandbedrag = afbetaling_maandbedrag
		order.klantnummer = klantnummer
		order.verkoper = verkoper

		res = append(res, order)
	}
	if err := tmpl.ExecuteTemplate(w, "IndexBestelling", res); err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

//func Show(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Show", emp)
//	defer db.Close()
//}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//func Edit(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Edit", emp)
//	defer db.Close()
//}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 302)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 302)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 302)
}

func main() {
	log.Println("Server started on: http://localhost:8898")
	//http.HandleFunc("/", Login)
	//http.HandleFunc("/show", Show)
	//http.HandleFunc("/index", Index)
	http.HandleFunc("/", IndexBestelling)
	http.HandleFunc("/new", New)
	//http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/authenticate", Authenticate)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8898", nil)
}