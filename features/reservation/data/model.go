package data

import (
	"time"

	"gorm.io/gorm"
)

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
	User          User           `gorm:"foreignKey:UserID"`
	Homestay      Homestay       `gorm:"foreignKey:HomestayID"`
	PaymentID     string
	Payment       Payment `gorm:"foreignKey:PaymentID"`
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
	Homestays       []Homestay     `gorm:"foreignKey:UserID"`
}

type Homestay struct {
	HomestayID   string         `gorm:"primaryKey;type:varchar(21)"`
	UserID       string         `gorm:"type:varchar(21)"`
	Name         string         `gorm:"type:text"`
	Description  string         `gorm:"type:text"`
	Address      string         `gorm:"type:text"`
	Status       bool           `gorm:"type:boolean"`
	CreatedAt    time.Time      `gorm:"type:datetime"`
	UpdatedAt    time.Time      `gorm:"type:datetime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	User         User           `gorm:"foreignKey:UserID"`
	Reservations []Reservation  `gorm:"foreignKey:HomestayID"`
}

type Payment struct {
	PaymentID     string `gorm:"primaryKey;type:varchar(255)"`
	ReservationID string `gorm:"type:varchar(21)"`
	Amount        string
	BankAccount   string         `gorm:"type:enum('bca', 'bri', 'bni', 'mandiri'); default:'bca'"`
	VANumber      string         `gorm:"type:varchar(21)"`
	Status        string         `gorm:"type:varchar(21)"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
