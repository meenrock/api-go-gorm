package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	NickName     string `json:"nick_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Weight       int
	Height       int
	Age          int
	Sex          string `json:"sex"`
	AllergicFood string `json:"allergic_food"`
	FavFood      string `json:"fav_food"`
	ExpectedBMI  int
}
