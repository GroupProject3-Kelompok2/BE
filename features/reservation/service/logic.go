package service

import (
	"fmt"

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

func (service *reservationService) CheckAvailability(input reservation.ReservationCore) (reservation.ReservationCore, error) {
	result, err := service.reservationData.CheckAvailability(input)
	if err != nil {
		return reservation.ReservationCore{}, err
	}

	if result != 0 {
		return reservation.ReservationCore{Availability: reservation.Availability{Status: false}}, err
	}

	homestay, err := service.reservationData.SelectHomestay(input.HomestayID)
	if err != nil {
		return reservation.ReservationCore{}, err
	}

	reservationDuration := input.CheckoutDate.Sub(input.CheckinDate).Hours() / 24
	grossAmount := homestay.Price * reservationDuration

	availability := reservation.Availability{
		Status:              true,
		ReservationDuration: int(reservationDuration),
		GrossAmount:         grossAmount,
	}

	reservationCore := reservation.ReservationCore{
		CheckinDate:  input.CheckinDate,
		CheckoutDate: input.CheckoutDate,
		Homestay:     homestay,
		Availability: availability,
	}

	return reservationCore, nil
}

func (service *reservationService) GetById(reservationID string) (reservation.ReservationCore, error) {
	reservationCore, err := service.reservationData.SelectById(reservationID)
	if err != nil {
		return reservation.ReservationCore{}, err
	}

	homestay, err := service.reservationData.SelectHomestay(reservationCore.HomestayID)
	if err != nil {
		return reservation.ReservationCore{}, err
	}

	reservationDuration := reservationCore.CheckoutDate.Sub(reservationCore.CheckinDate).Hours() / 24
	grossAmount := homestay.Price * reservationDuration

	reservationCore.Homestay = homestay
	reservationCore.Availability = reservation.Availability{
		Status:              true,
		ReservationDuration: int(reservationDuration),
		GrossAmount:         grossAmount,
	}

	fmt.Println(reservationCore.Availability)

	return reservationCore, err
}

func (service *reservationService) GetAllByUserId(userID string) ([]reservation.ReservationCore, error) {
	data, err := service.reservationData.SelectAllByUserId(userID)
	if err != nil {
		return nil, err
	}
	return data, err
}
