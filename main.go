package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

//set templates
var tmpl = template.Must(template.ParseGlob("form/*"))

func main() {
	log.Println("Server started on: http://localhost:8898")


	//Call functions on URL


	//Login
	http.HandleFunc("/", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/authenticate", Authenticate)
	http.HandleFunc("/failed", Failed)


	//Bestelling
	http.HandleFunc("/indexbestelling", IndexBestelling)
	http.HandleFunc("/bekijkbestelling", BekijkBestelling)


	//Serve
	http.ListenAndServe(":8898", nil)
}
