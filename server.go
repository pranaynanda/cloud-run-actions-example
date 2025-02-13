package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Phrase struct {
	Text string `json:"phrase"`
}

func HelloWorld() Phrase {
	return Phrase{
		Text: "Hello Cloud Run!",
	}
}

func GetPhrase(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HelloWorld())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetPhrase).Methods("GET")

	port := os.Getenv("PORT")
	log.Print("Started API on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
