package execution

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/illusioniststg/hopeline/internal/domain"
)

func Handover(c *gin.Context) {
	var req domain.ValidateInputRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO : Implement actual Handover logic here
	//handover will have the name, age, boat id, shelter id, route id
	msg:=fmt.Sprintf("Handover successful: Handing over person with name, age, boat %s, shelter %s, and route %s", req.BoatID, req.ShelterID, req.RouteID)

	// msg := fmt.Sprintf("Handover successful: Handing over boat %s, shelter %s, and route %s", req.BoatID, req.ShelterID, req.RouteID)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}