package api

func (ors *OpenRouteService) GetCoordinates() (lon, lat float64) {
	if len(ors.Features) > 0 && len(ors.Features[0].Geometry.Coordinates) >= 2 {
		return ors.Features[0].Geometry.Coordinates[0], ors.Features[0].Geometry.Coordinates[1]
	}

	return 0, 0
}

func (d *Directions) GetDistance() (dist float64) {
	if len(d.Routes) > 0 && d.Routes[0].Summary.Distance >= 1 {
		return d.Routes[0].Summary.Distance
	}

	return 0
}
