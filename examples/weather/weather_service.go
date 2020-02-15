package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "http://api.openweathermap.org/data/2.5/"

type WeatherService struct {
	APIKey string
}

type WeatherForecast struct {
	Temp          float64
	TempFeelsLike float64
	Icon          string
}

// https://openweathermap.org/weather-conditions
type WeatherResponse struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{APIKey: apiKey}
}

func (w *WeatherService) GetForecast(location string) (WeatherForecast, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/weather", baseURL), nil)
	if err != nil {
		return WeatherForecast{}, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("appid", w.APIKey)
	q.Add("q", location)
	q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return WeatherForecast{}, err
	}
	defer resp.Body.Close()

	weatherForecastResponse := new(WeatherResponse)
	err = json.NewDecoder(resp.Body).Decode(weatherForecastResponse)
	if err != nil {
		return WeatherForecast{}, err
	}

	return mapResponse(weatherForecastResponse), nil
}

func mapResponse(resp *WeatherResponse) WeatherForecast {
	return WeatherForecast{
		TempFeelsLike: resp.Main.FeelsLike,
		Temp:          resp.Main.Temp,
		Icon:          mapCondition(resp.Weather[0].ID),
	}
}

// https://unicode.org/emoji/charts/full-emoji-list.html
func mapCondition(conditionID int) string {
	switch true {
	case conditionID < 300:
		return "🌩" // thunderstorm
	case conditionID < 600:
		return "🌧" // drizzle or rain
	case conditionID < 700:
		return "🌨" // snow
	case conditionID < 800:
		return "🌫" // fog
	case conditionID == 800:
		return "☀" // clear sky
	case conditionID > 800:
		return "☁" // cloud
	default:
		return "❓"
	}
}
