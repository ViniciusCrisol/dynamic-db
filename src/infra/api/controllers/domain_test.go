package controllers_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareDeleteDomainEnv(domainURL string) {
	domainRepository := repositories.NewDomainRepository()
	domainRepository.SaveDomain(domainURL)
}

func TestDomain(testSwitch *testing.T) {
	testSwitch.Run("[SaveDomain] - it should be able to save a domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

		domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
		_, saveDomainErr := domainController.SaveDomain("domain_name")
		if saveDomainErr != nil {
			test.Error(saveDomainErr)
		}
	})

	testSwitch.Run("[SaveDomain] - it should not be able to save a domain with a repeated name", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

		domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
		domainController.SaveDomain("domain_name")

		_, saveDomainErr := domainController.SaveDomain("domain_name")
		if saveDomainErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveDomain] - it should not be able to save a domain in an invalid dir", func(test *testing.T) {
		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService("invalid_dir")
		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

		domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
		_, saveDomainErr := domainController.SaveDomain("domain_name")
		if saveDomainErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[DeleteDomain] - it should be able to delete a domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareDeleteDomainEnv(domainURL)

		domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
		deleteDomainErr := domainController.DeleteDomain("domain_name")
		if deleteDomainErr != nil {
			test.Error(deleteDomainErr)
		}
	})

	testSwitch.Run("[DeleteDomain] - it should be able to delete a domain with clusters inside", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteDomainEnv(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
		deleteDomainErr := domainController.DeleteDomain("domain_name")
		if deleteDomainErr != nil {
			test.Error(deleteDomainErr)
		}
	})

	testSwitch.Run("[DeleteDomain] - it should not be able to delete a domain in an invalid dir", func(test *testing.T) {
		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService("invalid_dir")
		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

		domainURL := assembleDomainURLService.Execute("domain_name")
		prepareDeleteDomainEnv(domainURL)

		domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
		deleteDomainErr := domainController.DeleteDomain("domain_name")
		if deleteDomainErr == nil {
			test.Error()
		}
	})
}
