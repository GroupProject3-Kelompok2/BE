package user

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/features/review"
	"github.com/labstack/echo/v4"
)

type UserCore struct {
	UserID          string
	Fullname        string
	Email           string
	Password        string
	ProfilePricture string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	Homestays       []homestay.HomestayCore
	Reviews         []review.ReviewCore
	Payments        []payment.PaymentCore
}

type UserHandler interface {
	Register() echo.HandlerFunc
}

type UserService interface {
	Register(request UserCore) (UserCore, error)
}

type UserData interface {
	Register(request UserCore) (UserCore, error)
}
