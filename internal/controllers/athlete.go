package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stravaclient "github.com/kjaenicke/strava-api-test/internal/strava_client"
)


func GetAthlete(c *gin.Context) {
	athlete, err := stravaclient.GetCurrentAthlete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, athlete)
}