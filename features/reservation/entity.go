package reservation

import "time"

type ReservationCore struct {
	ReservationID string `validate:"required"`
	UserID        string `validate:"required"`
	HomestayID    string `validate:"required"`
	CheckinDate   string `validate:"required"`
	CheckoutDate  string `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ReservationDataInterface interface {
	Insert(input ReservationCore) error
}

type ReservationServiceInterface interface {
	Create(input ReservationCore) error
}
