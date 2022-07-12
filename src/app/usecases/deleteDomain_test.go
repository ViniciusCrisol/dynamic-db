package usecases_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareDeleteDomainEnv(domainURL string) {
	domainRepository := repositories.NewDomainRepository()
	domainRepository.SaveDomain(domainURL)
}

func TestDeleteDomain(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to delete a domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareDeleteDomainEnv(domainURL)

		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainErr := deleteDomainUsecase.Execute("domain_name")
		if deleteDomainErr != nil {
			test.Error(deleteDomainErr)
		}
	})

	testSwitch.Run("[Execute] - it should be able to delete a domain with clusters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")

		prepareDeleteDomainEnv(domainURL)

		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		clusterRepository.SaveCluster(clusterURL, nil)

		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainErr := deleteDomainUsecase.Execute("domain_name")
		if deleteDomainErr != nil {
			test.Error(deleteDomainErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to delete a domain in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		invalidAssembleDomainURLService := services.NewAssembleDomainURLService("invalid_dir")

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareDeleteDomainEnv(domainURL)

		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, invalidAssembleDomainURLService)
		deleteDomainErr := deleteDomainUsecase.Execute("domain_name")
		if deleteDomainErr == nil {
			test.Error()
		}
	})
}
