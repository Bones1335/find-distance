package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRequest(apiKey string) error {
	client := &http.Client{}

	url := fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%s&text=VESOUL", apiKey)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json, application/geo-json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to server")
		return err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	coors := OpenRouteService{}
	err = decoder.Decode(&coors)
	if err != nil {
		fmt.Println("error decoding response", err)
		return err
	}

	fmt.Println(res.Status)
	lon, lat := coors.GetCoordinates()

	fmt.Printf("Longitude: %f, Latitude: %f\n", lon, lat)

	return nil
}
