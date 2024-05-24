package main

import (
	"log"
	"net/http"

	"github.com/pinilla123/cloud-medical-device-simulator/api"
)

func main() {
	http.HandleFunc("/data", api.GetData)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
