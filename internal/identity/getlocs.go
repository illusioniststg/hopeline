package identity

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illusioniststg/hopeline/internal/data"
	"github.com/illusioniststg/hopeline/internal/domain"
)

func GetLocs(c *gin.Context) {
	var req domain.SurvivorSearchRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	survivor, exists := data.Survivors[req.Phone]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Survivor not found"})
		return
	}

	msg := fmt.Sprintf("Search results for person: %+v", survivor.Locs)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
