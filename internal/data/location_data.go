package data

import "github.com/illusioniststg/hopeline/internal/domain"

func GetLocations() []domain.Location {
	return []domain.Location{
		{Latitude: 26.1445, Longitude: 91.7362}, // Guwahati
		{Latitude: 27.4728, Longitude: 94.9120}, // Dibrugarh
		{Latitude: 26.7509, Longitude: 94.2037}, // Jorhat
		{Latitude: 26.9853, Longitude: 94.6361}, // Sivasagar
		{Latitude: 26.2006, Longitude: 92.9376}, // Central Assam
	}
}
