package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/MaximDik/snippetbox/cmd/web/handlers"
)

func main() {

	addr := flag.String("addr", ":4000", "Сетевой фдрес HTTP")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("server is listening on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
