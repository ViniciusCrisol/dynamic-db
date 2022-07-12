package repositories

import (
	"encoding/gob"
	"os"

	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type clusterRepository struct{}

const CLUSTER_EXTENSION = ".gob"

func NewClusterRepository() *clusterRepository {
	return &clusterRepository{}
}

func (repo *clusterRepository) ReadCluster(clusterURL string) ([]*entities.Schema, error) {
	cluster, readClusterErr := os.Open(clusterURL)
	if readClusterErr != nil {
		utils.ErrorLogger.Println(readClusterErr)
		return nil, readClusterErr
	}
	defer cluster.Close()

	schemas := []*entities.Schema{}
	decoder := gob.NewDecoder(cluster)
	decodeSchemasErr := decoder.Decode(&schemas)
	if decodeSchemasErr != nil {
		utils.ErrorLogger.Println(decodeSchemasErr)
	}
	return schemas, decodeSchemasErr
}

func (repo *clusterRepository) SaveCluster(clusterURL string, clusterSchemas []*entities.Schema) error {
	cluster, saveClusterErr := os.Create(clusterURL)
	if saveClusterErr != nil {
		utils.ErrorLogger.Println(saveClusterErr)
		return saveClusterErr
	}
	defer cluster.Close()

	encoder := gob.NewEncoder(cluster)
	encodeSchemasErr := encoder.Encode(clusterSchemas)
	if encodeSchemasErr != nil {
		utils.ErrorLogger.Println(encodeSchemasErr)
	}
	return encodeSchemasErr
}

func (repo *clusterRepository) ClusterExists(clusterURL string) bool {
	_, clusterStatus := os.Stat(clusterURL)
	clusterNotExists := os.IsNotExist(clusterStatus)
	return !clusterNotExists
}

func (repo *clusterRepository) DeleteCluster(clusterURL string) error {
	deleteClusterErr := os.Remove(clusterURL)
	if deleteClusterErr != nil {
		utils.ErrorLogger.Println(deleteClusterErr)
	}
	return deleteClusterErr
}
