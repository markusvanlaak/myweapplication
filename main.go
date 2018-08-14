package main

import (
	"github.com/markusvanlaak/go/myweapplication/internal"
	"net/http"
	"log"
)

func init() {

}

func main() {
	http.HandleFunc("/person", internal.GetData)
	log.Println("Starting Webserver")
	http.ListenAndServe(":8081", http.DefaultServeMux)
}
