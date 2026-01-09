package viz

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/illusioniststg/hopeline/internal/domain"
)

func GetKPISummary(c *gin.Context) {
	summary := domain.KPISummaryResponse{
		Rescued:          1248,
		ConfirmedSafe:    980,
		Missing:          312,
		BoatsAvailable:   23,
		BoatsTotal:       41,
		ShelterOccupancy: 68.4,
		LastUpdated:      time.Now().String(),
	}

	c.JSON(http.StatusOK, summary)
}
