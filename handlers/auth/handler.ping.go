package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"authservice/helpers"
	services "authservice/services/auth"
)

type handlerPing struct {
	service services.ServicePing
}

func NewHandlerPing(service services.ServicePing) *handlerPing {
	return &handlerPing{service: service}
}

func (h *handlerPing) PingHandler(ctx *gin.Context) {
	res := h.service.PingService()
	helpers.APIResponse(ctx, res, http.StatusOK, http.MethodGet, nil)
}
