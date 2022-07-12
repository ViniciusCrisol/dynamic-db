package controllers

import (
	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
)

type clusterController struct {
	saveClusterUsecase   core.SaveClusterUsecase
	deleteClusterUsecase core.DeleteClusterUsecase
}

func NewClusterController(
	saveClusterUsecase core.SaveClusterUsecase,
	deleteClusterUsecase core.DeleteClusterUsecase,
) *clusterController {
	return &clusterController{
		saveClusterUsecase:   saveClusterUsecase,
		deleteClusterUsecase: deleteClusterUsecase,
	}
}

func (ctrl *clusterController) SaveCluster(saveClusterUsecaseDTO controllerDTOs.SaveCluster) (*entities.Cluster, error) {
	return ctrl.saveClusterUsecase.Execute(
		saveClusterUsecaseDTO.DomainName,
		saveClusterUsecaseDTO.ClusterName,
	)
}

func (ctrl *clusterController) DeleteCluster(deleteClusterUsecaseDTO controllerDTOs.DeleteCluster) error {
	return ctrl.deleteClusterUsecase.Execute(
		deleteClusterUsecaseDTO.DomainName,
		deleteClusterUsecaseDTO.ClusterName,
	)
}
