package main

import (
	"crypto/tls"
	"net/http"
	"log"
	"github.com/gorilla/handlers"
	"os"
	"github.com/golang/protobuf/proto"
	"github.com/niksko/light-pet-data-capture/http-handlers"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
    	http_handlers.RootHandler(response, request, proto.UnmarshalText)
	})

    loggedMux := handlers.LoggingHandler(os.Stdout, mux)

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
        Handler:      loggedMux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }

    log.Fatal(srv.ListenAndServeTLS("certs/10.0.0.30.crt", "certs/10.0.0.30.key"))
}
