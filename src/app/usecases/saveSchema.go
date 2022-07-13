package usecases

import (
	"errors"

	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type saveSchemaUsecase struct {
	clusterRepository         core.ClusterRepository
	assembleClusterURLService core.AssembleClusterURLService
}

func NewSaveSchemaUsecase(
	clusterRepository core.ClusterRepository,
	assembleClusterURLService core.AssembleClusterURLService,
) *saveSchemaUsecase {
	return &saveSchemaUsecase{
		clusterRepository:         clusterRepository,
		assembleClusterURLService: assembleClusterURLService,
	}
}

func (ucs *saveSchemaUsecase) Execute(domainName, clusterName string, schemaContent entities.SchemaContent) (*entities.Schema, error) {
	clusterURL := ucs.assembleClusterURLService.Execute(domainName, clusterName)
	clusterExists := ucs.clusterRepository.ClusterExists(clusterURL)
	if !clusterExists {
		return nil, errors.New("cluster-does-not-exists")
	}

	currentSchemas, readClusterErr := ucs.clusterRepository.ReadCluster(clusterURL)
	if readClusterErr != nil {
		return nil, readClusterErr
	}

	schema := entities.NewSchema(schemaContent)
	updatedSchemas := append(currentSchemas, schema)

	saveClusterErr := ucs.clusterRepository.SaveCluster(clusterURL, updatedSchemas)
	if saveClusterErr != nil {
		utils.ErrorLogger.Println(saveClusterErr)
	}
	return schema, saveClusterErr
}
