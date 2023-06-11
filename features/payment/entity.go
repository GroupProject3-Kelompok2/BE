package payment

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
)

type PaymentCore struct {
	PaymentID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Buyer     string
	Item      []homestay.HomestayCore
}
