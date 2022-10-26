package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

func Test_ServerIsHealthy(t *testing.T) {
	//Arrange
	expectedStatus := "healthy"
	//Act
	server := httptest.NewServer(createHttpMux(WeatherAPIMock{}))
	defer server.Close()
	res, err := http.Get(server.URL + "/health")
	// Assert
	if err != nil {
		t.Error(err)
	}
	status, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Should be %d", http.StatusOK)
	}
	if string(status) != expectedStatus {
		t.Errorf("Should be %s", expectedStatus)
	}
}

func TestBasicClothesForDateRanges(t *testing.T) {
	// Arrange
	testCases := []struct {
		From               string
		To                 string
		ExpectedStatusCode int
		ExpectedResponse   []string
	}{
		{
			From:               "2022-10-26",
			To:                 "2022-10-27",
			ExpectedStatusCode: 200,
			ExpectedResponse:   []string{"1 x tshirt", "1 x pair of socks", "1 x panties"},
		},
		{
			From:               "202-10-26",
			To:                 "202-10-26",
			ExpectedStatusCode: 400,
			ExpectedResponse:   []string{},
		},
		{
			From:               "2022-10-26",
			To:                 "2022-10-28",
			ExpectedStatusCode: 200,
			ExpectedResponse:   []string{"2 x tshirt", "2 x pair of socks", "2 x panties"},
		},
	}
	server := httptest.NewServer(createHttpMux(WeatherAPIMock{}))
	defer server.Close()
	// Act
	for _, c := range testCases {
		httpResponse, err := http.Get(server.URL + fmt.Sprintf("/recommendation?from=%s&to=%s", c.From, c.To))
		if err != nil {
			t.Error(err)
		}

		response, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			t.Error("Oops")
		}

		t.Logf("%v", httpResponse.StatusCode)
		if httpResponse.StatusCode != c.ExpectedStatusCode {
			t.FailNow()
		}
		data := []string{}
		json.Unmarshal(response, &data)
		// Assert
		if !slices.Equal(data, c.ExpectedResponse) {
			t.Errorf("Should be %v, but got %s", c.ExpectedResponse, response)
		}
	}
}

func TestWeatherApi(t *testing.T) {
	// Arrange
	weatherAPI := visualCrossingWeatherAPI{ApiKey: os.Getenv("VISUAL_CROSSING_WEATHER_KEY")}
	from, _ := time.Parse("2006-01-02", "2022-01-02")
	middleDay, _ := time.Parse("2006-01-02", "2022-01-03")
	to, _ := time.Parse("2006-01-02", "2022-01-04")
	// Act
	temperatures := weatherAPI.GetTemperatures(from, to, "oslo")
	// Assert
	if val, ok := temperatures[from]; !ok || val != 0 {
		t.Errorf("Either the day '%v' was not given a temperature or it was the wrong one, expected 0, got %d", from, val)
	}
	if val, ok := temperatures[middleDay]; !ok || val != 3 {
		t.Errorf("Either the day '%v' was not given a temperature or it was the wrong one, expected 3, got %d", middleDay, val)
	}
	if val, ok := temperatures[to]; !ok || val != 0 {
		t.Errorf("Either the day '%v' was not given a temperature or it was the wrong one, expected 0, got %d", to, val)
	}
}

func TestTemperatureBelow11DegreesShouldReturnJacket(t *testing.T) {
	// Arrange
	api := WeatherAPIMock{ExpectedResult: map[time.Time]int{time.Now(): 1}}
	server := httptest.NewServer(createHttpMux(api))
	defer server.Close()
	// Act
	httpResponse, err := http.Get(server.URL + fmt.Sprint("/recommendation?from=2022-10-26&to=2022-10-28&location=oslo"))
	// Assert
	if err != nil {
		t.Error(err)
	}
	response, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		t.Error("Oops")
	}
	if httpResponse.StatusCode != 200 {
		t.FailNow()
	}
	data := []string{}
	json.Unmarshal(response, &data)
	expectedResponse := []string{"2 x tshirt", "2 x pair of socks", "2 x panties", "1 x jacket"}
	// Assert
	if !slices.Equal(data, expectedResponse) {
		t.Errorf("Should be %v, but got %s", expectedResponse, response)
	}
}

type WeatherAPIMock struct {
	ExpectedResult map[time.Time]int
}

func (m WeatherAPIMock) GetTemperatures(from, to time.Time, city string) map[time.Time]int {
	return m.ExpectedResult
}
