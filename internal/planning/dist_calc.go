package planning

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pymaxion/geographiclib-go/geodesic"
)

// HTTP handler
func DistanceAlgo(c *gin.Context) {
	var req WaterRouteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	distanceMeters := calculateKarneyDistance(req.Lat1, req.Lon1, req.Lat2, req.Lon2)

	resp := WaterRouteResponse{
		DistanceMeters: distanceMeters,
		DistanceKm:     distanceMeters / 1000,
		Algorithm:      "Karney Geodesic",
		EarthModel:     "WGS-84 Ellipsoid",
	}

	if req.BoatSpeedKmph > 0 {
		resp.ETAHours = resp.DistanceKm / req.BoatSpeedKmph
	}

	msg := fmt.Sprintf("Distance in Km: %+v", distanceMeters/1000)
	c.JSON(http.StatusOK, gin.H{"message": msg})

	c.JSON(http.StatusOK, resp)
}

// Karney geodesic calculation (WATER distance)
func calculateKarneyDistance(lat1, lon1, lat2, lon2 float64) float64 {
	g := geodesic.WGS84
	result := g.Inverse(lat1, lon1, lat2, lon2)
	return result.S12 // distance in meters
}
