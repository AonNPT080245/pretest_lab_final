package entity

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

type User struct {
	gorm.Model
	Name        string `valid:"required~Name cannot be blank"`
	Email       string `valid:"email~Email invalid format"`
	StudentID   string `gorm:"uniqurIndex" valid:"mathces(^[B]\\d{7,7}$)"`
	PhoneNumber string `valid:"mathces(^([+]66|0)+[0-9]{9,9}$)"`
	Password    string `valid:"matches(^minstringlength(8)$)~Password is not less than 8 characters ,maxstringlength(20)~Password is not more than 20 characters"`
}
