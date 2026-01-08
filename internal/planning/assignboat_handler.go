package planning

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func AssignBoat(c *gin.Context) {
	// Parse request to get survivorID and boatID
	survivor_id := c.Query("survivor_id")              // returns "" if not present
    boat_id  := c.Query("boat_id")

	msg := fmt.Sprintf("Boat assigned successfully: Assigning boat %s to survivor %s", boat_id, survivor_id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}