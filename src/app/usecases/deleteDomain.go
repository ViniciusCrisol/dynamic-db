package usecases

import (
	"errors"

	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type deleteDomainUsecase struct {
	domainRepository         core.DomainRepository
	assembleDomainURLService core.AssembleDomainURLService
}

func NewDeleteDomainUsecase(
	domainRepository core.DomainRepository,
	assembleDomainURLService core.AssembleDomainURLService,
) *deleteDomainUsecase {
	return &deleteDomainUsecase{
		domainRepository:         domainRepository,
		assembleDomainURLService: assembleDomainURLService,
	}
}

func (ucs *deleteDomainUsecase) Execute(domainName string) error {
	domainURL := ucs.assembleDomainURLService.Execute(domainName)
	domainExists := ucs.domainRepository.DomainExists(domainURL)
	if !domainExists {
		return errors.New("domain-does-not-exists")
	}

	deleteDomainErr := ucs.domainRepository.DeleteDomain(domainURL)
	if deleteDomainErr != nil {
		utils.ErrorLogger.Println(deleteDomainErr)
	}
	return deleteDomainErr
}
