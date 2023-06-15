package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/payment"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

var log = middlewares.Log()

type paymentHandler struct {
	service payment.PaymentService
}

func New(us payment.PaymentService) payment.PaymentHandler {
	return &paymentHandler{
		service: us,
	}
}

func (tc *paymentHandler) Payment() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := createPaymentRequest{}
		_, _, err := middlewares.ExtractToken(c)
		if err != nil {
			log.Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "", "Missing or Malformed JWT", nil, nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			log.Error("error on bind request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request"+errBind.Error(), nil, nil))
		}

		payment, err := tc.service.Payment(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "unsupported bank account") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, unsupported bank account", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		log.Sugar().Infoln(payment)
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Successful Operation", paymentResp(payment), nil))
	}
}

func (tc *paymentHandler) Notification() echo.HandlerFunc {
	return func(c echo.Context) error {
		midtransResponse := midtransCallback{}
		errBind := c.Bind(&midtransResponse)
		if errBind != nil {
			log.Sugar().Errorf("error on binding notification input", errBind)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request: "+errBind.Error(), nil, nil))
		}

		log.Sugar().Infof("callback midtrans: %s, bank: %s, reservation ID: %s, transaction ID: %s",
			midtransResponse.TransactionStatus, midtransResponse.Bank,
			midtransResponse.OrderID, midtransResponse.TransactionID)

		return nil
	}
}
