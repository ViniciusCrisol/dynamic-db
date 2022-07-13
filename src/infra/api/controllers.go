package api

import (
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
)

type SchemaController interface {
	SaveSchema(saveSchemaDTO controllerDTOs.SaveSchema) (*entities.Schema, error)
	FindSchema(findSchemaDTO controllerDTOs.FindSchema) ([]*entities.Schema, error)
}

type DomainController interface {
	SaveDomain(domainName string) (*entities.Domain, error)
	DeleteDomain(domainName string) error
}

type ClusterController interface {
	SaveCluster(createClusterDTO controllerDTOs.SaveCluster) (*entities.Cluster, error)
	DeleteCluster(deleteClusterDTO controllerDTOs.DeleteCluster) error
}
