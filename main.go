/*
 * ACME Reviews - PSIDI
 *
 * Swagger proposed server for the Review Product infrastructure at ACME .Inc
 *
 * API version: 0.1
 * Contact: 1171071@isep.ipp.pt
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
	"os"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
)

func main() {
	log.Printf("Server started")
	router := NewRouter()
	port := os.Getenv("PORT")

	log.Println("TESTE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
