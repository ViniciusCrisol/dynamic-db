package core

import "github.com/ViniciusCrisol/dynamic-db/app/entities"

type SaveSchemaUsecase interface {
	Execute(domainName, clusterName string, schemaContent entities.SchemaContent) (*entities.Schema, error)
}

type FindSchemaUsecase interface {
	Execute(domainName, clusterName string, schemaToMatch entities.SchemaContent) ([]*entities.Schema, error)
}
