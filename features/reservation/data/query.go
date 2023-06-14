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

func (repo *reservationQuery) Insert(input reservation.ReservationCore) error {
	reservationInputGorm := ReservationModel(input)
	tx := repo.db.Create(&reservationInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
