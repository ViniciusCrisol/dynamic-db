package controllerDTOs

type SaveCluster struct {
	DomainName  string
	ClusterName string
}

type DeleteCluster struct {
	DomainName  string
	ClusterName string
}
