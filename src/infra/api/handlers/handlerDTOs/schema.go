package handlerDTOs

import "github.com/ViniciusCrisol/dynamic-db/app/entities"

type SaveSchema struct {
	SchemaContent entities.SchemaContent `json:"schema_content"`
}
