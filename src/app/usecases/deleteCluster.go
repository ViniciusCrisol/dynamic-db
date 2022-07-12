package usecases

import (
	"errors"

	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type deleteClusterUsecase struct {
	clusterRepository         core.ClusterRepository
	assembleClusterURLService core.AssembleClusterURLService
}

func NewDeleteClusterUsecase(
	clusterRepository core.ClusterRepository,
	assembleClusterURLService core.AssembleClusterURLService,
) *deleteClusterUsecase {
	return &deleteClusterUsecase{
		clusterRepository:         clusterRepository,
		assembleClusterURLService: assembleClusterURLService,
	}
}

func (ucs *deleteClusterUsecase) Execute(domainName, clusterName string) error {
	clusterURL := ucs.assembleClusterURLService.Execute(domainName, clusterName)
	clusterExists := ucs.clusterRepository.ClusterExists(clusterURL)
	if !clusterExists {
		return errors.New("cluster-does-not-exists")
	}

	deleteClusterErr := ucs.clusterRepository.DeleteCluster(clusterURL)
	if deleteClusterErr != nil {
		utils.ErrorLogger.Println(deleteClusterErr)
	}
	return deleteClusterErr
}
