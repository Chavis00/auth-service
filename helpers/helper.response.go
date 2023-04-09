package helpers

import (
	schemas "authservice/schemas"

	"github.com/gin-gonic/gin"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := schemas.SchemaErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	}

	ctx.AbortWithStatusJSON(StatusCode, errResponse)
}
