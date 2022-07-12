package middlewares

import (
	"github.com/ViniciusCrisol/dynamic-db/infra/api"
	"github.com/ViniciusCrisol/dynamic-db/utils"
	"github.com/gin-gonic/gin"
)

func SendJSON(status int, data interface{}, context *gin.Context) {
	response := api.Response{
		Data:   data,
		Status: status,
	}
	context.JSON(status, response)
}

func HandleErr(err error, context *gin.Context) {
	message, status := utils.GetMessageAndHTTPStatusFromErr(err)
	response := api.Response{
		Status:  status,
		Message: message,
	}
	context.JSON(status, response)
}

func SendRouteNotFound(context *gin.Context) {
	status := 404
	message := utils.ErrMessages["route-not-found"]
	response := api.Response{
		Status:  status,
		Message: message,
	}
	context.JSON(status, response)
}

func SendInternalServerErr(context *gin.Context) {
	status := 500
	message := utils.ErrMessages["internal-server-err"]
	response := api.Response{
		Status:  status,
		Message: message,
	}
	context.JSON(status, response)
}
