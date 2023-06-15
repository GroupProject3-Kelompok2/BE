package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"gorm.io/gorm"
)

var log = middlewares.Log()

type paymentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.PaymentData {
	return &paymentQuery{
		db: db,
	}
}

func (pq *paymentQuery) Payment(request payment.PaymentCore) (payment.PaymentCore, error) {
	paymentData := chargeMidtrans(request)

	result := pq.db.Create(&paymentData)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("payment record not found")
		return payment.PaymentCore{}, errors.New("payment record not found")
	}

	if result.RowsAffected == 0 {
		log.Warn("no charge payment has been created")
		return payment.PaymentCore{}, errors.New("row affected : 0")
	}

	if result.Error != nil {
		log.Error("error while charging payment")
		return payment.PaymentCore{}, errors.New("internal server error")
	}

	return paymentModels(paymentData), nil
}

func (pq *paymentQuery) UpdateStatus(orderID string, status string) error {
	result := pq.db.Table("payments").Where("reservation_id = ?", orderID).Update("status", status)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("payment record not found")
		return errors.New("payment record not found")
	}
	if result.RowsAffected == 0 {
		log.Warn("no payment record has been updated")
		return errors.New("no payment record has been updated")
	}
	if result.Error != nil {
		log.Error("error while updating payment status")
		return errors.New("internal server error")
	}
	return nil
}
