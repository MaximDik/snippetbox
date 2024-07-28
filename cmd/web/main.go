package main

import (
	"log"
	"net/http"

	"github.com/MaximDik/snippetbox/cmd/web/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	log.Println("server is listening...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
