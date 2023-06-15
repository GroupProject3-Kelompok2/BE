package handler

import (
	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
)

type ReservationResponse struct {
	ReservationID string  `json:"reservation_id,omitempty"`
	HomestayName  string  `json:"homestay_name,omitempty"`
	CheckinDate   string  `json:"checkin_date,omitempty"`
	CheckOutDate  string  `json:"checkout_date,omitempty"`
	HomestayPrice float64 `json:"homestay_price,omitempty"`
	Duration      int     `json:"duration,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	BankAccount   string  `json:"bank_account,omitempty"`
	VaNumber      string  `json:"va_number,omitempty"`
	Status        string  `json:"status,omitempty"`
}

func NewReservationResponse(reservationID string) ReservationResponse {
	return ReservationResponse{
		ReservationID: reservationID,
	}
}

func ReservationResponseData(reservation reservation.ReservationCore) ReservationResponse {
	return ReservationResponse{
		ReservationID: reservation.ReservationID,
		HomestayName:  reservation.Homestay.Name,
		CheckinDate:   reservation.CheckinDate,
		CheckOutDate:  reservation.CheckoutDate,
		HomestayPrice: reservation.Homestay.Price,
		Duration:      reservation.Availability.ReservationDuration,
		Amount:        reservation.Availability.GrossAmount,
	}
}
