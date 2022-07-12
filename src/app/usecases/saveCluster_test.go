package usecases_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareSaveClusterEnv(domainURL string) {
	domainRepository := repositories.NewDomainRepository()
	domainRepository.SaveDomain(domainURL)
}

func TestSaveCluster(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to save a cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareSaveClusterEnv(domainURL)

		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		_, saveClusterErr := saveClusterUsecase.Execute("domain_name", "cluster_name")
		if saveClusterErr != nil {
			test.Error(saveClusterErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a cluster with a repeated name", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareSaveClusterEnv(domainURL)

		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		saveClusterUsecase.Execute("domain_name", "cluster_name")

		_, saveClusterErr := saveClusterUsecase.Execute("domain_name", "cluster_name")
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a cluster in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareSaveClusterEnv(domainURL)

		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		_, saveClusterErr := saveClusterUsecase.Execute("domain_name", "cluster_name")
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a cluster in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareSaveClusterEnv(domainURL)

		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		_, saveClusterErr := saveClusterUsecase.Execute("invalid_domain_name", "cluster_name")
		if saveClusterErr == nil {
			test.Error()
		}
	})
}
