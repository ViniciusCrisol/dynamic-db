package repositories_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareClusterRepositoryEnv(domainURL string) {
	domainRepository := repositories.NewDomainRepository()
	domainRepository.SaveDomain(domainURL)
}

func TestClusterRepository(testSwitch *testing.T) {
	testSwitch.Run("[ReadCluster] - it should be able to read a cluster without schemas", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		schemas, readClusterErr := clusterRepository.ReadCluster(clusterURL)
		if readClusterErr != nil || len(schemas) != 0 {
			test.Error(readClusterErr)
		}
	})

	testSwitch.Run("[ReadCluster] - it should be able to read a cluster with schemas", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)
		clusterContent := []*entities.Schema{
			entities.NewSchema(entities.SchemaContent{"name": "Emma Stone", "role": "Mia", "rate": 5.0}),
			entities.NewSchema(entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}),
		}
		clusterRepository.SaveCluster(clusterURL, clusterContent)

		schemas, readClusterErr := clusterRepository.ReadCluster(clusterURL)
		if readClusterErr != nil || len(schemas) != 2 {
			test.Error(readClusterErr)
		}
	})

	testSwitch.Run("[ReadCluster] - it should not be able to read a cluster in an invalid dir", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)
		invalidAssembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		invalidClusterURL := invalidAssembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		_, readClusterErr := clusterRepository.ReadCluster(invalidClusterURL)
		if readClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[ReadCluster] - it should not be able to read a cluster in an invalid domain", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("invalid_domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		_, readClusterErr := clusterRepository.ReadCluster(clusterURL)
		if readClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveCluster] - it should be able to save a cluster without schemas", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)

		saveClusterErr := clusterRepository.SaveCluster(clusterURL, nil)
		if saveClusterErr != nil {
			test.Error(saveClusterErr)
		}
	})

	testSwitch.Run("[SaveCluster] - it should be able to save a cluster with schemas", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)

		clusterContent := []*entities.Schema{
			entities.NewSchema(entities.SchemaContent{"name": "Emma Stone", "role": "Mia", "rate": 5.0}),
			entities.NewSchema(entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}),
		}
		saveClusterErr := clusterRepository.SaveCluster(clusterURL, clusterContent)
		if saveClusterErr != nil {
			test.Error(saveClusterErr)
		}
	})

	testSwitch.Run("[SaveCluster] - it should not be able to save a cluster in an invalid dir", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)

		saveClusterErr := clusterRepository.SaveCluster(clusterURL, nil)
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveCluster] - it should not be able to save a cluster in an invalid domain", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("invalid_domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)

		saveClusterErr := clusterRepository.SaveCluster(clusterURL, nil)
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[DomainExists] - it should be able to verify if a cluster exists and return true", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		clusterExists := clusterRepository.ClusterExists(clusterURL)
		if !clusterExists {
			test.Error()
		}
	})

	testSwitch.Run("[DomainExists] - it should be able to verify if a cluster exists and return false", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)

		clusterExists := clusterRepository.ClusterExists(clusterURL)
		if clusterExists {
			test.Error()
		}
	})

	testSwitch.Run("[SaveCluster] - it should be able to delete a cluster", func(test *testing.T) {
		tmpStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tmpStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tmpStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareClusterRepositoryEnv(domainURL)
		clusterRepository.SaveCluster(clusterURL, nil)

		saveClusterErr := clusterRepository.DeleteCluster(clusterURL)
		if saveClusterErr != nil {
			test.Error(saveClusterErr)
		}
	})
}
