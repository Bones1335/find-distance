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

	lon1, lat1, err := api.GetGeocodeRequest(apiKey, url.PathEscape("70000 VESOUL"))
	if err != nil {
		fmt.Printf("error requesting data: %+v", err)
		return
	}
	fmt.Printf("Vesoul => Longitude: %f, Latitude: %f\n", lon1, lat1)

	lon2, lat2, err := api.GetGeocodeRequest(apiKey, url.PathEscape("25000 BESANCON"))
	if err != nil {
		fmt.Printf("error requesting data: %+v", err)
		return
	}
	fmt.Printf("Besak => Longitude: %f, Latitude: %f\n", lon2, lat2)

	cities := [][]float64{{lon1, lat1}, {lon2, lat2}}

	dist, err := api.PostDirectionsRequest(apiKey, cities)
	if err != nil {
		fmt.Println("Error getting distance", err)
		return
	}

	fmt.Println(dist / 1000)

	//	for _, internship := range internships {
	//		fmt.Println("Hello", internship.FirstName, internship.LastName)
	//	}

}
