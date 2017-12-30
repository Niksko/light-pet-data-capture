package main

import (
	"crypto/tls"
	"net/http"
	"log"
	"fmt"
)

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	response.Header().Add("X-Frame-Options", "DENY")
	response.Header().Add("X-Content-Type-Options", "nosniff")

	log.Print(fmt.Sprintf("Handled request at / from %s", request.RemoteAddr))

	if (request.Method == http.MethodPost) {
		response.WriteHeader(http.StatusOK)
		log.Print("Request method was POST, sending 200 OK")
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		log.Print(fmt.Sprintf("Request method was %s, sending 405 Method Not Allowed", request.Method))
	}
}


func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", RootHandler)

    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
        	tls.TLS_RSA_WITH_AES_128_CBC_SHA,
        	tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        	tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
        },
    }
    srv := &http.Server{
        Addr:         ":443",
        Handler:      mux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }

    log.Fatal(srv.ListenAndServeTLS("certs/10.0.0.30.crt", "certs/10.0.0.30.key"))
}
