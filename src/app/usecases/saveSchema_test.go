package usecases_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareSaveSchemaEnv(domainURL, clusterURL string) {
	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	domainRepository.SaveDomain(domainURL)
	clusterRepository.SaveCluster(clusterURL, nil)
}

func TestSaveSchema(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to save a schema", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		_, saveSchemaErr := saveSchemaUsecase.Execute("domain_name", "cluster_name", schemaContent)
		if saveSchemaErr != nil {
			test.Error(saveSchemaErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a schema in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		invalidAssembleClusterURL := services.NewAssembleClusterURLService("invalid_dir")

		prepareSaveSchemaEnv(domainURL, clusterURL)
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, invalidAssembleClusterURL)
		_, saveSchemaErr := saveSchemaUsecase.Execute("domain_name", "cluster_name", schemaContent)
		if saveSchemaErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a schema in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		_, saveSchemaErr := saveSchemaUsecase.Execute("invalid_domain_name", "cluster_name", schemaContent)
		if saveSchemaErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[Execute] - it should not be able to save a schema in an invalid cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		_, saveSchemaErr := saveSchemaUsecase.Execute("domain_name", "invalid_cluster_name", schemaContent)
		if saveSchemaErr == nil {
			test.Error()
		}
	})
}
