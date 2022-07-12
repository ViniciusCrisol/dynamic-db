package services_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func TestAssembleClusterURL(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to assemble a cluster URL", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		assembleDomainURLService := services.NewAssembleClusterURLService(tempStorageDir)
		domainURL := assembleDomainURLService.Execute("domain_name", "cluster_name")
		domainURLToCompare := tempStorageDir + "/domain_name/cluster_name" + repositories.CLUSTER_EXTENSION
		if domainURL != domainURLToCompare {
			test.Error()
		}
	})
}
