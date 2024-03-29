package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type WeatherData struct {
	Temperature      float64 // Kelvin
	Pressure         float64 // Pascal
	RelativeHumidity float64 // in percentage
	WindSpeed        float64 // in m/s
	Area             string
	Icon             string
	Format           string
}

func (w WeatherData) String() string {
	result := w.Format

	// icon
	result = strings.ReplaceAll(result, "%c", w.Icon)

	// humidity
	re := regexp.MustCompile(`\%h(?P<unit>(?: ?))(?P<format>(?:\d*\.\d*f)?)`)
	subMatchMap, err := GetSubMatchMap(re, result)
	if err != nil {
		return ""
	}

	humidity_format := "%.0f%%"
	if strings.TrimSpace(subMatchMap["format"]) != "" {
		humidity_format = "%" + strings.ReplaceAll(subMatchMap["format"], "%", "%%")
	}
	result = re.ReplaceAllString(result, fmt.Sprintf(humidity_format+subMatchMap["unit"], w.RelativeHumidity))

	// temperature
	re = regexp.MustCompile(`\%t(?P<unit>(?: ?(?:°C|K)))(?P<format>(?:\d*\.\d*f)?)`)
	subMatchMap, err = GetSubMatchMap(re, result)
	if err != nil {
		return ""
	}
	temperature := w.Temperature
	if subMatchMap["unit"] == "°C" {
		temperature -= 273.15
	}
	temperature_format := "%.1f"
	if strings.TrimSpace(subMatchMap["format"]) != "" {
		temperature_format = "%" + subMatchMap["format"]
	}
	result = re.ReplaceAllString(result, fmt.Sprintf(temperature_format+subMatchMap["unit"], temperature))

	// wind
	re = regexp.MustCompile(`\%w(?P<unit>(?: ?(?:m/s|km/h|m/h|mph)))(?P<format>(?:\d*\.\d*f)?)`)
	subMatchMap, err = GetSubMatchMap(re, result)
	if err != nil {
		return ""
	}
	wind := w.WindSpeed
	if strings.TrimSpace(subMatchMap["unit"]) == "km/h" {
		wind *= 3.6
	}
	if (strings.TrimSpace(subMatchMap["unit"]) == "m/h") || (strings.TrimSpace(subMatchMap["unit"]) == "mph") {
		wind *= 2.236936
	}

	wind_format := "%.1f"
	if strings.TrimSpace(subMatchMap["format"]) != "" {
		wind_format = "%" + subMatchMap["format"]
	}
	result = re.ReplaceAllString(result, fmt.Sprintf(wind_format+subMatchMap["unit"], wind))

	// area
	result = strings.ReplaceAll(result, "%l", w.Area)

	return result
}

