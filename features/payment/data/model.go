package data

import (
	"strconv"
	"strings"
	"time"

	"github.com/GroupProject3-Kelompok2/BE/app/config"
	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type Payment struct {
	PaymentID     string `gorm:"primaryKey;type:varchar(255)"`
	ReservationID string `gorm:"type:varchar(21)"`
	Amount        string
	BankAccount   string         `gorm:"type:enum('bca', 'bri', 'bni', 'mandiri'); default:'bca'"`
	VANumber      string         `gorm:"type:varchar(21)"`
	Status        string         `gorm:"type:varchar(21)"`
	CreatedAt     time.Time      `gorm:"type:datetime"`
	UpdatedAt     time.Time      `gorm:"type:datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func paymentModels(p Payment) payment.PaymentCore {
	return payment.PaymentCore{
		PaymentID:     p.PaymentID,
		ReservationID: p.ReservationID,
		Amount:        p.Amount,
		BankAccount:   p.BankAccount,
		VANumber:      p.VANumber,
		Status:        p.Status,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func paymentEntities(p payment.PaymentCore) Payment {
	return Payment{
		PaymentID:     p.PaymentID,
		ReservationID: p.ReservationID,
		Amount:        p.Amount,
		BankAccount:   p.BankAccount,
		VANumber:      p.VANumber,
		Status:        p.Status,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func chargeMidtrans(request payment.PaymentCore) Payment {
	var bankMap = map[string]coreapi.BankTransferDetails{
		"bni": {
			Bank: midtrans.BankBni,
		},
		"bca": {
			Bank: midtrans.BankBca,
		},
		"bri": {
			Bank: midtrans.BankBri,
		},
	}

	var c = coreapi.Client{}
	c.New(config.MIDTRANS_SERVERKEY, midtrans.Sandbox)

	amount, err := strconv.ParseInt(request.Amount, 10, 64)
	if err != nil {
		return Payment{}
	}

	log.Info(request.BankAccount)
	bankTransfer, ok := bankMap[request.BankAccount]
	if !ok {
		return Payment{}
	}

	req := &coreapi.ChargeReq{
		PaymentType:  coreapi.PaymentTypeBankTransfer,
		BankTransfer: &bankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.ReservationID,
			GrossAmt: amount,
		},
	}

	resp, _ := c.ChargeTransaction(req)
	banks := make([]string, len(resp.VaNumbers))
	for i, bank := range resp.VaNumbers {
		banks[i] = bank.Bank
	}
	banksStr := strings.Join(banks, ",")

	vaNumbers := make([]string, len(resp.VaNumbers))
	for i, vaNumber := range resp.VaNumbers {
		vaNumbers[i] = vaNumber.VANumber
	}
	vaNumbersStr := strings.Join(vaNumbers, ",")
	createdAt, _ := time.Parse(time.RFC3339, resp.TransactionTime)
	return Payment{
		PaymentID:     resp.TransactionID,
		ReservationID: resp.OrderID,
		Amount:        resp.GrossAmount,
		BankAccount:   banksStr,
		VANumber:      vaNumbersStr,
		Status:        resp.TransactionStatus,
		CreatedAt:     createdAt,
	}
}
