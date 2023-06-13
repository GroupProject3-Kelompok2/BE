package review

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ReviewCore struct {
	ReviewID   string
	UserID     string
	HomestayID string `validate:"required"`
	Review     string `validate:"required"`
	Rating     uint8
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Reviewer   string
	User       UserCore
	Homestay   HomestayCore
}

type UserCore struct {
	UserID      string
	Username    string
	Email       string
	Password    string
	UserPicture string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Homestays   []HomestayCore
	Reviews     []ReviewCore
}

type HomestayCore struct {
	HomestayID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Hoster     string
}

type ReviewHandler interface {
	AddReview() echo.HandlerFunc
	EditReview() echo.HandlerFunc
}

type ReviewService interface {
	AddReview(userId string, request ReviewCore) error
	EditReview(userId string, reviewId string, request ReviewCore) error
}

type ReviewData interface {
	AddReview(userId string, request ReviewCore) error
	EditReview(userId string, reviewId string, request ReviewCore) error
}
