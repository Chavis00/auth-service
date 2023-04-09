package handlers

import (
	"authservice/helpers"
	pkg "authservice/pkg"
	"authservice/schemas"
	services "authservice/services/auth"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handlerActivation struct {
	service services.ServiceActivation
}

func NewHandlerActivation(service services.ServiceActivation) *handlerActivation {
	return &handlerActivation{service: service}
}

func (h *handlerActivation) ActivationHandler(ctx *gin.Context) {

	var input schemas.SchemaAuth

	token := ctx.Param("token")
	resultToken, errToken := pkg.VerifyToken(token, os.Getenv("JWT_SECRET"))

	if errToken != nil {
		defer logrus.Error(errToken.Error())
		helpers.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	result := pkg.DecodeToken(resultToken)
	input.Email = result.Claims.Email
	input.Active = true

	_, err := h.service.ActivationService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "User account is not exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "User account hash been active please login", err.Code, http.MethodPost, nil)
		return
	case "error_03":
		helpers.APIResponse(ctx, "Activation account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Activation account success", http.StatusOK, http.MethodPost, nil)
	}
}
