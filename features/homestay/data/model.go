package data

import (
	"time"

	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"
	"gorm.io/gorm"
)

type Homestay struct {
	HomestayID       string            `gorm:"primaryKey;type:varchar(50)"`
	CreatedAt        time.Time         `gorm:"type:datetime"`
	UpdatedAt        time.Time         `gorm:"type:datetime"`
	DeletedAt        gorm.DeletedAt    `gorm:"index"`
	Hoster           string            `gorm:"type:varchar(50)"`
	HomestayPictures []HomestayPicture `gorm:"foreignKey:HomestayID"`
	Reviews          []review.Review   `gorm:"foreignKey:HomestayID"`
}

type HomestayPicture struct {
	HomestayPictureID string         `gorm:"primaryKey;type:varchar(50)"`
	HomestayPicture   string         `gorm:"type:varchar(255);not null"`
	CreatedAt         time.Time      `gorm:"type:datetime"`
	UpdatedAt         time.Time      `gorm:"type:datetime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	HomestayID        string         `gorm:"type:varchar(50)"`
}
