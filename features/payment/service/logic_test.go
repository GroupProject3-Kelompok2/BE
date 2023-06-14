package service_test

import (
	"errors"
	"testing"

	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/features/payment/service"
	"github.com/GroupProject3-Kelompok2/BE/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPayment(t *testing.T) {
	data := mocks.NewPaymentData(t)
	validate := validator.New()
	service := service.New(data, validate)

	t.Run("success charging payment", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "string",
			Amount:        "1200.00",
			BankAccount:   "bca",
			VANumber:      "string",
			Status:        "pending",
		}
		result := payment.PaymentCore{
			PaymentID:     "string",
			ReservationID: "string",
			Amount:        "1200.00",
			BankAccount:   "bca",
			VANumber:      "string",
			Status:        "pending",
		}

		data.On("Payment", mock.Anything).Return(result, nil).Once()
		res, err := service.Payment(request)
		assert.Nil(t, err)
		assert.Equal(t, result.PaymentID, res.PaymentID)
		assert.NotEmpty(t, result.VANumber)
		data.AssertExpectations(t)
	})

	t.Run("empty reservation ID", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "",
			Amount:        "1200.00",
			BankAccount:   "bca",
			VANumber:      "string",
			Status:        "pending",
		}

		res, err := service.Payment(request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "reservation_id cannot be empty")
		assert.Equal(t, payment.PaymentCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("empty bank account", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "string",
			Amount:        "1200.00",
			BankAccount:   "",
			VANumber:      "string",
			Status:        "pending",
		}

		res, err := service.Payment(request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "bank account cannot be empty")
		assert.Equal(t, payment.PaymentCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("empty amount", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "string",
			Amount:        "",
			BankAccount:   "bca",
			VANumber:      "string",
			Status:        "pending",
		}

		res, err := service.Payment(request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "amount cannot be empty")
		assert.Equal(t, payment.PaymentCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("unsupported bank account", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "string",
			Amount:        "1200.00",
			BankAccount:   "unsupported",
			VANumber:      "string",
			Status:        "pending",
		}

		res, err := service.Payment(request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "unsupported bank account")
		assert.Equal(t, payment.PaymentCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("payment record not found", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "string",
			Amount:        "1200.00",
			BankAccount:   "bca",
			VANumber:      "string",
			Status:        "pending",
		}

		data.On("Payment", mock.Anything).Return(payment.PaymentCore{}, errors.New("record not found")).Once()

		res, err := service.Payment(request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "payment record not found")
		assert.Equal(t, payment.PaymentCore{}, res)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		request := payment.PaymentCore{
			ReservationID: "string",
			Amount:        "1200.00",
			BankAccount:   "bca",
			VANumber:      "string",
			Status:        "pending",
		}

		data.On("Payment", mock.Anything).Return(payment.PaymentCore{}, errors.New("internal error")).Once()

		res, err := service.Payment(request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "internal server error")
		assert.Equal(t, payment.PaymentCore{}, res)
		data.AssertExpectations(t)
	})
}
