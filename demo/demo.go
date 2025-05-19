package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("Started")

	if err := http.ListenAndServe("localhost:8300", http.HandlerFunc(handle)); err != nil {
		log.Fatalf("%v", err)
	}
}

type request struct {
	Method     string      `json:"method,omitempty"`
	RequestUri string      `json:"requestUri,omitempty"`
	Host       string      `json:"host,omitempty"`
	URL        string      `json:"url,omitempty"`
	RemoteAddr string      `json:"remoteAddr,omitempty"`
	Header     http.Header `json:"header,omitempty"`
	Proto      string      `json:"proto,omitempty"`
}

func handle(rw http.ResponseWriter, r *http.Request) {
	req := request{
		Method:     r.Method,
		RequestUri: r.RequestURI,
		Host:       r.Host,
		RemoteAddr: r.RemoteAddr,
		Header:     r.Header,
		Proto:      r.Proto,
	}

	if r.URL != nil {
		req.URL = r.URL.String()
	}

	encLog := json.NewEncoder(os.Stderr)
	encLog.SetEscapeHTML(false)
	_ = encLog.Encode(req)

	enc := json.NewEncoder(rw)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	_ = enc.Encode(req)
}
