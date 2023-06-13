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
}
