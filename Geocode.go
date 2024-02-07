package OpenWeatherGO

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Geocode struct {
	ZIP     string  `json:"zip"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

type MultiGeocode struct {
	Results []Geocode `json:""`
}

func GeocodeByZIP(ZIP int, countryCode string) Geocode {
	var APIKEY = os.Getenv("OWKEY")
	baseURL := "http://api.openweathermap.org/geo/1.0/zip?zip=%v,%s&appid=%s"
	completeURL := fmt.Sprintf(baseURL, ZIP, countryCode, APIKEY)

	resp, err := http.Get(completeURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data *Geocode
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data

}

func GeocodeByName(cityName string) []Geocode {
	var APIKEY = os.Getenv("OWKEY")
	baseURL := "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
	completeURL := fmt.Sprintf(baseURL, cityName, APIKEY)

	resp, err := http.Get(completeURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []Geocode
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
