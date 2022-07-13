package usecases

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
)

type findSchemaUsecase struct {
	clusterRepository         core.ClusterRepository
	assembleClusterURLService core.AssembleClusterURLService
}

func NewFindSchemaUsecase(
	clusterRepository core.ClusterRepository,
	assembleClusterURLService core.AssembleClusterURLService,
) *findSchemaUsecase {
	return &findSchemaUsecase{
		clusterRepository:         clusterRepository,
		assembleClusterURLService: assembleClusterURLService,
	}
}

func (ucs *findSchemaUsecase) Execute(domainName, clusterName string, schemaToMatch entities.SchemaContent) ([]*entities.Schema, error) {
	clusterURL := ucs.assembleClusterURLService.Execute(domainName, clusterName)
	clusterExists := ucs.clusterRepository.ClusterExists(clusterURL)
	if !clusterExists {
		return nil, errors.New("cluster-does-not-exists")
	}

	currentSchemas, readClusterErr := ucs.clusterRepository.ReadCluster(clusterURL)
	if readClusterErr != nil {
		return nil, readClusterErr
	}

	if schemaToMatch == nil {
		return currentSchemas, nil
	}

	filteredSchemas := ucs.filterSchemas(currentSchemas, schemaToMatch)
	return filteredSchemas, nil
}

func (ucs *findSchemaUsecase) filterSchemas(schemasToFilter []*entities.Schema, schemaToMatch entities.SchemaContent) []*entities.Schema {
	filteredSchemas := []*entities.Schema{}
	for _, schema := range schemasToFilter {
		schemaMatches := ucs.schemaMatches(schema, schemaToMatch)
		if !schemaMatches {
			continue
		}
		filteredSchemas = append(filteredSchemas, schema)
	}
	return filteredSchemas
}

func (ucs *findSchemaUsecase) schemaMatches(schema *entities.Schema, schemaToMatch entities.SchemaContent) bool {
	for field, valueToMatch := range schemaToMatch {
		sanitizedField := ucs.sanitizeValueToFilter(schema.SchemaContent[field])
		sanitizedValue := ucs.sanitizeValueToFilter(valueToMatch)
		if sanitizedField != sanitizedValue {
			return false
		}
	}
	return true
}

func (ucs *findSchemaUsecase) sanitizeValueToFilter(rawValue interface{}) string {
	parsedValue := fmt.Sprint(rawValue)
	return strings.ToLower(parsedValue)
}
