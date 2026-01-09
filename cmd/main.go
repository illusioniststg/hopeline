package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illusioniststg/hopeline/internal/data"
	"github.com/illusioniststg/hopeline/internal/execution"
	"github.com/illusioniststg/hopeline/internal/identity"
	"github.com/illusioniststg/hopeline/internal/planning"
	"github.com/illusioniststg/hopeline/internal/viz"
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

	api := router.Group("/api")
	adminGroup := api.Group("/admin")
	{
		adminGroup.GET("/readyornot", data.GetSurvivors)
		adminGroup.GET("/start", data.GetSurvivors)
		adminGroup.GET("/stop", data.GetSurvivors)
		adminGroup.GET("/status", data.GetSurvivors)
	}

	dataGroup := api.Group("/data")
	{
		dataGroup.GET("/survivors", data.GetSurvivors)
	}

	planGroup := api.Group("/plan")
	{
		// http://localhost:8080/assign?phone=05792219224

		planGroup.POST("/assign", planning.AssignBoat)
		// http://localhost:8080/route
		planGroup.POST("/route", planning.MultiPickupRoute)
	}
	identityGroup := api.Group("/identity")
	{
		identityGroup.POST("/search", identity.SearchPerson)
		// http://localhost:8080/locs
		identityGroup.POST("/locs", identity.GetLocs)
	}
	execGroup := api.Group("/exec")
	{
		// http://localhost:8080/validate
		execGroup.POST("/validate", execution.Validate)
	}
	vizGroup := api.Group("/viz")
	{
		// http://localhost:8080/metrics/summary
		vizGroup.GET("/metrics/summary", viz.GetKPISummary)
	}

	router.Run() // listens on 0.0.0.0:8080 by default
}
