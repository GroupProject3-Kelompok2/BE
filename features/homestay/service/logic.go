package service

import (
	"errors"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/go-playground/validator/v10"
)

var log = middlewares.Log()

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

func (service *homestayService) GetAll() ([]homestay.HomestayCore, error) {
	data, err := service.homestayData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *homestayService) GetById(homestayId string) (homestay.HomestayCore, error) {
	data, err := service.homestayData.SelectById(homestayId)
	if err != nil {
		return homestay.HomestayCore{}, err
	}
	return data, err
}

// HomestayPictures implements homestay.HomestayServiceInterface
func (hs *homestayService) HomestayPictures(homestayId string, req homestay.HomestayPictureCore) error {
	err := hs.homestayData.HomestayPictures(homestayId, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("homestay profile record not found")
			return errors.New("homestay profile record not found")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}
