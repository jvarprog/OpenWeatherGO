package OpenWeatherGO

import (
	"testing"
)

func TestGeocodeName(t *testing.T) {
	data := GeocodeByName("london")
	if data[0].Name != "London" && data[0].Country != "GB" {
		t.Errorf("Was expecting London and GB, got %s and %s", data[0].Name, data[0].Country)
	}
}

func TestGeocodeZIP(t *testing.T) {
	data := GeocodeByZIP(94040, "us") //Should get weather from Mountain View CA
	if data.Name != "Mountain View" {
		t.Errorf("Was expecting Mountain View, got %s", data.Name)
	}
}
