package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

const (
	TLS_KEY_PATH  = "KEY"
	TLS_CERT_PATH = "CERT"
)

func main() {
	key, ok := os.LookupEnv(TLS_KEY_PATH)
	if !ok {
		log.Fatalln("env KEY not found")
	}
	cert, ok := os.LookupEnv(TLS_CERT_PATH)
	if !ok {
		log.Fatalln("env CERT not found")
	}
	log.Println("Ready to serve request")
	http.Handle("/", http.FileServer(http.Dir("./src")))

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}
	log.Fatal(server.ListenAndServeTLS(cert, key))
}
