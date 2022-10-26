package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type visualCrossingWeatherAPI struct {
	ApiKey string
}

type DaysResponse struct {
	Temperature float32 `json:"tempmax"`
	Day         string  `json:"datetime"`
}

type WeatherApiResponse struct {
	Days []DaysResponse `json:"days"`
}

func (w visualCrossingWeatherAPI) GetTemperatures(from, to time.Time, city string) map[time.Time]int {
	fromStr := from.Format("2006-01-02")
	toStr := to.Format("2006-01-02")
	response, err := http.Get(fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?unitGroup=metric&key=%s&contentType=json", city, fromStr, toStr, w.ApiKey))
	if err != nil {
		fmt.Print(err.Error())
		return map[time.Time]int{}
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return map[time.Time]int{}
	}
	var responseObject WeatherApiResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Print(err.Error())
		return map[time.Time]int{}
	}
	responseMap := map[time.Time]int{}
	for _, v := range responseObject.Days {
		dayInDateTime, err := time.Parse("2006-01-02", v.Day)
		if err != nil {
			fmt.Print(err.Error())
			return map[time.Time]int{}
		}
		responseMap[dayInDateTime] = int(v.Temperature)
	}
	return responseMap
}
