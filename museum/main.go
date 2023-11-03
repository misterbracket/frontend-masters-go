package main

import (
	"fmt"
	"html/template"
	"net/http"

	"maxklammer.com/go/museum/api"
	"maxklammer.com/go/museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from my Server"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		fmt.Printf("Template error: %v", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = html.Execute(w, data.GetAll())
	if err != nil {
		return
	}
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/templates", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	// server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Println(err)
	}
}
