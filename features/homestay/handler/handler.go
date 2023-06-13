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

func (handler *HomestayHandler) CreateHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, userRole, errExtract := middlewares.ExtractToken(c)
		if errExtract != nil {
			log.Error("failed to extract token")
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		if userRole != "hoster" {
			log.Error("unathorized")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "", "Unauthorize", nil, nil))
		}

		homestayInput := HomestayRequest{}
		errBind := c.Bind(&homestayInput)
		if errBind != nil {
			log.Error("bad request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request", nil, nil))
		}

		homestayCore := HomestayRequestCore(homestayInput)
		homestayCore.UserID = userId
		homestayCore.HomestayID, _ = helper.GenerateId()

		err := handler.homestayService.Create(homestayCore)
		if err != nil {
			if strings.Contains(err.Error(), "validation") {
				log.Error("bad request")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, required field cannot be empty", nil, nil))
			} else {
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Successfully created an account.", nil, nil))
	}

}
