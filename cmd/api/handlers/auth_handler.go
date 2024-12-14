package handlers

import (
	"bouguette/cmd/api/requests"
	"bouguette/common"

	"github.com/labstack/echo/v4"
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

	if validationErrors != nil && len(validationErrors) > 0 {
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	return common.SendSuccessResponse(c, "User registered successfully", nil)
}
