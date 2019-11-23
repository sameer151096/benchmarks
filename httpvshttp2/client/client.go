package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{}
	buf := new(bytes.Buffer)
	reqobj := Request{}
	reqobj.Key = 1

	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client.Transport = &http2.Transport{TLSClientConfig: tlsConfig}

	json.NewEncoder(buf).Encode(&reqobj)

	response, err := client.Post("https://localhost:8080/serve", "", buf)
	if err != nil {
		fmt.Print("\n error while sending request.. Error ", err)
	}
	responseobj := Response{}
	err = json.NewDecoder(response.Body).Decode(&responseobj)
	if err != nil {
		fmt.Print("\n error while decoding response.. Error : ", err)
	}
	fmt.Print("\n response received is ", responseobj.Value)

	fmt.Print("\n Protocol - ", response.Proto)
}

type Request struct {
	Key int
}

type Response struct {
	Value string
}
