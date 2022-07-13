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

func prepareSaveAndFindSchemaEnv(domainURL, clusterURL string) {
	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	domainRepository.SaveDomain(domainURL)
	clusterRepository.SaveCluster(clusterURL, []*entities.Schema{
		entities.NewSchema(entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}),
		entities.NewSchema(entities.SchemaContent{"name": "Emma Stone", "role": "Mia", "rate": 4.5}),
		entities.NewSchema(entities.SchemaContent{"name": "Lady Gaga", "role": "Ally", "rate": 5.0}),
	})
}

func TestSchema(testSwitch *testing.T) {
	testSwitch.Run("[SaveSchema] - it should be able to save a schema", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaContent: schemaContent}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

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
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaContent: schemaContent}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

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
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "invalid_domain_name", ClusterName: "cluster_name", SchemaContent: schemaContent}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

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
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaContent := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		saveSchemaDTO := controllerDTOs.SaveSchema{DomainName: "domain_name", ClusterName: "invalid_cluster_name", SchemaContent: schemaContent}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		_, saveSchemaErr := clusterController.SaveSchema(saveSchemaDTO)
		if saveSchemaErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[FindSchema] - it should be able to find a schema with filters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		findSchemaDTO := controllerDTOs.FindSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaToMatch: schemaToMatch}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		schemas, findSchemaErr := clusterController.FindSchema(findSchemaDTO)
		if findSchemaErr != nil || len(schemas) != 1 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[FindSchema] - it should be able to find many schemas with filters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"rate": 4.5}
		findSchemaDTO := controllerDTOs.FindSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaToMatch: schemaToMatch}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		schemas, findSchemaErr := clusterController.FindSchema(findSchemaDTO)
		if findSchemaErr != nil || len(schemas) != 2 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[FindSchema] - it should be able to find many schemas without filters", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		findSchemaDTO := controllerDTOs.FindSchema{DomainName: "domain_name", ClusterName: "cluster_name"}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		schemas, findSchemaErr := clusterController.FindSchema(findSchemaDTO)
		if findSchemaErr != nil || len(schemas) != 3 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[FindSchema] - it should not be able to find a schema in an invalid dir", func(test *testing.T) {

		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		invalidAssembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, invalidAssembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		findSchemaDTO := controllerDTOs.FindSchema{DomainName: "domain_name", ClusterName: "cluster_name", SchemaToMatch: schemaToMatch}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		schemas, findSchemaErr := clusterController.FindSchema(findSchemaDTO)
		if findSchemaErr == nil || len(schemas) != 0 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[FindSchema] - it should not be able to find a schema in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		findSchemaDTO := controllerDTOs.FindSchema{DomainName: "invalid_domain_name", ClusterName: "cluster_name", SchemaToMatch: schemaToMatch}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		schemas, findSchemaErr := clusterController.FindSchema(findSchemaDTO)
		if findSchemaErr == nil || len(schemas) != 0 {
			test.Error(findSchemaErr)
		}
	})

	testSwitch.Run("[FindSchema] - it should not be able to find a schema in an invalid cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
		findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
		clusterController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)

		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")
		schemaToMatch := entities.SchemaContent{"name": "Emma Watson", "role": "Hermione Granger", "rate": 4.5}
		findSchemaDTO := controllerDTOs.FindSchema{DomainName: "domain_name", ClusterName: "invalid_cluster_name", SchemaToMatch: schemaToMatch}

		prepareSaveAndFindSchemaEnv(domainURL, clusterURL)

		schemas, findSchemaErr := clusterController.FindSchema(findSchemaDTO)
		if findSchemaErr == nil || len(schemas) != 0 {
			test.Error(findSchemaErr)
		}
	})
}
