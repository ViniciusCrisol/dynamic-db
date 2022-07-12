package controllers

import (
	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/infra/api/controllers/controllerDTOs"
)

type schemaController struct {
	saveSchemaUsecase core.SaveSchemaUsecase
}

func NewSchemaController(saveSchemaUsecase core.SaveSchemaUsecase) *schemaController {
	return &schemaController{saveSchemaUsecase}
}

func (ctrl *schemaController) SaveSchema(saveSchemaDTO controllerDTOs.SaveSchema) (*entities.Schema, error) {
	return ctrl.saveSchemaUsecase.Execute(
		saveSchemaDTO.DomainName,
		saveSchemaDTO.ClusterName,
		saveSchemaDTO.SchemaContent,
	)
}
