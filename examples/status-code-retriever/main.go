// Steve Phillips / elimisteve
// 2013.09.03

package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	LISTEN_ADDR = ":8080"
)

var (
	router = mux.NewRouter()
)

func init() {
	router.HandleFunc("/", ShowSchema).Methods("OPTIONS")

	http.Handle("/", router)
}

func main() {
	serve(router)
}

func ShowSchema(w http.ResponseWriter, r *http.Request) {
	w.Write(BLOCK_DEFINITION)
}

func serve(handler http.Handler) {
	server := &http.Server{
		Addr:           LISTEN_ADDR,
		Handler:        handler,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("HTTP server trying to listen on %s...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP listen failed: %v\n", err)
	}
}

var BLOCK_DEFINITION = []byte(`
{
  "name": "HTTP Status Code Retriever",
  "url": "http://status-code-retriever.herokuapp.com",
  "description": "Visits the given URL and returns the HTTP status code.",
  "inputs": {
      "name": "url",
      "type": "String",
      "description": "URL to be visited."
  },
  "outputs": {
      "name": "status_code",
      "type": "Number",
      "description": "HTTP status code returned by given URL."
  }
}
`)
