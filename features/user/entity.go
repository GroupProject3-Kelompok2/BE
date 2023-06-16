package user

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/features/review"
	"github.com/labstack/echo/v4"
)

type UserCore struct {
	UserID         string
	Fullname       string `validate:"required"`
	Email          string `validate:"required,email"`
	Phone          string `validate:"required"`
	Password       string `validate:"required"`
	ProfilePicture string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Homestays      []homestay.HomestayCore
	Reviews        []review.ReviewCore
	Payments       []payment.PaymentCore
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	ProfileUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	UpgradeUser() echo.HandlerFunc
	DeactiveUser() echo.HandlerFunc
}

type UserService interface {
	Register(request UserCore) (UserCore, error)
	Login(request UserCore) (UserCore, string, error)
	ProfileUser(userId string) (UserCore, error)
	UpdateProfile(userId string, request UserCore) error
	UpgradeProfile(userId string, request UserCore) error
	DeactiveUser(userId string) error
}

type UserData interface {
	Register(request UserCore) (UserCore, error)
	Login(request UserCore) (UserCore, string, error)
	ProfileUser(userId string) (UserCore, error)
	UpdateProfile(userId string, request UserCore) error
	UpgradeProfile(userId string, request UserCore) error
	DeactiveUser(userId string) error
}
