package main

import "testing"

func TestCurrentWeather(t *testing.T) {
	data := getCurrentByZIP(94040, "us") //Should get weather from Mountain View CA

	if data.COD != 200 {
		t.Errorf("Error talking to server, got code %d", data.COD)
	}
	if data.Name != "Mountain View" {
		t.Errorf("Was expecting Mountain View, got %s", data.Name)
	}
}
