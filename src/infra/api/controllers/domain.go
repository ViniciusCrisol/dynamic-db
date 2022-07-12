package controllers

import (
	"github.com/ViniciusCrisol/dynamic-db/app/core"
	"github.com/ViniciusCrisol/dynamic-db/app/entities"
)

type domainController struct {
	saveDomainUsecase   core.SaveDomainUsecase
	deleteDomainUsecase core.DeleteDomainUsecase
}

func NewDomainController(
	saveDomainUsecase core.SaveDomainUsecase,
	deleteDomainUsecase core.DeleteDomainUsecase,
) *domainController {
	return &domainController{
		saveDomainUsecase:   saveDomainUsecase,
		deleteDomainUsecase: deleteDomainUsecase,
	}
}

func (ctrl *domainController) SaveDomain(domainName string) (*entities.Domain, error) {
	return ctrl.saveDomainUsecase.Execute(domainName)
}

func (ctrl *domainController) DeleteDomain(domainName string) error {
	return ctrl.deleteDomainUsecase.Execute(domainName)
}
