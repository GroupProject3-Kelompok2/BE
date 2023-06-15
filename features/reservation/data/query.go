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
	query := "SELECT reservation_id from reservations " +
		"WHERE homestay_id = ? AND check_in_date BETWEEN ? AND ? OR check_out_date BETWEEN ? AND ?"
	tx := repo.db.Raw(query, reservationInputGorm.HomestayID,
		reservationInputGorm.CheckInDate, reservationInputGorm.CheckOutDate,
		reservationInputGorm.CheckInDate, reservationInputGorm.CheckOutDate).
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

func (repo *reservationQuery) SelectAllByUserId(userID string) ([]reservation.ReservationCore, error) {
	var reservationsData []ReservationData
	query := ("select r.reservation_id, h.name as homestay_name, r.check_in_date, r.check_out_date, " +
		"h.price as homestay_price, datediff(check_out_date, check_in_date) as duration, " +
		"p.amount, p.bank_account, p.va_number, p.status " +
		"from reservations as r inner join homestays as h on r.homestay_id = h.homestay_id " +
		"inner join payments as p on r.reservation_id = p.reservation_id " +
		"where r.user_id = ? order by r.created_at desc")
	tx := repo.db.Raw(query, userID).
		Scan(&reservationsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error reservations not found")
	}

	var reservationCoreAll []reservation.ReservationCore
	for _, value := range reservationsData {
		reservationCore := ReservationDataCore(value)
		reservationCoreAll = append(reservationCoreAll, reservationCore)
	}

	return reservationCoreAll, nil
}
