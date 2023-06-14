package handler

import "github.com/GroupProject3-Kelompok2/BE/features/payment"

type CreatePaymentRequest struct {
	ReservationID string `json:"reservation_id"`
	BankAccount   string `json:"bank_account"`
	Amount        string `json:"amount"`
}

func RequestToCore(data interface{}) payment.PaymentCore {
	res := payment.PaymentCore{}
	switch v := data.(type) {
	case CreatePaymentRequest:
		res.ReservationID = v.ReservationID
		res.BankAccount = v.BankAccount
		res.Amount = v.Amount
	default:
		return payment.PaymentCore{}
	}
	return res
}
