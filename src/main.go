package main

import (
	"html/template"
	"log"
	"mime"
	"net/http"
	"os"
)

func main() {
    // Set MIME type for CSS files
    mime.AddExtensionType(".css", "text/css")

    // Log the current working directory
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Current working directory:", cwd)

    // Serve static files
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Serve the home page
    http.HandleFunc("/", HomeHandler)

    log.Println("Starting server on :8000")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatal(err)
    }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("src/templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}