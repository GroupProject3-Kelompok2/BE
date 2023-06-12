package data

import (
	"time"

	"gorm.io/gorm"
)

type HomestayPicture struct {
	HomestayPictureID  string         `gorm:"primaryKey;type:varchar(21)"`
	HomestayID         string         `gorm:"type:varchar(21)"`
	HomestayPictureURL string         `gorm:"type:text;not null"`
	CreatedAt          time.Time      `gorm:"type:datetime"`
	UpdatedAt          time.Time      `gorm:"type:datetime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Homestay           Homestay       `gorm:"foreignKey:HomestayID"`
}

type Homestay struct {
	HomestayID       string            `gorm:"primaryKey;type:varchar(21)"`
	UserID           string            `gorm:"type:varchar(21)"`
	Name             string            `gorm:"type:text"`
	Description      string            `gorm:"type:text"`
	Address          string            `gorm:"type:text"`
	Status           bool              `gorm:"type:boolean"`
	CreatedAt        time.Time         `gorm:"type:datetime"`
	UpdatedAt        time.Time         `gorm:"type:datetime"`
	DeletedAt        gorm.DeletedAt    `gorm:"index"`
	HomestayPictures []HomestayPicture `gorm:"foreignKey:HomestayID"`
}
