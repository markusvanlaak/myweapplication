package internal

import (
	"html/template"
	"gopkg.in/mgo.v2"
	"net/http"
	"log")

type Link struct {
	Link string
	Name string
}

type Person struct {
	Fname string
	Lname string
	Age   int
}

type Blog struct {
	Text    string
	Heading string
}

var Session *mgo.Session

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetData(w http.ResponseWriter, r *http.Request)  {
	Session, err := mgo.Dial("127.0.0.1:27017")
	Check(err)

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)

	defer Session.Close()

	P := getperson(Session)
	L := getlink(Session)

	Text0 := Blog{"This is my first text", "First"}

	Links := L
	Persons := P
	Texts := []Blog{Text0}

	data := struct {
		MyLinks   []Link
		MyPersons []Person
		MyBlog    []Blog
	}{
		Links,
		Persons,
		Texts,
	}

	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))

	tpl.ExecuteTemplate(w, "people", data)
	Check(err)
	log.Println(r.Method, r.URL)
}

func getperson(mongoSession *mgo.Session) []Person {

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var Result []Person

	d := sessionCopy.DB("test").C("people")
	d.Find(nil).All(&Result)

	//log.Println(Result)

	return Result
}

func getlink(mongoSession *mgo.Session) []Link {

	sessionCopy2 := mongoSession.Copy()
	defer sessionCopy2.Close()

	var l []Link

	d := sessionCopy2.DB("test").C("links")
	d.Find(nil).All(&l)

	return l
}
