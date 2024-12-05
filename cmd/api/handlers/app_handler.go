package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	healthCheck := struct {
		Health    bool   `json:"health"`
		Timestamp string `json:"timestamp"`
	}{
		Health:    true,
		Timestamp: time.Now().String(),
	}

	return c.JSON(http.StatusOK, healthCheck)
}
