package entities

type Cluster struct {
	DomainName  string `json:"domain_name"`
	ClusterURL  string `json:"-"`
	ClusterName string `json:"cluster_name"`
}

func NewCluster(domainName, clusterName, clusterURL string) *Cluster {
	return &Cluster{
		DomainName:  domainName,
		ClusterURL:  clusterURL,
		ClusterName: clusterName,
	}
}
