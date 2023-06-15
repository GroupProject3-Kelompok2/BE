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
	CheckInDate   time.Time      `gorm:"type:date"`
	CheckOutDate  time.Time      `gorm:"type:date"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type ReservationData struct {
	ReservationID string
	HomestayName  string
	CheckInDate   time.Time
	CheckOutDate  time.Time
	HomestayPrice float64
	Duration      int
	Amount        float64
	BankAccount   string
	VaNumber      string
	Status        string
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

func ReservationDataCore(reservationData ReservationData) reservation.ReservationCore {
	return reservation.ReservationCore{
		ReservationID: reservationData.ReservationID,
		CheckinDate:   reservationData.CheckInDate,
		CheckoutDate:  reservationData.CheckOutDate,
		Homestay: reservation.Homestay{
			Name:  reservationData.HomestayName,
			Price: reservationData.HomestayPrice,
		},
		Availability: reservation.Availability{
			ReservationDuration: reservationData.Duration,
			GrossAmount:         reservationData.Amount,
		},
		Payment: reservation.Payment{
			BankAccount: reservationData.BankAccount,
			VaNumber:    reservationData.VaNumber,
			Status:      reservationData.Status,
		},
	}
}
