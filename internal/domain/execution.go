package domain

type ValidateInputRequest struct {
	BoatID string `json:"boat_id"`
	ShelterID string `json:"shelter_id"`
	RouteID string `json:"route_id"`
}