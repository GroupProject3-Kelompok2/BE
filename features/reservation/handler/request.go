package handler

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
)

type ReservationRequest struct {
	HomestayID   string    `json:"homestay_id" form:"homestay_id"`
	CheckInDate  time.Time `json:"checkin_date" form:"checkin_date"`
	CheckOutDate time.Time `json:"checkout_date" form:"checkout_date"`
}

func ReservationRequestCore(reservationRequest ReservationRequest) reservation.ReservationCore {
	return reservation.ReservationCore{
		HomestayID:   reservationRequest.HomestayID,
		CheckinDate:  reservationRequest.CheckInDate,
		CheckoutDate: reservationRequest.CheckOutDate,
	}
}
