package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pinilla123/cloud-medical-device-simulator/src/api"
)

func main() {
	lambda.Start(api.LambdaHandler)

	http.HandleFunc("/data", api.HTTPHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
