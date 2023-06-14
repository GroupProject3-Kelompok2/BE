package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/reservation"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

var log = middlewares.Log()

type ReservationHandler struct {
	reservationService reservation.ReservationServiceInterface
}

func New(service reservation.ReservationServiceInterface) *ReservationHandler {
	return &ReservationHandler{
		reservationService: service,
	}
}

func (handler *ReservationHandler) CreateReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _, errExtract := middlewares.ExtractToken(c)
		if errExtract != nil {
			log.Error("failed to extract token")
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		reservationInput := ReservationRequest{}
		errBind := c.Bind(&reservationInput)
		if errBind != nil {
			log.Error("bad request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request", nil, nil))
		}

		reservationCore := ReservationRequestCore(reservationInput)
		reservationCore.UserID = userId
		reservationCore.ReservationID, _ = helper.GenerateId()

		err := handler.reservationService.Create(reservationCore)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				log.Error("bad request")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, required field cannot be empty", nil, nil))
			} else {
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Reservation created successfully", nil, nil))
	}

}
