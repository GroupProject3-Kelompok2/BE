package data

import (
	"time"

	homestay "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	"gorm.io/gorm"
)

type Payment struct {
	PaymentID string              `gorm:"primaryKey;type:varchar(50)"`
	CreatedAt time.Time           `gorm:"type:datetime"`
	UpdatedAt time.Time           `gorm:"type:datetime"`
	DeletedAt gorm.DeletedAt      `gorm:"index"`
	Buyer     string              `gorm:"type:varchar(50)"`
	Item      []homestay.Homestay `gorm:"many2many:Reservation;foreignKey:PaymentID;joinForeignKey:Invoice"`
}
