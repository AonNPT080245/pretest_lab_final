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
	StudentID   string `gorm:"uniqurIndex" valid:"matches(^[B]\\d{7}$)~StudentID invalid"`
	PhoneNumber string `valid:"matches(^([+]66|0)+[0-9]{9}$)~PhoneNumber invalid"`
	Password    string `valid:"minstringlength(8)~password must be between 8 and 20 characters, maxstringlength(20)~password must be between 8 and 20 characters"`
}
