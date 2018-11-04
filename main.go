package main

import (
	"log"
	"net/http"

	"github.com/markusvanlaak/go/myweapplication/internal"
)

func init() {

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	http.HandleFunc("/person", internal.GetData)

	http.HandleFunc("/about", internal.About)
	//http.HandleFunc("/req", internal.Reqapiexample)

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/api", internal.Api)

	log.Println("Starting Webserver")
	http.ListenAndServe(":8081", http.DefaultServeMux)
}
