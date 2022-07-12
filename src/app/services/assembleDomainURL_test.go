package services_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
)

func TestAssembleDomainURL(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to assemble a domain URL", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		domainURL := assembleDomainURLService.Execute("domain_name")
		domainURLToCompare := tempStorageDir + "/domain_name"
		if domainURL != domainURLToCompare {
			test.Error()
		}
	})
}
