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

		reservationID, err := handler.reservationService.Create(reservationCore)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				log.Error("bad request")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, required field cannot be empty", nil, nil))
			} else {
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		responseData := NewReservationResponse(reservationID)

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Reservation created successfully", responseData, nil))
	}

}

func (handler *ReservationHandler) CheckAvailability() echo.HandlerFunc {
	return func(c echo.Context) error {
		reservationInput := ReservationRequest{}
		errBind := c.Bind(&reservationInput)
		if errBind != nil {
			log.Error("bad request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request", nil, nil))
		}

		reservationCore := ReservationRequestCore(reservationInput)

		availability, homestay, err := handler.reservationService.CheckAvailability(reservationCore)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				log.Error("bad request")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, required field cannot be empty", nil, nil))
			} else {
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		if !availability.Status {
			return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Not Available", nil, nil))
		}

		responseData := ReservationResponseData(reservation.ReservationCore{}, homestay, availability)

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Available", responseData, nil))
	}
}

func (handler *ReservationHandler) GetReservationById() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("reservation_id")
		reservationCore, homestay, availability, err := handler.reservationService.GetById(paramId)
		if err != nil {
			log.Error("resource not found")
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "", "Resource not found", nil, nil))
		}

		responseData := ReservationResponseData(reservationCore, homestay, availability)

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Homestays read successfully", responseData, nil))
	}
}
