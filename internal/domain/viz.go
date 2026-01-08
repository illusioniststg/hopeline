package domain

type KPISummaryResponse struct {
	Rescued          int       `json:"rescued"`
	ConfirmedSafe    int       `json:"confirmed_safe"`
	Missing          int       `json:"missing"`
	BoatsAvailable   int       `json:"boats_available"`
	BoatsTotal       int       `json:"boats_total"`
	ShelterOccupancy float64   `json:"shelter_occupancy_percent"`
	LastUpdated      string    `json:"last_updated"`
}
