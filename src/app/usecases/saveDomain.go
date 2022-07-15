package usecases

import (
	"errors"

	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type saveDomainUsecase struct {
	domainRepository         core.DomainRepository
	assembleDomainURLService core.AssembleDomainURLService
}

func NewSaveDomainUsecase(
	domainRepository core.DomainRepository,
	assembleDomainURLService core.AssembleDomainURLService,
) *saveDomainUsecase {
	return &saveDomainUsecase{
		domainRepository:         domainRepository,
		assembleDomainURLService: assembleDomainURLService,
	}
}

func (ucs *saveDomainUsecase) Execute(domainName string) (*entities.Domain, error) {
	domainURL := ucs.assembleDomainURLService.Execute(domainName)
	domainAlreadyExists := ucs.domainRepository.DomainExists(domainURL)
	if domainAlreadyExists {
		return nil, errors.New("domain-already-exists")
	}

	domain := entities.NewDomain(domainName, domainURL)
	saveDomainErr := ucs.domainRepository.SaveDomain(domain.DomainURL)
	if saveDomainErr != nil {
		utils.ErrorLogger.Println(saveDomainErr)
		return nil, saveDomainErr
	}
	return domain, nil
}
