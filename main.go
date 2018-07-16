package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Ready to serve request")
	http.Handle("/", http.FileServer(http.Dir("./src")))
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
