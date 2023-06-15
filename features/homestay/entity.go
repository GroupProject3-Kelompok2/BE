package homestay

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/utils/pagination"
)

type HomestayCore struct {
	HomestayID    string  `validate:"required"`
	UserID        string  `validate:"required"`
	Name          string  `validate:"required"`
	Description   string  `validate:"required"`
	Address       string  `validate:"required"`
	Price         float64 `validate:"required"`
	TotalReviews  uint
	AverageRating float32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Pictures      []HomestayPictureCore
	Reviews       []ReviewCore
}

type HomestayPictureCore struct {
	PictureID  string
	HomestayID string
	URL        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

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
	Homestays      []HomestayCore
	Reviews        []ReviewCore
}

type HomestayDataInterface interface {
	Insert(input HomestayCore) error
	UpdateById(userId string, homestayId string, input HomestayCore) error
	DeleteById(userId string, homestayId string) error
	SelectAll(keyword string, page pagination.Pagination) ([]HomestayCore, error)
	SelectById(homestayID string) (HomestayCore, error)
	HomestayPictures(homestayId string, req HomestayPictureCore) error
}

type HomestayServiceInterface interface {
	Create(input HomestayCore) error
	UpdateById(userId string, homestayId string, input HomestayCore) error
	DeleteById(userId string, homestayId string) error
	GetAll(keyword string, page pagination.Pagination) ([]HomestayCore, error)
	GetById(homestayID string) (HomestayCore, error)
	HomestayPictures(homestayId string, req HomestayPictureCore) error
}
