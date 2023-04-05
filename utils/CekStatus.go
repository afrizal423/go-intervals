package utils

import "github.com/afrizal423/go-intervals/models"

func CekStatus(payload models.Cuaca) (string, string) {
	// Menentukan status water dan wind sesuai dengan ketentuan yang telah diberikan.
	statusWater := "aman"
	if payload.Water > 8 {
		statusWater = "bahaya"
	} else if payload.Water >= 6 {
		statusWater = "siaga"
	}
	statusWind := "aman"
	if payload.Wind > 15 {
		statusWind = "bahaya"
	} else if payload.Wind >= 7 {
		statusWind = "siaga"
	}

	return statusWater, statusWind
}
