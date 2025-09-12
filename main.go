package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
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

type Coordinates struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Type     string   `json:"type"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func main() {
	SetEnv(".env")

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

	internships := []*Internship{}

	err = gocsv.Unmarshal(file, &internships)
	if err != nil {
		log.Fatal(err)
	}

	err = Request(apiKey)
	if err != nil {
		fmt.Printf("error requesting data: %+v", err)
		return
	}

	//	for _, internship := range internships {
	//		fmt.Println("Hello", internship.FirstName, internship.LastName)
	//	}

}

func SetEnv(filename string) error {
	if filename != ".env" {
		return fmt.Errorf("can't load non .env file")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, "=", 2)
		if len(splitLine) != 2 {
			return fmt.Errorf("current line variable in .env causing errors: %v", line)
		}

		err = os.Setenv(splitLine[0], strings.ReplaceAll(splitLine[1], "\"", ""))
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Request(apiKey string) error {
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
	coors := Coordinates{}
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

func (c *Coordinates) GetCoordinates() (lon, lat float64) {
	if len(c.Features) > 0 && len(c.Features[0].Geometry.Coordinates) >= 2 {
		return c.Features[0].Geometry.Coordinates[0], c.Features[0].Geometry.Coordinates[1]
	}

	return 0, 0
}
