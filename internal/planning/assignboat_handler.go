package planning

import (
	"fmt"
	"net/http"

	"github.com/illusioniststg/hopeline/internal/data"
	"github.com/illusioniststg/hopeline/internal/domain"

	"github.com/gin-gonic/gin"
)

var boats = []domain.Boat{
	{
		ID:           "B1",
		Name:         "Rescue Alpha",
		Capacity:     6,
		Availability: 6,
		IsAssigned:   false,
	},
	{
		ID:           "B2",
		Name:         "Inflatable Bravo",
		Capacity:     4,
		Availability: 4,
		IsAssigned:   false,
	},
	{
		ID:           "B3",
		Name:         "Motor Charlie",
		Capacity:     10,
		Availability: 0,
		IsAssigned:   true,
	},
}

func AssignBoat(c *gin.Context) {

	// 1. Read Request (Search Survivor)
	var req domain.SurvivorSearchRequest // return phone number in request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	survivor, exists := data.Survivors[req.Phone]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Survivor not found"})
		return
	}

	msg := fmt.Sprintf("Search results for person: %+v", survivor.FirstName+" "+survivor.LastName)
	c.JSON(http.StatusOK, gin.H{"message": msg})

	// 3. Select Available Boat (simple logic)
	var selectedBoat *domain.Boat

	for i := range boats {
		if boats[i].IsAssigned {
			continue
		}

		// assume 1 survivor = 1 seat for now
		if boats[i].Availability >= 1 {
			selectedBoat = &boats[i]
			break
		}
	}

	if selectedBoat == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "no available boat",
		})
		return
	}

	// 4. Assign Boat
	selectedBoat.IsAssigned = true
	selectedBoat.Availability = selectedBoat.Availability - 1

	// simulate DB update
	survivor.BoatID = 1

	// 5. Final Response (ONLY ONE RESPONSE)
	c.JSON(http.StatusOK, gin.H{
		"status": "ASSIGNED",
		"survivor": gin.H{
			"name":  survivor.FirstName + " " + survivor.LastName,
			"phone": survivor.Phone,
		},
		"boat": gin.H{
			"id":           selectedBoat.ID,
			"name":         selectedBoat.Name,
			"capacity":     selectedBoat.Capacity,
			"availability": selectedBoat.Availability,
		},
	})
}
