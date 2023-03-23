package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	getHttpClientTest()

}

func getHttpClientTest() {
	mTLSConfig := getTlsConfig()

	tr := &http.Transport{
		TLSClientConfig: mTLSConfig,
	}
	c := &http.Client{Transport: tr}

	res, err := c.Get("https://api-dev.lean-tools.cz/v1/system/version")
	if err != nil {
		log.Panicf("error making http request: %q\n", err)
	}

	log.Printf("Request successfull with code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicf("client: could not read response body: %q\n", err)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}

func getTlsConfig() *tls.Config {
	return &tls.Config{
		// Certificates: []tls.Certificate{cert},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       true,
		VerifyConnection: func(state tls.ConnectionState) error {
			fmt.Println(utils.ObjectPrettyPrint(state))
			return nil
		},
	}
}
