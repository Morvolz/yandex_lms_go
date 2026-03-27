package main

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	coefficient := 1.0
	if weather.Condition == 2 {
		coefficient += 0.2
	} else if weather.Condition == 3 {
		coefficient += 0.15
	} else if weather.Condition == 1 {
		coefficient += 0.125
	}
	if weather.WindSpeed > 15 {
		coefficient += 0.1
	}

	return coefficient
}
