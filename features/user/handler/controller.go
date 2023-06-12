package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

var log = middlewares.Log()

type userHandler struct {
	service user.UserService
}

func New(us user.UserService) user.UserHandler {
	return &userHandler{
		service: us,
	}
}

// Register implements user.UserHandler
func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterRequest{}
		errBind := c.Bind(&request)
		if errBind != nil {
			log.Error("error on bind request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed", "Bad request", nil, nil))
		}

		result, err := uh.service.Register(RequestToCore(request))
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "empty"):
				log.Error("bad request, username, email, password cannot be empty")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed", "Bad request", nil, nil))
			case strings.Contains(err.Error(), "duplicated"):
				log.Error("bad request, duplicate data request")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed", "Bad request", nil, nil))
			case strings.Contains(err.Error(), "email"):
				log.Error("bad request, invalid email format")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed", "Bad request", nil, nil))
			case strings.Contains(err.Error(), "low password"):
				log.Error("bad request, low password strength")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed", "Bad request", nil, nil))
			case strings.Contains(err.Error(), "password"):
				log.Error("internal server error, hashing password")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Failed", "Internal server error", nil, nil))
			default:
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Failed", "Internal server error", nil, nil))
			}
		}

		resp := register(result)
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Failed", "Successfully created an account.", resp, nil))
	}
}
