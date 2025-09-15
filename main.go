package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Bones1335/find-distance/api"
	"github.com/Bones1335/find-distance/internal/env"
	"github.com/gocarina/gocsv"
)

func main() {
	err := env.SetEnv(".env")
	if err != nil {
		fmt.Printf("couldn't set environment: %s\n", err)
		return
	}

	apiKey := os.Getenv("API_KEY")

	args := os.Args
	defaultCity := args[1:3]
	fmt.Println(defaultCity)

	fileName := args[3]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	internships := []*api.Internship{}

	err = gocsv.Unmarshal(file, &internships)
	if err != nil {
		log.Fatal(err)
	}

	for _, internship := range internships {

		fixedCity := url.PathEscape(internship.FixedZipCode + " " + internship.FixedCity)

		currentCity := url.PathEscape(internship.CurrentZipCode + " " + internship.CurrentCity)

		lon1, lat1, err := api.GetGeocodeRequest(apiKey, fixedCity)
		if err != nil {
			fmt.Printf("error requesting data: %+v", err)
			return
		}

		lon2, lat2, err := api.GetGeocodeRequest(apiKey, currentCity)
		if err != nil {
			fmt.Printf("error requesting data: %+v", err)
			return
		}

		cities := [][]float64{{lon1, lat1}, {lon2, lat2}}

		dist, err := api.PostDirectionsRequest(apiKey, cities)
		if err != nil {
			fmt.Println("Error getting distance", err)
			return
		}

		fmt.Println(dist / 1000)

		fmt.Println("Fixed City info: ", fixedCity)
		fmt.Println("Current City info: ", currentCity)
	}

}
