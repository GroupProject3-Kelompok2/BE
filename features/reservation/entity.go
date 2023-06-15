package reservation

import "time"

type ReservationCore struct {
	ReservationID string
	UserID        string `validate:"required"`
	HomestayID    string `validate:"required"`
	CheckinDate   string `validate:"required"`
	CheckoutDate  string `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

type ReservationDataInterface interface {
	Insert(input ReservationCore) (string, error)
	CheckAvailability(input ReservationCore) (int64, error)
	SelectHomestay(homestayID string) (Homestay, error)
	SelectById(reservationID string) (ReservationCore, error)
	//SelectAllByUser(userID string) ([]ReservationCore, error)
}

type ReservationServiceInterface interface {
	Create(input ReservationCore) (string, error)
	CheckAvailability(input ReservationCore) (Availability, Homestay, error)
	GetById(reservationID string) (ReservationCore, Homestay, Availability, error)
	//GetAllByUser(userID string) ([]ReservationCore, error)
}
