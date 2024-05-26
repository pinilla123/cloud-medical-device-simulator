package device

import (
	"fmt"
	"math/rand"
	"time"
)

type VitalSigns struct {
	DeviceID      string
	HeartRate     int
	BloodPressure string
}

func GenerateVitalSigns(deviceID string) VitalSigns {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	heartRate := r.Intn(40) + 60 // Random heart rate between 60 and 100
	systolic := r.Intn(40) + 90  // Random systolic BP between 90 and 130
	diastolic := r.Intn(20) + 60 // Random diastolic BP between 60 and 80

	return VitalSigns{
		DeviceID:      deviceID,
		HeartRate:     heartRate,
		BloodPressure: fmt.Sprintf("%d/%d", systolic, diastolic),
	}
}
