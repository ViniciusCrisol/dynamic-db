package entities

import (
	"time"
)

type SchemaContent map[string]interface{}

type SchemaMetadata struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Schema struct {
	SchemaContent  SchemaContent  `json:"schema_content"`
	SchemaMetadata SchemaMetadata `json:"-"`
}

func NewSchema(schemaContent SchemaContent) *Schema {
	return &Schema{
		SchemaContent: schemaContent,
		SchemaMetadata: SchemaMetadata{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
