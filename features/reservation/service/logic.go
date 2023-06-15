package service

import (
	"fmt"
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
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

func (service *reservationService) Create(input reservation.ReservationCore) (string, error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return "", errValidate
	}

	input.ReservationID, _ = helper.GenerateId()
	reservationID, errInsert := service.reservationData.Insert(input)
	if errInsert != nil {
		return "", errInsert
	}

	return reservationID, nil
}

func (service *reservationService) CheckAvailability(input reservation.ReservationCore) (reservation.Availability, reservation.Homestay, error) {
	chekinDate, _ := time.Parse("2006-01-02", input.CheckinDate)
	checkoutDate, _ := time.Parse("2006-01-02", input.CheckoutDate)
	reservationDuration := checkoutDate.Sub(chekinDate).Hours() / 24

	date := fmt.Sprint(checkoutDate)
	input.CheckoutDate = date
	result, err := service.reservationData.CheckAvailability(input)
	if err != nil {
		return reservation.Availability{}, reservation.Homestay{}, err
	}

	if result != 0 {
		return reservation.Availability{Status: false}, reservation.Homestay{}, err
	}

	homestay, err := service.reservationData.SelectHomestay(input.HomestayID)
	if err != nil {
		return reservation.Availability{}, reservation.Homestay{}, err
	}

	grossAmount := homestay.Price * reservationDuration

	availability := reservation.Availability{
		Status:              true,
		ReservationDuration: int(reservationDuration),
		GrossAmount:         grossAmount,
	}

	return availability, homestay, nil
}

func (service *reservationService) GetById(reservationID string) (reservation.ReservationCore, reservation.Homestay, reservation.Availability, error) {
	reservationCore, err := service.reservationData.SelectById(reservationID)
	if err != nil {
		return reservation.ReservationCore{}, reservation.Homestay{}, reservation.Availability{}, err
	}

	homestay, err := service.reservationData.SelectHomestay(reservationCore.HomestayID)
	if err != nil {
		return reservation.ReservationCore{}, reservation.Homestay{}, reservation.Availability{}, err
	}

	chekinDate, _ := time.Parse("2006-01-02", reservationCore.CheckinDate)
	checkoutDate, _ := time.Parse("2006-01-02", reservationCore.CheckoutDate)
	reservationDuration := checkoutDate.Sub(chekinDate).Hours() / 24
	grossAmount := homestay.Price * reservationDuration

	availability := reservation.Availability{
		Status:              true,
		ReservationDuration: int(reservationDuration),
		GrossAmount:         grossAmount,
	}

	return reservationCore, homestay, availability, err
}
