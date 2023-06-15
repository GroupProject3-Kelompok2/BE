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
