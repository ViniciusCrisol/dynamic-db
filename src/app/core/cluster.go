package core

import "github.com/ViniciusCrisol/dynamic-db/app/entities"

type ClusterRepository interface {
	ReadCluster(clusterURL string) ([]*entities.Schema, error)
	SaveCluster(clusterURL string, clusterSchemas []*entities.Schema) error
	ClusterExists(clusterURL string) bool
	DeleteCluster(clusterURL string) error
}

type AssembleClusterURLService interface {
	Execute(domainName, clusterName string) string
}

type SaveClusterUsecase interface {
	Execute(domainName, clusterName string) (*entities.Cluster, error)
}

type DeleteClusterUsecase interface {
	Execute(domainName, clusterName string) error
}
