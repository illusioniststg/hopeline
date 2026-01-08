package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illusioniststg/hopeline/internal/data"
	"github.com/illusioniststg/hopeline/internal/identity"
	"github.com/illusioniststg/hopeline/internal/planning"
	"github.com/illusioniststg/hopeline/internal/execution"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong from server",
	})
}

func main() {
	data.ReadSurvivorsFromCSV("internal/data/basedata.csv")
	router := gin.Default()
	router.GET("/ping", pong)

	router.GET("/data/survivors", data.GetSurvivors)
	// http://localhost:8080/assign?survivor_id=123&boat_id=boatA
	router.GET("/assign", planning.AssignBoat)
	// http://localhost:8080/search
	router.POST("/search", identity.SearchPerson)
	// http://localhost:8080/search
	router.POST("/locs", identity.GetLocs)
	router.POST("/validate", execution.Validate)
	router.Run() // listens on 0.0.0.0:8080 by default
}
//branch created