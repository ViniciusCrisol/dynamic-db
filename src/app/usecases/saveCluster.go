package usecases

import (
	"errors"

	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type saveClusterUsecase struct {
	clusterRepository         core.ClusterRepository
	assembleClusterURLService core.AssembleClusterURLService
}

func NewSaveClusterUsecase(
	clusterRepository core.ClusterRepository,
	assembleClusterURLService core.AssembleClusterURLService,
) *saveClusterUsecase {
	return &saveClusterUsecase{
		clusterRepository:         clusterRepository,
		assembleClusterURLService: assembleClusterURLService,
	}
}

func (ucs *saveClusterUsecase) Execute(domainName, clusterName string) (*entities.Cluster, error) {
	clusterURL := ucs.assembleClusterURLService.Execute(domainName, clusterName)
	clusterAlreadyExists := ucs.clusterRepository.ClusterExists(clusterURL)
	if clusterAlreadyExists {
		return nil, errors.New("cluster-already-exists")
	}

	cluster := entities.NewCluster(domainName, clusterName, clusterURL)
	saveClusterErr := ucs.clusterRepository.SaveCluster(cluster.ClusterURL, nil)
	if saveClusterErr != nil {
		utils.ErrorLogger.Println(saveClusterErr)
		return nil, saveClusterErr
	}
	return cluster, nil
}
