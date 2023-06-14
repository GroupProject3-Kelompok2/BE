package handler

import "time"

type ReservationResponse struct {
	ReservationID string
	HomestayID    string
	CheckinDate   time.Time
	CheckOutDate  time.Time
}
