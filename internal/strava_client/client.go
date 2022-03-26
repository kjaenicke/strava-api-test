package stravaclient

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/kjaenicke/strava-api-test/internal/models"
)

const BASE_API_PATH = "https://www.strava.com/api/v3"

func getAccessToken() string {
	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
	
	if accessToken == "" {
		log.Fatal("No Strava access token found")
	}

	return accessToken
}

func GetCurrentAthlete() (*models.AthleteResponse, error) {
	var athlete models.AthleteResponse

	accessToken := getAccessToken()	
	client := resty.New()
	response, err := client.R().SetHeader("Authorization", "Bearer "+accessToken).Get(BASE_API_PATH + "/athlete")
	
	if err != nil {
		log.Printf("Got error %s", err.Error())
		return &athlete,  fmt.Errorf("Unable to decode athlete")
	}

	if err := json.Unmarshal(response.Body(), &athlete); err != nil {
		log.Printf("Got error %s", err.Error())
		return &athlete, fmt.Errorf("Unable to decode athlete")
	}

	return &athlete, nil
}

func GetActivities() (*models.ActivityResponse, error) {
	var activities models.ActivityResponse

	accessToken := getAccessToken()	
	client := resty.New()
	response, err := client.R().SetHeader("Authorization", "Bearer "+accessToken).Get(BASE_API_PATH + "/athlete/activities")
	
	if err != nil {
		log.Printf("Got error %s", err.Error())
		return &activities, fmt.Errorf("Unable to decode activities")
	}

	if err := json.Unmarshal(response.Body(), &activities); err != nil {
		log.Printf("Got error %s", err.Error())
		return &activities, fmt.Errorf("Unable to decode activities")
	}

	return &activities, nil
}