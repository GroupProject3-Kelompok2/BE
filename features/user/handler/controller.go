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
			log.Error("controller - error on bind request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request"+errBind.Error(), nil, nil))
		}

		_, err := uh.service.Register(RequestToCore(request))
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "empty"):
				log.Error("bad request, request cannot be empty")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, request cannot be empty", nil, nil))
			case strings.Contains(err.Error(), "duplicated"):
				log.Error("bad request, duplicate data request")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, duplicate data request", nil, nil))
			case strings.Contains(err.Error(), "email"):
				log.Error("bad request, invalid email format")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, invalid email format", nil, nil))
			case strings.Contains(err.Error(), "low password"):
				log.Error("bad request, password does not match")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, password does not match", nil, nil))
			case strings.Contains(err.Error(), "password"):
				log.Error("internal server error, hashing password")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			default:
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Successfully created an account.", nil, nil))
	}
}

// Login implements user.UserHandler
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := LoginRequest{}
		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind login input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request"+errBind.Error(), nil, nil))
		}

		resp, token, err := uh.service.Login(RequestToCore(request))
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "invalid email format"):
				log.Error("bad request, invalid email format")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, invalid email format", nil, nil))
			case strings.Contains(err.Error(), "password cannot be empty"):
				log.Error("bad request, password cannot be empty")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, password cannot be empty", nil, nil))
			case strings.Contains(err.Error(), "invalid email and password"):
				log.Error("bad request, invalid email and password")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, invalid email and password", nil, nil))
			case strings.Contains(err.Error(), "password does not match"):
				log.Error("bad request, password does not match")
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request, password does not match", nil, nil))
			case strings.Contains(err.Error(), "error while creating jwt token"):
				log.Error("internal server error, error while creating jwt token")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			default:
				log.Error("internal server error")
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Successful login", loginResponse{
			UserID: resp.UserID, Email: resp.Email, Token: token,
		}, nil))
	}
}

// ProfileUser implements user.UserHandler
func (uh *userHandler) ProfileUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _, err := middlewares.ExtractToken(c)
		if err != nil {
			log.Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "", "Missing or Malformed JWT", nil, nil))
		}

		user, err := uh.service.ProfileUser(userId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "", "The requested resource was not found", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		resp := profileUser(user)
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Successfully operation.", resp, nil))
	}
}
