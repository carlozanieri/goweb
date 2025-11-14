package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type PageData struct {
	Title string
	Nav   []string
	Items []string
	Year  int
}

func main() {
	// carica tutti i template
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Home — Sito con base.html",
			Nav:   []string{"Home", "About", "Contact"},
			Items: []string{"Elemento 1", "Elemento 2", "Elemento 3"},
			Year:  time.Now().Year(),
		}

		if err := tmpl.ExecuteTemplate(w, "home", data); err != nil {
			http.Error(w, "Errore", http.StatusInternalServerError)
			log.Println("template:", err)
		}
	})

	// file statici
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
