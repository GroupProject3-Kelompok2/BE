package data

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ReservationID      string `gorm:"primaryKey;type:varchar(50)"`
	Invoice            string `gorm:"primaryKey;type:varchar(50)"` // a.k.a PaymentID
	HomestayHomestayID string `gorm:"primaryKey;type:varchar(50)"`
	Quantity           uint
	Subtotal           uint
	CreatedAt          time.Time      `gorm:"type:datetime"`
	UpdatedAt          time.Time      `gorm:"type:datetime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
