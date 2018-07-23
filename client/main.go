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
	clientCert, err := tls.LoadX509KeyPair(key, cert)
	if err != nil {
		log.Fatalln("unable to load cert")
	}
	clientCACert, err := ioutil.ReadFile(TLS_CA_CERT_PATH)
	if err != nil {
		log.Fatal("Unable to open cert", err)
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
			TLSClientConfig: &tls.Config{
				RootCAs: clientCertPool,
			},
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
