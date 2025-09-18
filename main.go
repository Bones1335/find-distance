package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Bones1335/find-distance/api"
	"github.com/Bones1335/find-distance/internal/createCsv"
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
	defaultCity := url.PathEscape(args[1] + " " + args[2])
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

	var calculatedDistances []api.Distance
	for _, internship := range internships {

		fixedCity := url.PathEscape(internship.FixedZipCode + " " + internship.FixedCity)

		currentCity := url.PathEscape(internship.CurrentZipCode + " " + internship.CurrentCity)

		destinationCity := url.PathEscape(internship.DestinationZipCode + " " + internship.DestinationCity)

		if defaultCity == fixedCity || defaultCity == currentCity || (defaultCity == fixedCity && defaultCity == currentCity) {

			lon1, lat1, err := api.GetGeocodeRequest(apiKey, defaultCity)
			if err != nil {
				fmt.Printf("error requesting data: %+v", err)
				return
			}

			lon2, lat2, err := api.GetGeocodeRequest(apiKey, destinationCity)
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

			calculatedDistances = append(calculatedDistances, api.Distance{
				LastName:  internship.LastName,
				FirstName: internship.FirstName,
				Distance:  ((dist / 1000) * 2),
			})

			fmt.Printf("%s\n", defaultCity)

			continue
		}

		lon1, lat1, err := api.GetGeocodeRequest(apiKey, fixedCity)
		if err != nil {
			fmt.Printf("error requesting data: %+v\n", err)
			return
		}

		lon2, lat2, err := api.GetGeocodeRequest(apiKey, currentCity)
		if err != nil {
			fmt.Printf("error requesting data: %+v\n", err)
			return
		}

		lon3, lat3, err := api.GetGeocodeRequest(apiKey, destinationCity)
		if err != nil {
			fmt.Printf("error requesting data: %+v\n", err)
			return
		}

		fixedDestination := [][]float64{{lon1, lat1}, {lon3, lat3}}
		currentDestination := [][]float64{{lon2, lat2}, {lon3, lat3}}

		dist1, err := api.PostDirectionsRequest(apiKey, fixedDestination)
		if err != nil {
			fmt.Println("Error getting distance", err)
			return
		}

		dist2, err := api.PostDirectionsRequest(apiKey, currentDestination)
		if err != nil {
			fmt.Println("Error getting distance", err)
			return
		}

		if dist1 < dist2 {
			calculatedDistances = append(calculatedDistances, api.Distance{
				LastName:  internship.LastName,
				FirstName: internship.FirstName,
				Distance:  ((dist1 / 1000) * 2),
			})

			fmt.Printf("%s\n", fixedCity)

		} else {

			calculatedDistances = append(calculatedDistances, api.Distance{
				LastName:  internship.LastName,
				FirstName: internship.FirstName,
				Distance:  ((dist2 / 1000) * 2),
			})

			fmt.Printf("%s\n", currentCity)

		}
	}

	createCsv.CreateCsv(calculatedDistances)
}
