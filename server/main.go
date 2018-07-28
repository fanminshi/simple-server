package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	TLS_KEY_PATH     = "KEY"
	TLS_CERT_PATH    = "CERT"
	TLS_CA_CERT_PATH = "CA_CERT"
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
	http.Handle("/", http.FileServer(http.Dir("./src")))

	cacert, ok := os.LookupEnv(TLS_CA_CERT_PATH)
	if !ok {
		log.Fatalln("env CERT not found")
	}
	caCert, err := ioutil.ReadFile(cacert)
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalln("Unable to add certificate to certificate pool")
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}

	log.Println("Ready to serve request")
	log.Fatal(server.ListenAndServeTLS(cert, key))
}
