package handlers

import (
	"github.com/ViniciusCrisol/dynamic-db/infra/api"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/handlerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/middlewares"
	"github.com/gin-gonic/gin"
)

type clusterHandler struct {
	clusterController api.ClusterController
}

func NewClusterHandler(clusterController api.ClusterController) *clusterHandler {
	return &clusterHandler{clusterController}
}

func (handler *clusterHandler) SaveCluster(context *gin.Context) {
	var requestBody handlerDTOs.SaveCluster
	parseRequestBodyErr := context.ShouldBindJSON(&requestBody)
	if parseRequestBodyErr != nil {
		middlewares.SendInternalServerErr(context)
		return
	}
	domainName := context.Param("domain_name")

	cluster, saveClusterErr := handler.clusterController.SaveCluster(controllerDTOs.SaveCluster{
		DomainName:  domainName,
		ClusterName: requestBody.ClusterName,
	})
	if saveClusterErr != nil {
		middlewares.HandleErr(saveClusterErr, context)
		return
	}
	middlewares.SendJSON(200, cluster, context)
}

func (handler *clusterHandler) DeleteCluster(context *gin.Context) {
	domainName := context.Param("domain_name")
	clusterName := context.Param("cluster_name")
	deleteClusterErr := handler.clusterController.DeleteCluster(controllerDTOs.DeleteCluster{
		DomainName:  domainName,
		ClusterName: clusterName,
	})
	if deleteClusterErr != nil {
		middlewares.HandleErr(deleteClusterErr, context)
		return
	}
	middlewares.SendJSON(204, nil, context)
}
