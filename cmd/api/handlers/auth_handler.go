package handlers

import (
	"bouguette/cmd/api/requests"
	"bouguette/cmd/api/services"
	"bouguette/common"
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	// bind request
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		c.Logger().Error(err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	// validate request
	validationErrors := h.ValidateBodyRequest(c, *payload)

	if len(validationErrors) > 0 {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	//check if email exists
	userService := services.NewUserService(h.DB)
	_, error := userService.GetUserByEmail(payload.Email)
	if !errors.Is(error, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Email has already been taken")
	}
	//print(userExist)

	// create use in the database
	registeredUser, err := userService.RegisterUser(payload)
	if err != nil {
		return err
	}

	//send response
	return common.SendSuccessResponse(c, "User registered successfully", registeredUser)
}
