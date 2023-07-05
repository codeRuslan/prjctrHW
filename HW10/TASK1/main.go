package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

type WeatherData struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
}

func main() {
	handleRequest()
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/weather/{city}", GetCurrentWeather).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["city"]
	dataResponse := GetWeatherAPIResponse(key)
	json.NewEncoder(w).Encode(dataResponse.Main)
}

func GetWeatherAPIResponse(city string) (apiResponse WeatherData) {
	url := "https://open-weather13.p.rapidapi.com/city/" + city

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("X-RapidAPI-Key"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("X-RapidAPI-Host"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data WeatherData

	err := json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	return data
}
