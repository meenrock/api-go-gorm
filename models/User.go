package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id           int       `gorm:"column:id; primary_key serial"`
	UserUuid     uuid.UUID `gorm:"column:user_uuid; primary_key default:uuid_generate_v4()"`
	FirstName    string    `gorm:"column:first_name;"`
	LastName     string    `gorm:"column:last_name;"`
	NickName     string    `gorm:"column:nick_name;"`
	Email        string    `gorm:"column:email;"`
	PhoneNumber  string    `gorm:"column:phone_number;"`
	Weight       float32   `gorm:"column:weight;"`
	Height       float32   `gorm:"column:height;"`
	Age          int       `gorm:"column:age;"`
	Sex          string    `gorm:"column:sex;"`
	AllergicFood string    `gorm:"column:allergic_food;"`
	FavFood      string    `gorm:"column:fav_food;"`
	ExpectedBMI  float32   `gorm:"column:expected_bmi;"`
}
