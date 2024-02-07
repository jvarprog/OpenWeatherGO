package OpenWeatherGO

import (
	"fmt"
	"testing"
)

func TestGeocodeZIP(t *testing.T) {
	data := GeocodeByZIP(94040, "us") //Should get weather from Mountain View CA
	fmt.Println(data)
	if data.Name != "Mountain View" {
		t.Errorf("Was expecting Mountain View, got %s", data.Name)
	}
}
