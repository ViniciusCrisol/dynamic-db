package services

import "github.com/ViniciusCrisol/dynamic-db/infra/repositories"

type assembleClusterURLService struct {
	rootStorageDir string
}

func NewAssembleClusterURLService(rootStorageDir string) *assembleClusterURLService {
	return &assembleClusterURLService{rootStorageDir}
}

func (svc *assembleClusterURLService) Execute(domainName, clusterName string) string {
	return svc.rootStorageDir + "/" + domainName + "/" + clusterName + repositories.CLUSTER_EXTENSION
}
