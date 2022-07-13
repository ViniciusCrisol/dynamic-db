package main

import (
	"github.com/ViniciusCrisol/dynamic-db/app/services"
	"github.com/ViniciusCrisol/dynamic-db/app/usecases"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/handlers/middlewares"
	"github.com/ViniciusCrisol/dynamic-db/infra/repositories"
	"github.com/ViniciusCrisol/dynamic-db/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

const (
	ROOT_STORAGE_DIR   = "data"
	APPLICATION_PORT   = "8080"
	HTTP_ROUTES_PREFIX = "api/v1"
)

func main() {
	configApplication()

	router := getRouteHandler()
	setupHTTPHandlers(router)
	listenToHTTPRequests(router)
}

func configApplication() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

func getRouteHandler() *gin.Engine {
	utils.InfoLogger.Println("setting up HTTP listener")

	router := gin.Default()
	responseCompressor := gzip.Gzip(gzip.BestCompression)

	router.Use(responseCompressor)
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:  []string{"*"},
				AllowMethods:  []string{"*"},
				AllowHeaders:  []string{"*"},
				AllowWildcard: true,
			},
		),
	)
	router.NoRoute(middlewares.SendRouteNotFound)
	return router
}

func setupHTTPHandlers(router *gin.Engine) {
	utils.InfoLogger.Println("setting up HTTP handlers")

	apiRouter := router.Group(HTTP_ROUTES_PREFIX)

	domainRepository := repositories.NewDomainRepository()
	clusterRepository := repositories.NewClusterRepository()
	assembleDomainURLService := services.NewAssembleDomainURLService(ROOT_STORAGE_DIR)
	assembleClusterURLService := services.NewAssembleClusterURLService(ROOT_STORAGE_DIR)

	saveDomainUsecase := usecases.NewSaveDomainUsecase(domainRepository, assembleDomainURLService)
	deleteDomainUsecase := usecases.NewDeleteDomainUsecase(domainRepository, assembleDomainURLService)

	domainController := controllers.NewDomainController(saveDomainUsecase, deleteDomainUsecase)
	domainHandler := handlers.NewDomainHandler(domainController)
	apiRouter.POST("/domains", domainHandler.SaveDomain)
	apiRouter.DELETE("/domains/:domain_name", domainHandler.DeletedDomain)

	saveClusterUsecase := usecases.NewSaveClusterUsecase(clusterRepository, assembleClusterURLService)
	deleteClusterUsecase := usecases.NewDeleteClusterUsecase(clusterRepository, assembleClusterURLService)

	clusterController := controllers.NewClusterController(saveClusterUsecase, deleteClusterUsecase)
	clusterHandler := handlers.NewClusterHandler(clusterController)
	apiRouter.POST("/clusters/:domain_name", clusterHandler.SaveCluster)
	apiRouter.DELETE("/clusters/:domain_name/:cluster_name", clusterHandler.DeleteCluster)

	saveSchemaUsecase := usecases.NewSaveSchemaUsecase(clusterRepository, assembleClusterURLService)
	findSchemaUsecase := usecases.NewFindSchemaUsecase(clusterRepository, assembleClusterURLService)
	schemaController := controllers.NewSchemaController(saveSchemaUsecase, findSchemaUsecase)
	schemaHandler := handlers.NewSchemaHandler(schemaController)
	apiRouter.POST("/schemas/:domain_name/:cluster_name", schemaHandler.SaveSchema)
	apiRouter.GET("/schemas/:domain_name/:cluster_name", schemaHandler.FindSchema)
}

func listenToHTTPRequests(routerHTTP *gin.Engine) {
	utils.InfoLogger.Println("listening to HTTP requests")

	initHttpServerErr := routerHTTP.Run(":" + APPLICATION_PORT)
	if initHttpServerErr != nil {
		utils.PanicLogger.Panic(initHttpServerErr)
	}
}
