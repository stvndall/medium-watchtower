package main

import (
	"net/http"
	"log"
	"io"
	"fmt"
	"flag"
)



func main(){

port := flag.Int("port", 8080, "the port number you wish to serve on")
flag.Parse()

http.HandleFunc("/", getRootDetails)
fmt.Printf("Serving on port %v", *port)
log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",*port), nil))
}

func getRootDetails (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is the first go image")
}