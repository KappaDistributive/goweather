package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var logger = log.New(os.Stderr, "", 0)

func getWeather(city string, provider string, api_key string, format string) (WeatherData, error) {
	if strings.HasPrefix(provider, "wttr") {
		resp, err := http.Get(fmt.Sprintf("https://wttr.in/%s?format=j1", city))
		if err != nil {
			logger.Println(err)
			return WeatherData{}, err
		} else if resp.StatusCode != 200 {
			return WeatherData{}, errors.New(fmt.Sprintf("Status code %d", resp.StatusCode))
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return WeatherData{}, err
			}
			var data WttrPayload
			err = json.Unmarshal(body, &data)
			if err != nil {
				return WeatherData{}, err
			}
			return data.createWeather(format), nil
		}
	} else if strings.HasPrefix(provider, "openweather") {
		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key))
		if err != nil {
			return WeatherData{}, err
		} else if resp.StatusCode != 200 {
			return WeatherData{}, errors.New(fmt.Sprintf("Status code %d", resp.StatusCode))
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return WeatherData{}, err
			}
			var data OpenWeatherPayload
			err = json.Unmarshal(body, &data)
			if err != nil {
				return WeatherData{}, err
			}
			return data.createWeather(format), nil
		}
	}
	return WeatherData{}, errors.New("Reached end of function without return. This should never happen.")
}

func main() {

	var city string
	var weather_method string
	var api_key string
	var format string
	flag.StringVar(&city, "city", "n/a", "Name of the city.")
	flag.StringVar(&weather_method, "method", "n/a", "Method to use to obtain weather data.")
	flag.StringVar(&format, "format", "%c %t°C", "Output format.")
	flag.StringVar(&api_key, "api_key", "n/a", "API key.")
	flag.Parse()

	weather, err := getWeather(city, weather_method, api_key, format)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(weather)
}
