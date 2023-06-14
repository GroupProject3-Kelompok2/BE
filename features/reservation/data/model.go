package data

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
	"gorm.io/gorm"
)

type Reservation struct {
	ReservationID string         `gorm:"primaryKey;type:varchar(21)"`
	UserID        string         `gorm:"primaryKey;type:varchar(21)"`
	HomestayID    string         `gorm:"primaryKey;type:varchar(21)"`
	CheckInDate   string         `gorm:"type:date"`
	CheckOutDate  string         `gorm:"type:date"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	User          User           `gorm:"foreignKey:UserID"`
	Homestay      Homestay       `gorm:"foreignKey:HomestayID"`
	PaymentID     string
	Payment       Payment `gorm:"foreignKey:PaymentID"`
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

func ReservationCore(reservationData Reservation) reservation.ReservationCore {
	return reservation.ReservationCore{
		ReservationID: reservationData.ReservationID,
		UserID:        reservationData.UserID,
		HomestayID:    reservationData.HomestayID,
		CheckinDate:   reservationData.CheckInDate,
		CheckoutDate:  reservationData.CheckOutDate,
		CreatedAt:     reservationData.CreatedAt,
		UpdatedAt:     reservationData.UpdatedAt,
	}
}

func ReservationModel(dataCore reservation.ReservationCore) Reservation {
	return Reservation{
		ReservationID: dataCore.ReservationID,
		UserID:        dataCore.UserID,
		HomestayID:    dataCore.HomestayID,
		CheckInDate:   dataCore.CheckinDate,
		CheckOutDate:  dataCore.CheckoutDate,
		CreatedAt:     dataCore.CreatedAt,
		UpdatedAt:     dataCore.UpdatedAt,
	}
}
