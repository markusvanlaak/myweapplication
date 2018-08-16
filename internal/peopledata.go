package internal

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
)

type blog struct {
	Text string
	Heading string
}

func getmgosession() *mgo.Session {
	mgosession, err := mgo.Dial("localhost:27017")
	Check(err)
	mgosession.SetMode(mgo.Monotonic, true)
	defer mgosession.Close()
	return mgosession
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	s := getmgosession()
	fmt.Fprint(w, s.DB("test").C("people"))
}
