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

func (service *homestayService) UpdateById(userId string, homestayId string, input homestay.HomestayCore) error {
	errUpdate := service.homestayData.UpdateById(userId, homestayId, input)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (service *homestayService) DeleteById(userId string, homestayId string) error {
	errUpdate := service.homestayData.DeleteById(userId, homestayId)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
