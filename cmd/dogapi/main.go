package main

import (
	"log"
	"net/http"

	"github.com/dogdep/dog-engine/engine"
)

// The engine api.
//  Firstly would expose an api endpoint to receive archive file from API
//  and respond with OK. No need to save the file at this stage.
func main() {
	var listen = ":8080"

	log.Println("Starting HTTP on ", listen)
	http.HandleFunc("/api/deploy", engine.PostDeployHandler)

	log.Fatal(http.ListenAndServe(listen, nil))
}