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

type handlerLogin struct {
	service services.ServiceLogin
}

func NewHandlerLogin(service services.ServiceLogin) *handlerLogin {
	return &handlerLogin{service: service}
}

func (h *handlerLogin) LoginHandler(ctx *gin.Context) {

	var input schemas.SchemaAuth
	ctx.ShouldBindJSON(&input)
	res, err := h.service.LoginService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "User account is not registered", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "User account is not active", err.Code, http.MethodPost, nil)
		return
	case "error_03":
		helpers.APIResponse(ctx, "Username or password is wrong", err.Code, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, os.Getenv("JWT_SECRET"), 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		helpers.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}
