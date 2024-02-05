package openweathergo

import (
	"os"
	"testing"
)

func TestCurrentWeatherZIP(t *testing.T) {
	os.Setenv("OWKEY", "f0b7c0f3986510cac41a23797f3fb0a5")
	data := GetCurrentByZIP(94040, "us") //Should get weather from Mountain View CA

	if data.COD != 200 {
		t.Errorf("Error talking to server, got code %d", data.COD)
	}
	if data.Name != "Mountain View" {
		t.Errorf("Was expecting Mountain View, got %s", data.Name)
	}
}

func TestCurrentWeatherCOORD(t *testing.T) {
	os.Setenv("OWKEY", "f0b7c0f3986510cac41a23797f3fb0a5")
	data := GetCurrentByCoords(37.38, -122.08)
	if data.COD != 200 {
		t.Errorf("Error talking to server, got code %d", data.COD)
	}
	if data.Name != "Mountain View" {
		t.Errorf("Was expecting Mountain View, got %s", data.Name)
	}
}
