package handlers

import (
	"github.com/ViniciusCrisol/dynamic-db/infra/api"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/handlerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/middlewares"

	"github.com/gin-gonic/gin"
)

type domainHandler struct {
	domainController api.DomainController
}

func NewDomainHandler(domainController api.DomainController) *domainHandler {
	return &domainHandler{domainController}
}

func (handler *domainHandler) SaveDomain(context *gin.Context) {
	var requestBody handlerDTOs.SaveDomain
	parseRequestBodyErr := context.ShouldBindJSON(&requestBody)
	if parseRequestBodyErr != nil {
		middlewares.SendInternalServerErr(context)
		return
	}

	domain, saveDomainErr := handler.domainController.SaveDomain(requestBody.DomainName)
	if saveDomainErr != nil {
		middlewares.HandleErr(saveDomainErr, context)
		return
	}
	middlewares.SendJSON(200, domain, context)
}

func (handler *domainHandler) DeletedDomain(context *gin.Context) {
	domainName := context.Param("domain_name")
	deletedDomainErr := handler.domainController.DeleteDomain(domainName)
	if deletedDomainErr != nil {
		middlewares.HandleErr(deletedDomainErr, context)
		return
	}
	middlewares.SendJSON(204, nil, context)
}
