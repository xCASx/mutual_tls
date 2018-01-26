package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "crypto/x509"
    "io/ioutil"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
        w.Write([]byte("This is an example server.\n"))
    })

    caCert, err := ioutil.ReadFile("keys/rootCA.pem")
    if err != nil {
        log.Fatal(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    cfg := &tls.Config{
        ClientCAs:                caCertPool,
        ClientAuth:               tls.RequireAndVerifyClientCert,
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
//        CipherSuites: []uint16{
//            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
//            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
//            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
//            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
//        },
    }
    srv := &http.Server{
        Addr:         ":8443",
        Handler:      mux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("keys/server.pem", "keys/server.key"))
}
