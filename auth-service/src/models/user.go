package models

type User struct {
	ProfileId int    `json:"profile_id"`
	Number    string `json:"number"`
	Name      string `json:"name"`
}
