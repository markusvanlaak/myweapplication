package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Email string `json:"email"`
	PW    string `json:"pwd"`
}

func Api(w http.ResponseWriter, r *http.Request) {

	switch string(r.Method) {
	case "POST":
		log.Println(r.Method)
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var msg Message
		err = json.Unmarshal(b, &msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		output, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		log.Println(output)
		w.Header().Set("content-type", "application/json")
		w.Write(output)

	case "GET":
		log.Println(r.Method)
	}

	//log.Println(r)
	w.Header().Set("Content-Type", "application/json")
	//log.Println(r.Method, r.URL)
	fmt.Fprintf(w, "{\"fname\":\"Markus\",\"lname\": \"van Laak\",\"age\": 43}")
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "Mattis")
}
