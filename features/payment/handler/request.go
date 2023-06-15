package handler

import (
	"fmt"

	"github.com/GroupProject3-Kelompok2/BE/features/payment"
)

type createPaymentRequest struct {
	ReservationID string `json:"reservation_id"`
	BankAccount   string `json:"bank_account"`
	Amount        string `json:"amount"`
}

type midtransCallback struct {
	TransactionTime     string  `json:"transaction_time"`
	TransactionStatus   string  `json:"transaction_status"`
	TransactionID       string  `json:"transaction_id"`
	StatusMessage       string  `json:"status_message"`
	StatusCode          string  `json:"status_code"`
	SignatureKey        string  `json:"signature_key"`
	PaymentType         string  `json:"payment_type"`
	OrderID             string  `json:"order_id"`
	MerchantID          string  `json:"merchant_id"`
	MaskedCard          string  `json:"masked_card"`
	GrossAmount         float64 `json:"gross_amount"`
	FraudStatus         string  `json:"fraud_status"`
	ECI                 string  `json:"eci"`
	Currency            string  `json:"currency"`
	ChannelResponseMsg  string  `json:"channel_response_message"`
	ChannelResponseCode string  `json:"channel_response_code"`
	CardType            string  `json:"card_type"`
	Bank                string  `json:"bank"`
	ApprovalCode        string  `json:"approval_code"`
}

func RequestToCore(data interface{}) payment.PaymentCore {
	res := payment.PaymentCore{}
	switch v := data.(type) {
	case createPaymentRequest:
		res.ReservationID = v.ReservationID
		res.BankAccount = v.BankAccount
		res.Amount = v.Amount
	case midtransCallback:
		res.ReservationID = v.OrderID
		// res.TransactionTime = v.TransactionTime
		res.Status = v.TransactionStatus
		// res.TransactionID = v.TransactionID
		// res.PaymentType = v.PaymentType
		res.Amount = fmt.Sprintf("%.2f", v.GrossAmount)
		res.BankAccount = v.Bank
	default:
		return payment.PaymentCore{}
	}
	return res
}
