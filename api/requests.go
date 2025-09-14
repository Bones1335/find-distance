package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGeocodeRequest(apiKey, city string) (float64, float64, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%s&text=%s", apiKey, city)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json, application/geo-json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to server")
		return 0, 0, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	coors := OpenRouteService{}
	err = decoder.Decode(&coors)
	if err != nil {
		fmt.Println("error decoding response", err)
		return 0, 0, err
	}

	fmt.Println(res.Status)
	lon, lat := coors.GetCoordinates()

	return lon, lat, nil
}

func PostDirectionsRequest(apiKey string, cities [][]float64) (dist float64, err error) {
	client := &http.Client{}

	body := fmt.Sprintf("{\"coordinates\":[[%f,%f],[%f,%f]]}", cities[0][0], cities[0][1], cities[1][0], cities[1][1])
	fmt.Println(body)

	req, err := http.NewRequest("POST", "https://api.openrouteservice.org/v2/directions/driving-car", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Printf("error requesting directions: %s", err)
		return 0, err
	}

	req.Header.Add("Accept", "application/json, application/geo+json; charset=utf-8")
	req.Header.Add("Authorization", apiKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return 0, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	distance := Directions{}
	err = decoder.Decode(&distance)
	if err != nil {
		fmt.Println("error decoding response", err)
		return 0, err
	}
	fmt.Println(res.Status)
	dist = distance.GetDistance()

	return dist, nil
}
