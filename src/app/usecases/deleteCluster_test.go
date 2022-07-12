package usecases_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareDeleteClusterEnv(domainURL, clusterURL string) {
	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	domainRepository.SaveDomain(domainURL)
	clusterRepository.SaveCluster(clusterURL, nil)
}

func TestDeleteCluster(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to delete a cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteClusterEnv(domainURL, clusterURL)

		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterErr := deleteClusterUsecase.Execute("domain_name", "cluster_name")
		if deleteClusterErr != nil {
			test.Error(deleteClusterErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to delete a cluster in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		invalidAssembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteClusterEnv(domainURL, clusterURL)

		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, invalidAssembleClusterURLService)
		deleteClusterErr := deleteClusterUsecase.Execute("domain_name", "cluster_name")
		if deleteClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[Execute] - it should not be able to delete a cluster in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteClusterEnv(domainURL, clusterURL)

		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterErr := deleteClusterUsecase.Execute("invalid_domain_name", "cluster_name")
		if deleteClusterErr == nil {
			test.Error()
		}
	})
}
