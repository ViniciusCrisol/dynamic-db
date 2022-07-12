package usecases_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func TestSaveDomain(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to save a domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)

		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		_, saveDomainErr := saveDomainUsecase.Execute("domain_name")
		if saveDomainErr != nil {
			test.Error(saveDomainErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a domain with a repeated name", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)

		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		saveDomainUsecase.Execute("domain_name")

		_, saveDomainErr := saveDomainUsecase.Execute("domain_name")
		if saveDomainErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a domain in an invalid dir", func(test *testing.T) {
		domainRepository := repositories.NewDomainRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService("invalid_dir")

		saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
		_, saveDomainErr := saveDomainUsecase.Execute("domain_name")
		if saveDomainErr == nil {
			test.Error()
		}
	})
}
