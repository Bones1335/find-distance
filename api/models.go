package api

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

type OpenRouteService struct {
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

type Directions struct {
	Routes []Route `json:"routes"`
}
type Route struct {
	Summary Summary `json:"summary"`
}

type Summary struct {
	Distance float64 `json:"distance"`
}
