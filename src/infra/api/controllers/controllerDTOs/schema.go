package controllerDTOs

import "github.com/ViniciusCrisol/dynamic-db/app/entities"

type SaveSchema struct {
	DomainName    string
	ClusterName   string
	SchemaContent entities.SchemaContent
}
