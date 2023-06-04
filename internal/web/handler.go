package web

import (
	"github.com/freekoder/go-project-boilerplate/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

type SuccessResponse struct {
	Status string `json:"status"`
}

func (h *Handler) Index(e echo.Context) error {
	return e.JSON(http.StatusOK, SuccessResponse{Status: "ok"})
}
