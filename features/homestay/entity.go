package homestay

import (
	"time"
)

type HomestayCore struct {
	HomestayID  string  `validate:"required"`
	UserID      string  `validate:"required"`
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Address     string  `validate:"required"`
	Price       float64 `validate:"required"`
	Pictures    []HomestayPictureCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HomestayPictureCore struct {
	PictureID  string
	HomestayID string
	URL        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type HomestayDataInterface interface {
	Insert(input HomestayCore) error
	UpdateById(userId string, homestayId string, input HomestayCore) error
	DeleteById(userId string, homestayId string) error
	SelectAll() ([]HomestayCore, error)
	SelectById(homestayID string) (HomestayCore, error)
	HomestayPictures(homestayId string, req HomestayPictureCore) error
}

type HomestayServiceInterface interface {
	Create(input HomestayCore) error
	UpdateById(userId string, homestayId string, input HomestayCore) error
	DeleteById(userId string, homestayId string) error
	GetAll() ([]HomestayCore, error)
	GetById(homestayID string) (HomestayCore, error)
	HomestayPictures(homestayId string, req HomestayPictureCore) error
}
