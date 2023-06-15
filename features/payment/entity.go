package payment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentCore struct {
	PaymentID     string
	ReservationID string `validate:"required"`
	Amount        string `validate:"required"`
	BankAccount   string `validate:"required"`
	VANumber      string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PaymentHandler interface {
	Payment() echo.HandlerFunc
	Notification() echo.HandlerFunc
}

type PaymentService interface {
	Payment(request PaymentCore) (PaymentCore, error)
	UpdatePayment(request PaymentCore) error
}

type PaymentData interface {
	Payment(request PaymentCore) (PaymentCore, error)
	UpdatePayment(request PaymentCore) error
}