type WttrPayload struct { // auto-generated via https://mholt.github.io/json-to-go/
	CurrentCondition []struct {
		FeelsLikeC       string `json:"FeelsLikeC"`
		FeelsLikeF       string `json:"FeelsLikeF"`
		Cloudcover       string `json:"cloudcover"`
		Humidity         string `json:"humidity"`
		LocalObsDateTime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		PrecipInches     string `json:"precipInches"`
		PrecipMM         string `json:"precipMM"`
		Pressure         string `json:"pressure"`
		PressureInches   string `json:"pressureInches"`
		TempC            string `json:"temp_C"`
		TempF            string `json:"temp_F"`
		UvIndex          string `json:"uvIndex"`
		Visibility       string `json:"visibility"`
		VisibilityMiles  string `json:"visibilityMiles"`
		WeatherCode      string `json:"weatherCode"`
		WeatherDesc      []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WeatherIconURL []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		Winddir16Point string `json:"winddir16Point"`
		WinddirDegree  string `json:"winddirDegree"`
		WindspeedKmph  string `json:"windspeedKmph"`
		WindspeedMiles string `json:"windspeedMiles"`
	} `json:"current_condition"`
	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Population string `json:"population"`
		Region     []struct {
			Value string `json:"value"`
		} `json:"region"`
		WeatherURL []struct {
			Value string `json:"value"`
		} `json:"weatherUrl"`
	} `json:"nearest_area"`
	Request []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
		} `json:"astronomy"`
		AvgtempC string `json:"avgtempC"`
		AvgtempF string `json:"avgtempF"`
		Date     string `json:"date"`
		Hourly   []struct {
			DewPointC        string `json:"DewPointC"`
			DewPointF        string `json:"DewPointF"`
			FeelsLikeC       string `json:"FeelsLikeC"`
			FeelsLikeF       string `json:"FeelsLikeF"`
			HeatIndexC       string `json:"HeatIndexC"`
			HeatIndexF       string `json:"HeatIndexF"`
			WindChillC       string `json:"WindChillC"`
			WindChillF       string `json:"WindChillF"`
			WindGustKmph     string `json:"WindGustKmph"`
			WindGustMiles    string `json:"WindGustMiles"`
			Chanceoffog      string `json:"chanceoffog"`
			Chanceoffrost    string `json:"chanceoffrost"`
			Chanceofhightemp string `json:"chanceofhightemp"`
			Chanceofovercast string `json:"chanceofovercast"`
			Chanceofrain     string `json:"chanceofrain"`
			Chanceofremdry   string `json:"chanceofremdry"`
			Chanceofsnow     string `json:"chanceofsnow"`
			Chanceofsunshine string `json:"chanceofsunshine"`
			Chanceofthunder  string `json:"chanceofthunder"`
			Chanceofwindy    string `json:"chanceofwindy"`
			Cloudcover       string `json:"cloudcover"`
			Humidity         string `json:"humidity"`
			PrecipInches     string `json:"precipInches"`
			PrecipMM         string `json:"precipMM"`
			Pressure         string `json:"pressure"`
			PressureInches   string `json:"pressureInches"`
			TempC            string `json:"tempC"`
			TempF            string `json:"tempF"`
			Time             string `json:"time"`
			UvIndex          string `json:"uvIndex"`
			Visibility       string `json:"visibility"`
			VisibilityMiles  string `json:"visibilityMiles"`
			WeatherCode      string `json:"weatherCode"`
			WeatherDesc      []struct {
				Value string `json:"value"`
			} `json:"weatherDesc"`
			WeatherIconURL []struct {
				Value string `json:"value"`
			} `json:"weatherIconUrl"`
			Winddir16Point string `json:"winddir16Point"`
			WinddirDegree  string `json:"winddirDegree"`
			WindspeedKmph  string `json:"windspeedKmph"`
			WindspeedMiles string `json:"windspeedMiles"`
		} `json:"hourly"`
		MaxtempC    string `json:"maxtempC"`
		MaxtempF    string `json:"maxtempF"`
		MintempC    string `json:"mintempC"`
		MintempF    string `json:"mintempF"`
		SunHour     string `json:"sunHour"`
		TotalSnowCm string `json:"totalSnow_cm"`
		UvIndex     string `json:"uvIndex"`
	} `json:"weather"`
}

func (data WttrPayload) createWeather(format string) WeatherData {
	temperature, _ := strconv.ParseFloat(data.CurrentCondition[0].TempC, 64)
	temperature += 273.15
	pressure, _ := strconv.ParseFloat(data.CurrentCondition[0].Pressure, 64)
	pressure *= 1000.
	humidity, _ := strconv.ParseFloat(data.CurrentCondition[0].Humidity, 64)
	wind, _ := strconv.ParseFloat(data.CurrentCondition[0].WindspeedKmph, 64)
	wind /= 3.6
	icon := wttrCodeToIcon(data.CurrentCondition[0].WeatherCode)
	return WeatherData{
		temperature,
		pressure,
		humidity,
		wind,
		data.NearestArea[0].AreaName[0].Value,
		icon,
		format,
	}
}

func wttrDescriptionToIcon(description string) string {
	switch description {
	case "Cloudy":
		return "☁️ "
	case "Fog":
		return "🌫 "
	case "HeavyRain":
		return "🌧 "
	case "HeavyShowers":
		return "🌧 "
	case "HeavySnow":
		return "❄️ "
	case "HeavySnowShowers":
		return "❄️ "
	case "LightRain":
		return "🌦 "
	case "LightShowers":
		return "🌦 "
	case "LightSleet":
		return "🌧 "
	case "LightSleetShowers":
		return "🌧 "
	case "LightSnow":
		return "🌨 "
	case "LightSnowShowers":
		return "🌨 "
	case "PartlyCloudy":
		return "⛅️"
	case "Sunny":
		return "☀️ "
	case "ThunderyHeavyRain":
		return "🌩 "
	case "ThunderyShowers":
		return "⛈ "
	case "ThunderySnowShowers":
		return "⛈ "
	case "VeryCloudy":
		return "☁️ "
	default:
		return "✨"
	}
}

