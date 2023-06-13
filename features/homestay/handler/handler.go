package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

var log = middlewares.Log()

type HomestayHandler struct {
	homestayService homestay.HomestayServiceInterface
}

func New(service homestay.HomestayServiceInterface) *HomestayHandler {
	return &HomestayHandler{
		homestayService: service,
	}
}

func (handler *HomestayHandler) CreateHomestay(c echo.Context) error {
	homestayInput := HomestayRequest{}
	errBind := c.Bind(&homestayInput)
	if errBind != nil {
		log.Error("handler - error on bind request")
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request", nil, nil))
	}

	homestayCore := HomestayRequestCore(homestayInput)
	err := handler.homestayService.Create(homestayCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, required field cannot be empty", nil, nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}
	}

	return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Successfully created an account.", nil, nil))
}
