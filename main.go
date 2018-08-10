package main

import (
	"net/http"
	"fmt"
	"io"
	"myweapplication/internal"
	)


func person(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "bly")
}



func init() {

}

func main(){
	Hellfunc()





	http.HandleFunc("/person", person)
	fmt.Println("Starting Webserver")
	http.ListenAndServe(":8081", http.DefaultServeMux)
}
