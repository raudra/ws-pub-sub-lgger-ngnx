package models

type Otp struct {
	Number string `json:"number"`
	Otp    int    `json:"int"`
	Count  int    `json:"-"`
}
