package api

import (
	"encoding/json"
	"net/http"

	"github.com/pinilla123/cloud-medical-device-simulator/device"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	data := device.GenerateVitalSigns()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
