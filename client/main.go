package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	TLS_KEY_PATH     = "KEY"
	TLS_CERT_PATH    = "CERT"
	TLS_CA_CERT_PATH = "CA_CERT"
	SVC_ADDR         = "SVC"
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
	clientCert, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatalf("unable to load cert: %v", err)
	}
	cacert, ok := os.LookupEnv(TLS_CA_CERT_PATH)
	if !ok {
		log.Fatalln("env CA_CERT not found")
	}
	clientCACert, err := ioutil.ReadFile(cacert)
	if err != nil {
		log.Fatalf("Unable to open cert: %v", err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      clientCertPool,
	}
	tlsConfig.BuildNameToCertificate()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	svc, ok := os.LookupEnv(SVC_ADDR)
	if !ok {
		log.Fatalln("env SVC not found")
	}

	resp, err := client.Get(fmt.Sprintf("https://%v:8080", svc))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Lets print the message
	log.Println(string(body))
}
