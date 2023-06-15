package handler

import (
	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
)

type ReservationResponse struct {
	ReservationID       string  `json:"reservation_id,omitempty"`
	HomestayName        string  `json:"homestay_name,omitempty"`
	CheckinDate         string  `json:"checkin_date,omitempty"`
	CheckOutDate        string  `json:"checkout_date,omitempty"`
	HomestayPrice       float64 `json:"homestay_price,omitempty"`
	ReservationDuration int     `json:"reservation_duration,omitempty"`
	GrossAmount         float64 `json:"gross_amount,omitempty"`
}

func NewReservationResponse(reservationID string) ReservationResponse {
	return ReservationResponse{
		ReservationID: reservationID,
	}
}

func ReservationResponseData(reservation reservation.ReservationCore, homestay reservation.Homestay, availability reservation.Availability) ReservationResponse {
	return ReservationResponse{
		HomestayName:        homestay.Name,
		CheckinDate:         reservation.CheckinDate,
		CheckOutDate:        reservation.CheckoutDate,
		HomestayPrice:       homestay.Price,
		ReservationDuration: availability.ReservationDuration,
		GrossAmount:         availability.GrossAmount,
	}
}
