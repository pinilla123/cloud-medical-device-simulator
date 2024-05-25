package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pinilla123/cloud-medical-device-simulator/device"
)

func GetData(w http.ResponseWriter, r *http.Request) {
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