func wttrCodeToDescription(code string) string {
	switch code {
	case "113":
		return "Sunny"
	case "116":
		return "PartlyCloudy"
	case "119":
		return "Cloudy"
	case "122":
		return "VeryCloudy"
	case "143":
		return "Fog"
	case "176":
		return "LightShowers"
	case "179":
		return "LightSleetShowers"
	case "182":
		return "LightSleet"
	case "185":
		return "LightSleet"
	case "200":
		return "ThunderyShowers"
	case "227":
		return "LightSnow"
	case "230":
		return "HeavySnow"
	case "248":
		return "Fog"
	case "260":
		return "Fog"
	case "263":
		return "LightShowers"
	case "266":
		return "LightRain"
	case "281":
		return "LightSleet"
	case "284":
		return "LightSleet"
	case "293":
		return "LightRain"
	case "296":
		return "LightRain"
	case "299":
		return "HeavyShowers"
	case "302":
		return "HeavyRain"
	case "305":
		return "HeavyShowers"
	case "308":
		return "HeavyRain"
	case "311":
		return "LightSleet"
	case "314":
		return "LightSleet"
	case "317":
		return "LightSleet"
	case "320":
		return "LightSnow"
	case "323":
		return "LightSnowShowers"
	case "326":
		return "LightSnowShowers"
	case "329":
		return "HeavySnow"
	case "332":
		return "HeavySnow"
	case "335":
		return "HeavySnowShowers"
	case "338":
		return "HeavySnow"
	case "350":
		return "LightSleet"
	case "353":
		return "LightShowers"
	case "356":
		return "HeavyShowers"
	case "359":
		return "HeavyRain"
	case "362":
		return "LightSleetShowers"
	case "365":
		return "LightSleetShowers"
	case "368":
		return "LightSnowShowers"
	case "371":
		return "HeavySnowShowers"
	case "374":
		return "LightSleetShowers"
	case "377":
		return "LightSleet"
	case "386":
		return "ThunderyShowers"
	case "389":
		return "ThunderyHeavyRain"
	case "392":
		return "ThunderySnowShowers"
	case "395":
		return "HeavySnowShowers"
	default:
		return "Unknown"
	}
}

func wttrCodeToIcon(code string) string {
	return wttrDescriptionToIcon(wttrCodeToDescription(code))
}

type OpenWeatherPayload struct { // auto-generated via https://mholt.github.io/json-to-go/
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
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
		Gust  float64 `json:"gust"`
	} `json:"wind"`
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

func (data OpenWeatherPayload) createWeather(format string) WeatherData {
	return WeatherData{
		data.Main.Temp,
		float64(data.Main.Pressure) * 1000.,
		float64(data.Main.Humidity),
		data.Wind.Speed,
		data.Name,
		openWeatherIconToEmoji(data.Weather[0].Icon),
		format,
	}

}

func openWeatherIconToEmoji(icon string) string {
	switch icon {
	case "01d":
		return "☀️ " 
	case "02d":
		return "⛅️"
	case "03d":
		return "☁️ "
	case "04d":
		return "☁️ "
	case "09d":
		return "🌧 "
	case "10d":
		return "🌦 "
	case "11d":
		return "⛈ "
	case "13d":
		return "❄️ "
	case "50d":
		return "🌫 "
	case "01n":
		return "🌕"
	case "02n":
		return "⛅️"
	case "03n":
		return "☁️ "
	case "04n":
		return "☁️ "
	case "09n":
		return "🌧 "
	case "10n":
		return "☔️"
	case "11n":
		return "⛈ "
	case "13n":
		return "❄️ "
	case "50n":
		return "🌫 "
	default:
		return "✨"
	}
}
