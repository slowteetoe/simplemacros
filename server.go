package main

import (
	"golang.org/x/net/http2"
	"html/template"
	"log"
	"net/http"
)

func main() {
	var server http.Server

	http2.VerboseLogs = false
	server.Addr = ":8080"

	http2.ConfigureServer(&server, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPage := r.URL.Path[len("/"):]
		log.Printf("Serving up [%s]", requestedPage)
		renderPage(w, "templates/index.tmpl", map[string]interface{}{"title": "SimpleMacros"})
	})

	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}

func renderPage(w http.ResponseWriter, templateFile string, data map[string]interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	t.Execute(w, data)
}
