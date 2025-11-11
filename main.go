package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PageData contains data for the template
type PageData struct {
	Title       string
	Description string
	ThemeColor  string
}

func main() {
	// Serve static files (like momento.svg)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/momento.svg", fs)

	// Handle the main page
	http.HandleFunc("/", handleIndex)

	port := ":2129"
	log.Printf("Server is starting on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Only serve the index on the root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Parse the template
	tmplPath := filepath.Join(".", "templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare data for the template
	data := PageData{
		Title:       "Momento Mori",
		Description: "Moento",
		ThemeColor:  "#9f70ed",
	}

	// Execute the template
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
