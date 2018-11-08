package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

//set templates
var tmpl = template.Must(template.ParseGlob("Public/Templates/*"))

func main() {
	log.Println("Server started on: http://localhost:8898")

	//Handle folders for including css and javascript
	http.Handle("/Public/", http.StripPrefix("/Public/", http.FileServer(http.Dir("Public"))))

	http.HandleFunc("/", Login)

	//Call functions on URL

	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/authenticate", Authenticate)
	http.HandleFunc("/failed", Failed)
	http.HandleFunc("/newbestelling", ServeNewBestelling)

	//Bestelling
	http.HandleFunc("/indexbestelling", IndexBestelling)
	http.HandleFunc("/bekijkbestelling", BekijkBestelling)

	//Klanten
	http.HandleFunc("/indexklanten", IndexKlanten)
	http.HandleFunc("/newklant", NewKlant)
	http.HandleFunc("/insertklant", InsertKlant)

	//Serve
	http.ListenAndServe(":8898", nil)
}
