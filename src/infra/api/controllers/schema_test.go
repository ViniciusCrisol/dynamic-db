package controllers_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareSaveSchemaEnv(domainURL, clusterURL string) {
	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	domainRepository.SaveDomain(domainURL)
	clusterRepository.SaveCluster(clusterURL, nil)
}

func TestSchema(testSwitch *testing.T) {
	testSwitch.Run("[SaveSchema] - it should be able to save a schema", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaContent: schemaContent}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		_, saveSchemaErr := clusterController.SaveSchema(saveSchemaDTO)
		if saveSchemaErr != nil {
			test.Error(saveSchemaErr)
		}
	})

	testSwitch.Run("[SaveSchema] - it should not be able to save a schema in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		invalidAssembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, invalidAssembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaContent: schemaContent}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		_, saveSchemaErr := clusterController.SaveSchema(saveSchemaDTO)
		if saveSchemaErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveSchema] - it should not be able to save a schema in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "invalid_domain_name", ClusterName: "cluster_name", SchemaContent: schemaContent}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		_, saveSchemaErr := clusterController.SaveSchema(saveSchemaDTO)
		if saveSchemaErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveSchema] - it should not be able to save a schema in an invalid cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "domain_name", ClusterName: "invalid_cluster_name", SchemaContent: schemaContent}

		prepareSaveSchemaEnv(domainURL, clusterURL)

		_, saveSchemaErr := clusterController.SaveSchema(saveSchemaDTO)
		if saveSchemaErr == nil {
			test.Error()
		}
	})
}
