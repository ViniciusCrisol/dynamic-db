package core

import "github.com/ViniciusCrisol/dynamic-db/app/entities"

type SaveSchemaUsecase interface {
	Execute(domainName, clusterName string, schemaContent entities.SchemaContent) (*entities.Schema, error)
}
