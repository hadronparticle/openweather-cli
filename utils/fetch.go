package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiKey  = "charcharpachoychoy"
	baseURL = "http://api.openweathermap.org/data/2.5/weather"
)

type weatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`
	Message string `json:"message"`
}

func FetchWeather(city string) error {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Request Error: %v", err)
	}
	defer resp.Body.Close()

	var data weatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("Response parsing error: %v", err)
	}

	if resp.StatusCode == http.StatusOK {
		weatherDescription := data.Weather[0].Description
		temperature := data.Main.Temp
		feelsLike := data.Main.FeelsLike
		pressure := data.Main.Pressure
		humidity := data.Main.Humidity
		windSpeed := ConvertSpeed(data.Wind.Speed)
		windDirection := GetWindDirection(data.Wind.Deg)
		sunrise := ConvertToLocalTime(data.Sys.Sunrise, 8)
		sunset := ConvertToLocalTime(data.Sys.Sunset, 8)

		fmt.Printf("Weather in %s:\n", city)
		fmt.Printf("Description: %s\n", weatherDescription)
		fmt.Printf("Temperature: %.2f°C\n", temperature)
		fmt.Printf("Feels Like: %.2f°C\n", feelsLike)
		fmt.Printf("Pressure: %.2f hPa\n", pressure)
		fmt.Printf("Humidity: %.2f%%\n", humidity)
		fmt.Printf("Wind Speed: %.2f km/h (%s)\n", windSpeed, windDirection)
		fmt.Printf("Sunrise: %s (UTC+8)\n", sunrise)
		fmt.Printf("Sunset: %s (UTC+8)\n", sunset)
	} else {
		return fmt.Errorf("Error: %s", data.Message)
	}

	return nil
}
