package api

func (ors *OpenRouteService) GetCoordinates() (lon, lat float64) {
	if len(ors.Features) > 0 && len(ors.Features[0].Geometry.Coordinates) >= 2 {
		return ors.Features[0].Geometry.Coordinates[0], ors.Features[0].Geometry.Coordinates[1]
	}

	return 0, 0
}
