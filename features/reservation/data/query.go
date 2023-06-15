package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationDataInterface {
	return &reservationQuery{
		db: db,
	}
}

func (repo *reservationQuery) Insert(input reservation.ReservationCore) (string, error) {
	reservationInputGorm := ReservationModel(input)
	tx := repo.db.Create(&reservationInputGorm)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", errors.New("insert failed, row affected = 0")
	}

	reservationCore := ReservationCore(reservationInputGorm)

	return reservationCore.ReservationID, nil
}

func (repo *reservationQuery) CheckAvailability(input reservation.ReservationCore) (int64, error) {
	reservationInputGorm := ReservationModel(input)
	tx := repo.db.Raw("SELECT reservation_id from reservations WHERE homestay_id = ? AND check_in_date BETWEEN ? AND ? OR check_out_date BETWEEN ? AND ?",
		reservationInputGorm.HomestayID, reservationInputGorm.CheckInDate, reservationInputGorm.CheckOutDate, reservationInputGorm.CheckInDate, reservationInputGorm.CheckOutDate).
		Scan(&reservationInputGorm)

	return tx.RowsAffected, tx.Error
}

func (repo *reservationQuery) SelectById(reservationId string) (reservation.ReservationCore, error) {
	var reservationGorm Reservation
	tx := repo.db.Where("reservation_id = ?", reservationId).First(&reservationGorm)
	if tx.Error != nil {
		return reservation.ReservationCore{}, errors.New("error reservation not found")
	}

	reservationCore := ReservationCore(reservationGorm)
	return reservationCore, nil
}

func (repo *reservationQuery) SelectHomestay(homestayID string) (reservation.Homestay, error) {
	var homestayData reservation.Homestay

	tx := repo.db.Raw("SELECT name, price FROM homestays WHERE homestay_id = ?", homestayID).Scan(&homestayData)
	if tx.Error != nil {
		return reservation.Homestay{}, tx.Error
	}

	return homestayData, nil
}
