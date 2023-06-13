package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	storages "github.com/GroupProject3-Kelompok2/BE/utils/storage"
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
			log.Error("forbidden")
			return c.JSON(http.StatusForbidden, helper.ResponseFormat(http.StatusForbidden, "", "Forbidden", nil, nil))
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

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Homestay created successfully", nil, nil))
	}

}

func (handler *HomestayHandler) UpdateHomestayById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, userRole, errExtract := middlewares.ExtractToken(c)
		if errExtract != nil {
			log.Error("failed to extract token")
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		if userRole != "hoster" {
			log.Error("forbidden")
			return c.JSON(http.StatusForbidden, helper.ResponseFormat(http.StatusForbidden, "", "Forbidden", nil, nil))
		}

		paramId := c.Param("homestay_id")

		homestayInput := HomestayRequest{}
		errBind := c.Bind(&homestayInput)
		if errBind != nil {
			log.Error("bad request")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad request", nil, nil))
		}

		homestayCore := HomestayRequestCore(homestayInput)
		err := handler.homestayService.UpdateById(userId, paramId, homestayCore)
		if err != nil {
			log.Error("resource not found")
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "", "Resource not found", nil, nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Homestay updated successfully", nil, nil))
	}
}

func (handler *HomestayHandler) DeleteHomestayById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, userRole, errExtract := middlewares.ExtractToken(c)
		if errExtract != nil {
			log.Error("failed to extract token")
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
		}

		if userRole != "hoster" {
			log.Error("forbidden")
			return c.JSON(http.StatusForbidden, helper.ResponseFormat(http.StatusForbidden, "", "Forbidden", nil, nil))
		}

		paramId := c.Param("homestay_id")
		err := handler.homestayService.DeleteById(userId, paramId)
		if err != nil {
			log.Error("resource not found")
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "", "Resource not found", nil, nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Homestay deleted successfully", nil, nil))
	}
}

func (handler *HomestayHandler) GetAllHomestay() echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := handler.homestayService.GetAll()
		if err != nil {
			log.Error("resource not found")
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "", "Resource not found", nil, nil))
		}

		var homestaysResponse []HomestayResponse
		for _, value := range results {
			homestaysResponse = append(homestaysResponse, HomestayCoreResponse(value))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Homestays read successfully", homestaysResponse, nil))
	}
}

func (handler *HomestayHandler) GetHomestayById() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("homestay_id")
		results, err := handler.homestayService.GetById(paramId)
		if err != nil {
			log.Error("resource not found")
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "", "Resource not found", nil, nil))
		}

		homestayResponse := HomestayCoreResponse(results)

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "", "Homestays read successfully", homestayResponse, nil))
	}
}

func (handler *HomestayHandler) HomestayPictures() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := HomestayPicturesRequest{}
		_, role, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "", "Missing or Malformed JWT.", nil, nil))
		}

		if role != "hoster" {
			c.Logger().Error("unauthorized access")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "", "Unauthorized access", nil, nil))
		}

		var imageURL string
		file, err1 := c.FormFile("profile_picture")
		if err1 == nil {
			imageURL, err1 = storages.UploadImage(c, file)
			if err1 != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Failed to upload image", nil, nil))
			}
			request.HomestayPicture = &imageURL
		}

		homestayId := c.Param("id")
		err := handler.homestayService.HomestayPictures(homestayId, HomestayPictureRequestToCore(&request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "", "Bad Request, data cannot be empty while updating", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "", "Internal server error", nil, nil))
			}
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "", "Successfully upload a homestay_picture.", nil, nil))
	}
}
