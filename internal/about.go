package internal

import (
	"net/http"
	"html/template"
	"log"
)

func About(w http.ResponseWriter, r *http.Request)  {

	Text0 := Blog{"This is my first text", "First"}
	Texts := []Blog{Text0}

	data := struct {
		MyBlog    []Blog
	}{
		Texts,
	}

	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))

	tpl.ExecuteTemplate(w, "about", data)
	log.Println(r.Method, r.URL)
}
