package service

import (
	"errors"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/go-playground/validator/v10"
)

var log = middlewares.Log()

type paymentService struct {
	query    payment.PaymentData
	validate *validator.Validate
}

func New(ud payment.PaymentData, v *validator.Validate) payment.PaymentService {
	return &paymentService{
		query:    ud,
		validate: v,
	}
}

// Payment implements payment.PaymentService
func (ps *paymentService) Payment(request payment.PaymentCore) (payment.PaymentCore, error) {
	err := ps.validate.Struct(request)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "ReservationID"):
			log.Warn("reservation_id cannot be empty")
			return payment.PaymentCore{}, errors.New("reservation_id cannot be empty")
		case strings.Contains(err.Error(), "BankAccount"):
			log.Warn("bank account cannot be empty")
			return payment.PaymentCore{}, errors.New("bank account cannot be empty")
		case strings.Contains(err.Error(), "Amount"):
			log.Warn("amount cannot be empty")
			return payment.PaymentCore{}, errors.New("amount cannot be empty")
		}
	}

	if request.BankAccount != "bca" && request.BankAccount != "bri" && request.BankAccount != "bni" {
		log.Error("only bca bni, and bni are avalaible atm")
		return payment.PaymentCore{}, errors.New("unsupported bank account")
	}

	result, err := ps.query.Payment(request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("payment record not found")
			return payment.PaymentCore{}, errors.New("payment record not found")
		} else {
			log.Error("internal server error")
			return payment.PaymentCore{}, errors.New("internal server error")
		}
	}

	log.Sugar().Infof("new user has been created: %s", result.PaymentID)
	return result, nil
}

// func (ps *paymentService) UpdateStatus(status, orderID string) error {
// 	if err := ps.query.UpdateStatus(orderID, status); err != nil {
// 		return err
// 	}
// 	return nil
// }
