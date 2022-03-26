package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kjaenicke/strava-api-test/internal/controllers"
)

func CreateServer() *gin.Engine {
	server := gin.Default()

	server.GET("/ping", controllers.Ping)
	server.GET("/athlete", controllers.GetAthlete)
	server.GET("/activities", controllers.ListActivities)

	return server
}