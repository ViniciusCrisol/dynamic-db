package entities

type Domain struct {
	DomainURL  string `json:"-"`
	DomainName string `json:"domain_name"`
}

func NewDomain(domainName, domainURL string) *Domain {
	return &Domain{
		DomainURL:  domainURL,
		DomainName: domainName,
	}
}
