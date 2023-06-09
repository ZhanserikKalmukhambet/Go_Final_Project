package models

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`

	Role   enums.Role   `json:"role" gorm:"default:User"`
	Gender enums.Gender `json:"gender"`

	FavoriteItems  []FavoriteItem  `json:"favoriteItems" gorm:"foreignKey:UserID"`
	Chats          []Chat          `json:"chats" gorm:"foreignKey:UserID"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:UserID"`
	Complaints     []Complaint     `json:"complaints" gorm:"foreignKey: UserID"`
}

type SignUpInput struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`

	Gender enums.Gender `json:"gender"`
}

type SignInInput struct {
	PhoneNumber string `json:"phoneNumber" gorm:"unique"`
	Password    string `json:"password"`
}
