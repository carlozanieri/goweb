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

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Chi siamo",
			Nav:   []string{"Home", "About", "Contact"},
			Year:  time.Now().Year(),
		}
		if err := tmpl.ExecuteTemplate(w, "about", data); err != nil {
			http.Error(w, "Errore", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Contatti",
			Nav:   []string{"Home", "About", "Contact"},
			Year:  time.Now().Year(),
		}
		if err := tmpl.ExecuteTemplate(w, "contact", data); err != nil {
			http.Error(w, "Errore", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/article/", func(w http.ResponseWriter, r *http.Request) {
		// Estrarre l'ID dalla URL: /article/123 → "123"
		id := r.URL.Path[len("/article/"):]

		// Se non c’è ID → 404
		if id == "" {
			http.NotFound(w, r)
			return
		}

		// Fake database
		articles := map[string]string{
			"1": "Introduzione a Go",
			"2": "Usare i template in Go",
			"3": "Costruire un sito modulare",
		}

		title, ok := articles[id]
		if !ok {
			title = "Articolo non trovato"
		}

		data := PageData{
			Title: "Articolo " + id,
			Nav:   []string{"Home", "About", "Contact"},
			Items: []string{title}, // Riutilizziamo Items come contenuto articolo
			Year:  time.Now().Year(),
		}

		if err := tmpl.ExecuteTemplate(w, "article", data); err != nil {
			http.Error(w, "Errore rendering articolo", http.StatusInternalServerError)
		}
	})

	// file statici
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
