package core

import "github.com/ViniciusCrisol/dynamic-db/app/entities"

type DomainRepository interface {
	SaveDomain(domainURL string) error
	DomainExists(domainURL string) bool
	DeleteDomain(domainURL string) error
}

type AssembleDomainURLService interface {
	Execute(domainName string) string
}

type SaveDomainUsecase interface {
	Execute(domainName string) (*entities.Domain, error)
}

type DeleteDomainUsecase interface {
	Execute(domainName string) error
}
