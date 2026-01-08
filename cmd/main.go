package main

import (
	 "github.com/gin-gonic/gin"
	 "github.com/illusioniststg/hopeline/internal/planning"
)

func pong(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong from server",
    })
  }

func main() {
	router := gin.Default()
	router.GET("/ping", pong)
	// http://localhost:8080/assign?survivor_id=123&boat_id=boatA
	router.GET("/assign", planning.AssignBoat)
	router.Run() // listens on 0.0.0.0:8080 by default
}
