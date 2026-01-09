package planning

type WaterRouteRequest struct {
	Lat1          float64 `json:"lat1" binding:"required"`
	Lon1          float64 `json:"lon1" binding:"required"`
	Lat2          float64 `json:"lat2" binding:"required"`
	Lon2          float64 `json:"lon2" binding:"required"`
	BoatSpeedKmph float64 `json:"boat_speed_kmph"`
}

type WaterRouteResponse struct {
	DistanceMeters float64 `json:"distance_meters"`
	DistanceKm     float64 `json:"distance_km"`
	ETAHours       float64 `json:"eta_hours,omitempty"`
	Algorithm      string  `json:"algorithm"`
	EarthModel     string  `json:"earth_model"`
}
