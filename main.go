package main

import (
	"github.com/markusvanlaak/go/myweapplication/internal"
	"net/http"
	"log"
)

func init() {

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/person", internal.GetData)
	http.HandleFunc("/", internal.About)
	http.HandleFunc("/blogs", internal.GetBlogs)
	log.Println("Starting Webserver")
	http.ListenAndServe(":8081", http.DefaultServeMux)
}
