package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const HostUrl = "localhost:1234"

type weatherAPI interface {
	GetTemperatures(from, to time.Time, city string) map[time.Time]int
}

func checkForTemperaturesBelow11(weatherApi weatherAPI, from, to time.Time, city string) bool {
	temperatures := weatherApi.GetTemperatures(from, to, city)
	for _, v := range temperatures {
		if v < 11 {
			return true
		}
	}
	return false
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "healthy")
}

func getRecommendationForRanges(from, to time.Time, api weatherAPI, city string) []string {
	shouldAddJacket := checkForTemperaturesBelow11(api, from, to, city)
	rangeOfTime := to.Sub(from)
	days := int(rangeOfTime.Hours() / 24)
	expectedClothes := []string{fmt.Sprintf("%d x tshirt", days), fmt.Sprintf("%d x pair of socks", days), fmt.Sprintf("%d x panties", days)}
	if shouldAddJacket {
		expectedClothes = append(expectedClothes, "1 x jacket")
	}
	return expectedClothes
}

func getRecommendationHandler(api weatherAPI) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		city := r.URL.Query().Get("location")
		fromTime, err := time.Parse("2006-01-02", from)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		toTime, err := time.Parse("2006-01-02", to)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		rec := getRecommendationForRanges(fromTime, toTime, api, city)
		json.NewEncoder(w).Encode(rec)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}

func createHttpMux(api weatherAPI) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheckHandler)
	mux.HandleFunc("/recommendation", getRecommendationHandler(api))
	return mux
}

func main() {
	apiKey := os.Getenv("VISUAL_CROSSING_WEATHER_KEY")
	api := visualCrossingWeatherAPI{ApiKey: apiKey}
	log.Fatal(http.ListenAndServe(HostUrl, createHttpMux(api)))
}
