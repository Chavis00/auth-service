package handlers

import (
	"authservice/helpers"
	"authservice/pkg"
	"authservice/schemas"
	services "authservice/services/auth"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handlerRegister struct {
	service services.ServiceRegister
}

func NewHandlerRegister(service services.ServiceRegister) *handlerRegister {
	return &handlerRegister{service: service}
}

func (h *handlerRegister) RegisterHandler(ctx *gin.Context) {

	var input schemas.SchemaAuth
	ctx.ShouldBindJSON(&input)

	errResponse := make(map[string]string)

	// Validación de campo "Fullname"
	if input.Fullname == "" {
		errResponse["Fullname"] = "fullname is required on body"
	} else if strings.ToLower(input.Fullname) != input.Fullname {
		errResponse["Fullname"] = "fullname must be using lowercase"
	}

	// Validación de campo "Email"
	if input.Email == "" {
		errResponse["Email"] = "email is required on body"
	} else {
		// Validación del formato de correo electrónico
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !re.MatchString(input.Email) {
			errResponse["Email"] = "email format is not valid"
		}
	}

	// Validación de campo "Password"
	if input.Password == "" {
		errResponse["Password"] = "password is required on body"
	} else if len(input.Password) < 8 {
		errResponse["Password"] = "password minimum must be 8 character"
	}

	if len(errResponse) > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	res, err := h.service.RegisterService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Email already exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Register new account failed", err.Code, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}

		accessToken, errToken := pkg.Sign(accessTokenData, os.Getenv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		_, errSendMail := pkg.SendGridMail(res.Fullname, res.Email, "Activation Account", "template_register", accessToken)

		if errSendMail != nil {
			defer logrus.Error(errSendMail.Error())
			helpers.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		helpers.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
