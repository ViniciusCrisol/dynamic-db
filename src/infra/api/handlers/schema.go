package handlers

import (
	"github.com/ViniciusCrisol/dynamic-db/infra/api"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/handlerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/middlewares"
	"github.com/gin-gonic/gin"
)

type schemaHandler struct {
	schemaController api.SchemaController
}

func NewSchemaHandler(schemaController api.SchemaController) *schemaHandler {
	return &schemaHandler{schemaController}
}

func (handler *schemaHandler) SaveSchema(context *gin.Context) {
	var requestBody handlerDTOs.SaveSchema
	parseRequestBodyErr := context.ShouldBindJSON(&requestBody)
	if parseRequestBodyErr != nil {
		middlewares.SendInternalServerErr(context)
		return
	}
	domainName := context.Param("domain_name")
	clusterName := context.Param("cluster_name")

	schema, saveSchemaErr := handler.schemaController.SaveSchema(controllerDTOs.SaveSchema{
		DomainName:    domainName,
		ClusterName:   clusterName,
		SchemaContent: requestBody.SchemaContent,
	})
	if saveSchemaErr != nil {
		middlewares.HandleErr(saveSchemaErr, context)
		return
	}
	middlewares.SendJSON(200, schema, context)
}
