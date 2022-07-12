package repositories_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func TestDomainRepository(testSwitch *testing.T) {
	testSwitch.Run("[SaveDomain] - it should be able to save a domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		domainURL := assembleDomainURLService.Execute("domain_name")

		domainRepository := repositories.NewDomainRepository()
		saveDomainErr := domainRepository.SaveDomain(domainURL)
		if saveDomainErr != nil {
			test.Error(saveDomainErr)
		}
	})

	testSwitch.Run("[SaveDomain] - it should not be able to save a domain in an invalid dir", func(test *testing.T) {
		assembleDomainURLService := services.NewAssembleDomainURLService("invalid_dir")
		domainURL := assembleDomainURLService.Execute("domain_name")

		domainRepository := repositories.NewDomainRepository()
		saveDomainErr := domainRepository.SaveDomain(domainURL)
		if saveDomainErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[DomainExists] - it should be able to verify if a domain exists and return true", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		domainURL := assembleDomainURLService.Execute("domain_name")

		domainRepository := repositories.NewDomainRepository()
		domainRepository.SaveDomain(domainURL)

		domainExists := domainRepository.DomainExists(domainURL)
		if !domainExists {
			test.Error()
		}
	})

	testSwitch.Run("[DomainExists] - it should be able to verify if a domain exists and return false", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		domainURL := assembleDomainURLService.Execute("domain_name")

		domainRepository := repositories.NewDomainRepository()
		domainExists := domainRepository.DomainExists(domainURL)
		if domainExists {
			test.Error()
		}
	})

	testSwitch.Run("[DeleteDomain] - it should be able to delete a domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		domainURL := assembleDomainURLService.Execute("domain_name")

		domainRepository := repositories.NewDomainRepository()
		domainRepository.SaveDomain(domainURL)

		deleteDomainErr := domainRepository.DeleteDomain(domainURL)
		if deleteDomainErr != nil {
			test.Error()
		}
	})

	testSwitch.Run("[DeleteDomain] - it should be able to delete a domain with clusters inside", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		domainRepository.SaveDomain(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		deleteDomainErr := domainRepository.DeleteDomain(domainURL)
		if deleteDomainErr != nil {
			test.Error()
		}
	})
}
