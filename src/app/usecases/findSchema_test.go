package usecases_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareFindSchemaEnv(domainURL, clusterURL string) {
	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	domainRepository.SaveDomain(domainURL)
	clusterRepository.SaveCluster(clusterURL, []*entities.Schema{
		entities.NewSchema(entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}),
		entities.NewSchema(entities.SchemaContent{"name": "Emma Stone", "role": "Mia", "rate": 4.5}),
		entities.NewSchema(entities.SchemaContent{"name": "Lady Gaga", "role": "Ally", "rate": 5.0}),
	})
}

func TestFindSchema(testSwitch *testing.T) {
	testSwitch.Run("[Execute] - it should be able to find a schema with filters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareFindSchemaEnv(domainURL, clusterURL)

		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		schemas, findSchemaErr := findSchemaUsecase.Execute("domain_name", "cluster_name", schemaToMatch)
		if findSchemaErr != nil || len(schemas) != 1 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[Execute] - it should be able to find many schemas with filters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"rate": 4.5}

		prepareFindSchemaEnv(domainURL, clusterURL)

		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		schemas, findSchemaErr := findSchemaUsecase.Execute("domain_name", "cluster_name", schemaToMatch)
		if findSchemaErr != nil || len(schemas) != 2 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[Execute] - it should be able to find many schemas without filters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareFindSchemaEnv(domainURL, clusterURL)

		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		schemas, findSchemaErr := findSchemaUsecase.Execute("domain_name", "cluster_name", nil)
		if findSchemaErr != nil || len(schemas) != 3 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to find a schema in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		invalidAssembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareFindSchemaEnv(domainURL, clusterURL)

		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, invalidAssembleClusterURLService)
		schemas, findSchemaErr := findSchemaUsecase.Execute("domain_name", "cluster_name", schemaToMatch)
		if findSchemaErr == nil || len(schemas) != 0 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to find a schema in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareFindSchemaEnv(domainURL, clusterURL)

		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		schemas, findSchemaErr := findSchemaUsecase.Execute("invalid_domain_name", "cluster_name", schemaToMatch)
		if findSchemaErr == nil || len(schemas) != 0 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[Execute] - it should not be able to find a schema in an invalid cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}

		prepareFindSchemaEnv(domainURL, clusterURL)

		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		schemas, findSchemaErr := findSchemaUsecase.Execute("domain_name", "invalid_cluster_name", schemaToMatch)
		if findSchemaErr == nil || len(schemas) != 0 {
			test.Error(findSchemaErr)
		}
	})
}
