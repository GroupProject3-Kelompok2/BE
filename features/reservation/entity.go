package reservation

import "time"

type Reservation struct {
	ReservationID      string
	Invoice            string // a.k.a PaymentID
	HomestayHomestayID string
	Quantity           uint
	Subtotal           uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
}
