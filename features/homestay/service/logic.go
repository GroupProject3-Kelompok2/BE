package service

import (
	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/go-playground/validator"
)

type homestayService struct {
	homestayData homestay.HomestayDataInterface
	validate     *validator.Validate
}

func New(repo homestay.HomestayDataInterface) homestay.HomestayServiceInterface {
	return &homestayService{
		homestayData: repo,
		validate:     validator.New(),
	}
}

func (service *homestayService) Create(input homestay.HomestayCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.homestayData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
