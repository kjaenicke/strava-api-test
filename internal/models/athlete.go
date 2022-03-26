package models

type AthleteResponse struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Bio string `json:"bio"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	Shoes []ShoeResponse `json:"shoes"`
}

type ShoeResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
}