package models

type User struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	City       string `json:"city"`
	Country    string `json:"country"`
	State      string `json:"state"`
	Street     string `json:"street"`
	Profile    string `json:"profile "`
}
