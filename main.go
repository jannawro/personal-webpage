package main

import (
	"log"
	"net/http"
	"path/filepath"
	"html/template"
)

func main() {
    fs := http.FileServer(http.Dir("./assets"))
    http.Handle("/assets/", http.StripPrefix("/assets/", fs))

    http.HandleFunc("/", serveTemplate)

    log.Print("Listening on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    baseLayout := filepath.Join("templates", "layout.html")
    pageBody := filepath.Join("templates", filepath.Clean(r.URL.Path) + ".html")

    tmpl, err := template.New("").ParseFiles(baseLayout, pageBody)
    if err != nil {
        log.Print(err)
    }
    tmpl.ExecuteTemplate(w ,"layout", nil)
}
