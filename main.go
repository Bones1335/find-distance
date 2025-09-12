package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

type Internship struct {
	LastName           string `csv:"last_name"`
	FirstName          string `csv:"first_name"`
	FixedZipCode       string `csv:"fixed_zip_code"`
	FixedCity          string `csv:"fixed_city"`
	CurrentZipCode     string `csv:"current_zip_code"`
	CurrentCity        string `csv:"current_city"`
	DestinationZipCode string `csv:"destination_zip_code"`
	DestinationCity    string `csv:"destination_city"`
}

func main() {
	args := os.Args

	defaultCity := args[1:3]
	fmt.Println(defaultCity)

	fileName := args[3]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	internships := []*Internship{}

	err = gocsv.Unmarshal(file, &internships)
	if err != nil {
		log.Fatal(err)
	}

	for _, internship := range internships {
		fmt.Println("Hello", internship.FirstName, internship.LastName)
	}
}
