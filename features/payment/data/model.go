package data

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	PaymentID     string `gorm:"primaryKey;type:varchar(21)"`
	ReservationID string `gorm:"type:varchar(21)"`
	Amount        uint
	BankAccount   string         `gorm:"type:enum('bca', 'bri', 'bni', 'mandiri'); default:'bca'"`
	VANumber      uint           `gorm:"type:varchar(21)"`
	Status        bool           `gorm:"type:boolean"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Reservation   Reservation    `gorm:"foreignKey:PaymentID;references:ReservationID"`
}

type Reservation struct {
	ReservationID string `gorm:"primaryKey;type:varchar(21)"`
	UserID        string `gorm:"primaryKey;type:varchar(21)"`
	HomestayID    string `gorm:"primaryKey;type:varchar(21)"`
	CheckInDate   time.Time
	CheckOutDate  time.Time
	Duration      uint8
	Price         uint
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
