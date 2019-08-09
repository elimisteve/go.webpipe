package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	listenPort := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		listenPort = envPort
	}

	router := mux.NewRouter()

	router.HandleFunc("/", ShowSchema).Methods("OPTIONS")
	router.HandleFunc("/", RetrieveStatusCode).Methods("POST")

	http.Handle("/", router)

	server := &http.Server{
		Addr:           ":" + listenPort,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("HTTP server trying to listen on %s...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("HTTP listen failed: %v\n", err)
	}
}

// RetrieveStatusCode expects a POST body containing {"inputs":
// {"url": "http://..."}} and returns the status code retrieved by
// visiting the given URL
func RetrieveStatusCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Read Request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Bad Request
		log.Printf("Error at %s: %v\n", r.URL, err)
		http.Error(w, `{"outputs": []}`, 400)
		return
	}
	defer r.Body.Close()

	// Parse JSON
	input := URLInput{}
	if err := json.Unmarshal(body, &input); err != nil {
		// Bad Request
		log.Printf("Error at %s: %v\n", r.URL, err)
		http.Error(w, `{"outputs": []}`, 400)
		return
	}

	// Perform HTTP request for user
	resp, err := http.Get(input.Inputs.URL)
	if err != nil {
		// Something weird happened...
		log.Printf("Error at %s: %v\n", r.URL, err)
		http.Error(w, `{"outputs": []}`, 500)
		return
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, `{"outputs": [{"status_code": %d}]}`, resp.StatusCode)
}

type URLInput struct {
	Inputs struct {
		URL string `json:"url"`
	} `json:"inputs"`
}

// ShowSchema responds to an OPTIONS request with the JSON schema
// describing this service.
func ShowSchema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`
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
`))
}
