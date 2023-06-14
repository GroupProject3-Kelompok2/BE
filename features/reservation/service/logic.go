package service

import (
	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservation.ReservationDataInterface
	validate        *validator.Validate
}

func New(repo reservation.ReservationDataInterface) reservation.ReservationServiceInterface {
	return &reservationService{
		reservationData: repo,
		validate:        validator.New(),
	}
}

func (service *reservationService) Create(input reservation.ReservationCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := service.reservationData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
