package controllers_test

import (
	"testing"

	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
)

func prepareSaveClusterEnv(domainURL string) {
	domainRepository := repositories.NewDomainRepository()
	domainRepository.SaveDomain(domainURL)
}

func prepareDeleteClusterEnv(domainURL, clusterURL string) {
	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	domainRepository.SaveDomain(domainURL)
	clusterRepository.SaveCluster(clusterURL, nil)
}

func TestCluster(testSwitch *testing.T) {
	testSwitch.Run("[SaveCluster] - it should be able to save a cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		saveClusterDTO := controllerDTOs.SaveCluster{DomainName: "domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")

		prepareSaveClusterEnv(domainURL)

		_, saveClusterErr := clusterController.SaveCluster(saveClusterDTO)
		if saveClusterErr != nil {
			test.Error(saveClusterErr)
		}
	})

	testSwitch.Run("[SaveCluster] - it should not be able to save a cluster with a repeated name", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		saveClusterDTO := controllerDTOs.SaveCluster{DomainName: "domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")

		prepareSaveClusterEnv(domainURL)
		clusterController.SaveCluster(saveClusterDTO)

		_, saveClusterErr := clusterController.SaveCluster(saveClusterDTO)
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveCluster] - it should not be able to save a cluster in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")
		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		saveClusterDTO := controllerDTOs.SaveCluster{DomainName: "domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")

		prepareSaveClusterEnv(domainURL)

		_, saveClusterErr := clusterController.SaveCluster(saveClusterDTO)
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[SaveCluster] - it should not be able to save a cluster in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		saveClusterDTO := controllerDTOs.SaveCluster{DomainName: "invalid_domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")

		prepareSaveClusterEnv(domainURL)

		_, saveClusterErr := clusterController.SaveCluster(saveClusterDTO)
		if saveClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[DeleteCluster] - it should be able to delete a cluster", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		deleteClusterDTO := controllerDTOs.DeleteCluster{DomainName: "domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteClusterEnv(domainURL, clusterURL)

		deleteClusterErr := clusterController.DeleteCluster(deleteClusterDTO)
		if deleteClusterErr != nil {
			test.Error(deleteClusterErr)
		}
	})

	testSwitch.Run("[DeleteCluster] - it should not be able to delete a cluster in an invalid dir", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)
		invalidAssembleClusterURLService := services.NewAssembleClusterURLService("invalid_dir")
		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, invalidAssembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, invalidAssembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		deleteClusterDTO := controllerDTOs.DeleteCluster{DomainName: "domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteClusterEnv(domainURL, clusterURL)

		deleteClusterErr := clusterController.DeleteCluster(deleteClusterDTO)
		if deleteClusterErr == nil {
			test.Error()
		}
	})

	testSwitch.Run("[DeleteCluster] - it should not be able to delete a cluster in an invalid domain", func(test *testing.T) {
		tempStorageDir := test.TempDir()

		clusterRepository := repositories.NewClusterRepository()
		assembleDomainURLService := services.NewAssembleDomainURLService(tempStorageDir)
		assembleClusterURLService := services.NewAssembleClusterURLService(tempStorageDir)

		saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
		deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

		clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)

		deleteClusterDTO := controllerDTOs.DeleteCluster{DomainName: "invalid_domain_name", ClusterName: "cluster_name"}
		domainURL := assembleDomainURLService.Execute("domain_name")
		clusterURL := assembleClusterURLService.Execute("domain_name", "cluster_name")

		prepareDeleteClusterEnv(domainURL, clusterURL)

		deleteClusterErr := clusterController.DeleteCluster(deleteClusterDTO)
		if deleteClusterErr == nil {
			test.Error()
		}
	})
}
