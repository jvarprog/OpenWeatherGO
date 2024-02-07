package OpenWeatherGO

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Coordinate struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lon"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Cloud struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type CurrentWeather struct {
	Coord      Coordinate `json:"coord"`
	Weather    []Weather  `json:"weather"`
	Base       string     `json:"base"`
	Main       Main       `json:"main"`
	Visibility int        `json:"visibility"`
	Wind       Wind       `json:"wind"`
	Clouds     Cloud      `json:"clouds"`
	Dt         int64      `json:"dt"`
	Sys        Sys        `json:"sys"`
	Timezone   int        `json:"timezone"`
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	COD        int        `json:"cod"`
}

func GetCurrentByZIP(ZIP int, countryCode string) CurrentWeather {
	var APIKEY = os.Getenv("OWKEY")
	baseURL := "https://api.openweathermap.org/data/2.5/weather?zip=%d,%s&appid=%s"
	completeURL := fmt.Sprintf(baseURL, ZIP, countryCode, APIKEY)

	resp, err := http.Get(completeURL)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var data *CurrentWeather
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data
}

func GetCurrentByCoords(long float32, lat float32) CurrentWeather {
	var APIKEY = os.Getenv("OWKEY")
	baseURL := "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v"
	completeURL := fmt.Sprintf(baseURL, long, lat, APIKEY)

	resp, err := http.Get(completeURL)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var data *CurrentWeather
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return *data

}
