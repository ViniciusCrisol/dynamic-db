package controllers

import (
	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
)

type schemaController struct {
	saveSchemaUsecase core.SaveSchemaUsecase
	findSchemaUsecase core.FindSchemaUsecase
}

func NewSchemaController(saveSchemaUsecase core.SaveSchemaUsecase, findSchemaUsecase core.FindSchemaUsecase) *schemaController {
	return &schemaController{
		saveSchemaUsecase: saveSchemaUsecase,
		findSchemaUsecase: findSchemaUsecase,
	}
}

func (ctrl *schemaController) SaveSchema(saveSchemaDTO controllerDTOs.SaveSchema) (*entities.Schema, error) {
	return ctrl.saveSchemaUsecase.Execute(
		saveSchemaDTO.DomainName,
		saveSchemaDTO.ClusterName,
		saveSchemaDTO.SchemaContent,
	)
}

func (ctrl *schemaController) FindSchema(saveSchemaDTO controllerDTOs.FindSchema) ([]*entities.Schema, error) {
	return ctrl.findSchemaUsecase.Execute(
		saveSchemaDTO.DomainName,
		saveSchemaDTO.ClusterName,
		saveSchemaDTO.SchemaToMatch,
	)
}
