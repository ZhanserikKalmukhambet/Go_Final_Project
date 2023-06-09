package models

import (
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/enums"
	"gorm.io/gorm"
)

type Complaint struct {
	gorm.Model
	UserID          int                   `json:"userID"`
	AdvertisementID int                   `json:"advertisementID"`
	ComplaintReason enums.ComplaintReason `json:"complaintReason"`
	Text            string                `json:"text"`
}
