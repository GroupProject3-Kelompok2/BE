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
	Price            float64                            `gorm:"type:decimal(15,2)"`
	Status           bool                               `gorm:"type:boolean;default false"`
	CreatedAt        time.Time                          `gorm:"type:datetime"`
	UpdatedAt        time.Time                          `gorm:"type:datetime"`
	DeletedAt        gorm.DeletedAt                     `gorm:"index"`
	HomestayPictures []homestay_picture.HomestayPicture `gorm:"foreignKey:HomestayID"`
	Reservations     []reservation.Reservation          `gorm:"foreignKey:HomestayID"`
	Reviews          []review.Review                    `gorm:"foreignKey:HomestayID"`
}
