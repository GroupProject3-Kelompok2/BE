package handler

import (
	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
)

type paymentResponse struct {
	PaymentID     string           `json:"payment_id"`
	ReservationID string           `json:"reservation_id"`
	Amount        string           `json:"amount"`
	BankAccount   string           `json:"bank_account"`
	VANumber      string           `json:"va_number"`
	Status        string           `json:"status"`
	CreatedAt     helper.LocalTime `json:"created_at"`
	UpdatedAt     helper.LocalTime `json:"updated_at"`
}

func paymentResp(p payment.PaymentCore) paymentResponse {
	return paymentResponse{
		PaymentID:     p.PaymentID,
		ReservationID: p.ReservationID,
		Amount:        p.Amount,
		BankAccount:   p.BankAccount,
		VANumber:      p.VANumber,
		Status:        p.Status,
		CreatedAt:     helper.LocalTime(p.CreatedAt),
		UpdatedAt:     helper.LocalTime(p.UpdatedAt),
	}
}

// type notificationResponse struct {
// 	StatusCode        string `json:"status_code"`
// 	OrderID           string `json:"order_id"`
// 	TransactionTime   string `json:"transaction_time"`
// 	TransactionStatus string `json:"transaction_status"`
// 	FraudStatus       string `json:"fraud_status"`
// 	PaymentType       string `json:"payment_type"`
// 	GrossAmount       int    `json:"gross_amount"`
// }
