package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var addr string
var root string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "addr to listen on")
	flag.StringVar(&root, "root", ".", "directory to serve")
}

func logged(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("%s %s %s %s", r.Method, r.URL, r.RemoteAddr, duration)
	})
}

func main() {
	flag.Parse()
	handler := logged(http.FileServer(http.Dir(root)))
	log.Printf("serving directory %q on %q", root, addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
