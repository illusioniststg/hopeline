package viz

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illusioniststg/hopeline/internal/data"
)

func GetLocations(c *gin.Context) {
	locations := data.GetLocations()
	c.JSON(http.StatusOK, locations)
}
