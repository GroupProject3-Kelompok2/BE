package homestay

import (
	"time"
)

type HomestayCore struct {
	HomestayID  string
	UserID      string  `validate:"required"`
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Address     string  `validate:"required"`
	Price       float64 `validate:"required"`
	Status      bool
	Pictures    HomestayPictureCore
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
}

type HomestayServiceInterface interface {
	Create(input HomestayCore) error
}
