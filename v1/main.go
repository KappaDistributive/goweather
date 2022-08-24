package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	var city string
	var weather_method string
	var api_key string
	var format string
	flag.StringVar(&city, "city", "n/a", "Name of the city.")
	flag.StringVar(&weather_method, "method", "n/a", "Method to use to obtain weather data.")
	flag.StringVar(&format, "format", "%c %t°C", "Output format.")
	flag.StringVar(&api_key, "api_key", "n/a", "API key.")
	flag.Parse()

	if strings.HasPrefix(weather_method, "wttr") {
		resp, err := http.Get("https://wttr.in/" + city + "?format=j1")
		if err != nil {
			logger.Println(err)
			os.Exit(1)
		} else if resp.StatusCode != 200 {
			logger.Printf("Status code %d", resp.StatusCode)
			os.Exit(1)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				logger.Println(err)
				os.Exit(1)
			}
			var data WttrPayload
			err = json.Unmarshal(body, &data)
			if err != nil {
				logger.Println(err)
				os.Exit(1)
			}
			weather := data.createWeather(format)
			fmt.Println(weather)
		}
	} else if strings.HasPrefix(weather_method, "openweather") {
		resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + api_key)
		if err != nil {
			logger.Println(err)
			os.Exit(1)
		} else if resp.StatusCode != 200 {
			logger.Printf("Status code %d", resp.StatusCode)
			os.Exit(1)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				logger.Println(err)
				os.Exit(1)
			}
			var data OpenWeatherPayload
			err = json.Unmarshal(body, &data)
			weather := data.createWeather(format)
			fmt.Println(weather)
		}
	}
}
