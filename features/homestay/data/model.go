package data

import (
	"time"

	homestay_picture "github.com/GroupProject3-Kelompok2/BE/features/homestay_picture/data"
	reservation "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"

	"gorm.io/gorm"
)

type Homestay struct {
	HomestayID       string                             `gorm:"primaryKey;type:varchar(21)"`
	UserID           string                             `gorm:"type:varchar(21)"`
	Name             string                             `gorm:"type:text"`
	Description      string                             `gorm:"type:text"`
	Address          string                             `gorm:"type:text"`
	Status           bool                               `gorm:"type:boolean"`
	CreatedAt        time.Time                          `gorm:"type:datetime"`
	UpdatedAt        time.Time                          `gorm:"type:datetime"`
	DeletedAt        gorm.DeletedAt                     `gorm:"index"`
	User             User                               `gorm:"foreignKey:UserID"`
	HomestayPictures []homestay_picture.HomestayPicture `gorm:"foreignKey:HomestayID"`
	Reservations     []reservation.Reservation          `gorm:"foreignKey:HomestayID"`
	Reviews          []review.Review                    `gorm:"foreignKey:HomestayID"`
}

type User struct {
	UserID          string         `gorm:"primaryKey;type:varchar(21)"`
	Fullname        string         `gorm:"type:varchar(100);not null"`
	Email           string         `gorm:"type:varchar(100);not null;unique"`
	Password        string         `gorm:"type:text"`
	ProfilePricture string         `gorm:"type:text"`
	Role            string         `gorm:"type:enum('user', 'hoster'); default:'user'"`
	CreatedAt       time.Time      `gorm:"type:datetime"`
	UpdatedAt       time.Time      `gorm:"type:datetime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
