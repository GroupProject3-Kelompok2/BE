package homestay

import (
	"time"
)

type HomestayCore struct {
	HomestayID   string
	Hoster       string  `validate:"required"`
	Name         string  `validate:"required"`
	Descriptions string  `validate:"required"`
	Address      string  `validate:"required"`
	Price        float64 `validate:"required"`
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type HomestayDataInterface interface {
	Insert(input HomestayCore) (uint64, error)
}

type HomestayServiceInterface interface {
	Create(input HomestayCore) (HomestayCore, error)
}
