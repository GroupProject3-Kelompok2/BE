package reservation

import (
	"time"
)

type ReservationCore struct {
	ReservationID string
	UserID        string `validate:"required"`
	HomestayID    string `validate:"required"`
	CheckinDate   string `validate:"required"`
	CheckoutDate  string `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Homestay      Homestay
	Availability  Availability
	Payment       Payment
}

type Availability struct {
	Status              bool
	ReservationDuration int
	GrossAmount         float64
}

type Homestay struct {
	Name  string
	Price float64
}

type Payment struct {
	BankAccount string
	VaNumber    string
	Status      string
}

type ReservationDataInterface interface {
	Insert(input ReservationCore) (string, error)
	CheckAvailability(input ReservationCore) (int64, error)
	SelectHomestay(homestayID string) (Homestay, error)
	SelectById(reservationID string) (ReservationCore, error)
	SelectAllByUserId(userID string) ([]ReservationCore, error)
}

type ReservationServiceInterface interface {
	Create(input ReservationCore) (string, error)
	CheckAvailability(input ReservationCore) (ReservationCore, error)
	GetById(reservationID string) (ReservationCore, error)
	GetAllByUserId(userID string) ([]ReservationCore, error)
}
