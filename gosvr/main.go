package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := flag.Int("port", 8080, "the port number you wish to serve on")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", getRootDetailsAgain)
	http.Handle("/", r)
	fmt.Printf("Serving on port %v", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}

func getRootDetailsAgain(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "take details from string")
}
