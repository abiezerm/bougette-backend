package handlers

import (
	"bouguette/cmd/api/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	// bind request
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "Bad request")
	}

	// validate request
	validationErrors := h.ValidateBodyRequest(c, *payload)
	if len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	return c.JSON(http.StatusBadRequest, validationErrors)
}
