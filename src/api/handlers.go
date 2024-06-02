package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pinilla123/cloud-medical-device-simulator/src/device"
)

type Request struct {
	DeviceID string `json:"device_id"`
}

type Response struct {
	Message string            `json:"message"`
	Device  device.VitalSigns `json:"device"`
}

// LambdaHandler is the AWS Lambda handler function
func LambdaHandler(ctx context.Context, req Request) (Response, error) {
	log.Printf("Received request: %+v", req)

	if req.DeviceID == "" {
		return Response{Message: "device_id is required"}, fmt.Errorf("device_id is required")
	}

	data := device.GenerateVitalSigns(req.DeviceID)
	err := device.SaveVitalSigns(data)
	if err != nil {
		log.Printf("Failed to save data: %v", err)
		return Response{Message: "Failed to save data"}, err
	}

	return Response{
		Message: "success",
		Device:  data,
	}, nil
}

// HTTPHandler is the HTTP handler function
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("device_id")
	if deviceID == "" {
		http.Error(w, "device_id is required", http.StatusBadRequest)
		return
	}

	data := device.GenerateVitalSigns(deviceID)
	err := device.SaveVitalSigns(data)
	if err != nil {
		log.Printf("Failed to save data: %v", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
