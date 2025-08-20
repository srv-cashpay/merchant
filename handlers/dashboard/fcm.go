package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TokenRequest struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

func (h *domainHandler) SaveToken(c echo.Context) error {
	var req TokenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.serviceDashboard.SaveToken(req.UserID, req.Token); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *domainHandler) SendBroadcast(c echo.Context) error {
	var req struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Enqueue broadcast ke worker pool
	h.serviceDashboard.EnqueueFCM(req.Title, req.Body)

	return c.JSON(http.StatusOK, map[string]string{"status": "broadcast enqueued"})
}
