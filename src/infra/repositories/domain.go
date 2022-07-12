package repositories

import (
	"os"

	"github.com/ViniciusCrisol/dynamic-db/utils"
)

type domainRepository struct{}

func NewDomainRepository() *domainRepository {
	return &domainRepository{}
}

func (repo *domainRepository) SaveDomain(domainURL string) error {
	saveDomainErr := os.Mkdir(domainURL, os.ModePerm)
	if saveDomainErr != nil {
		utils.ErrorLogger.Println(saveDomainErr)
	}
	return saveDomainErr
}

func (repo *domainRepository) DomainExists(domainURL string) bool {
	_, domainStatus := os.Stat(domainURL)
	domainNotExists := os.IsNotExist(domainStatus)
	return !domainNotExists
}

func (repo *domainRepository) DeleteDomain(domainURL string) error {
	deleteDomainErr := os.RemoveAll(domainURL)
	if deleteDomainErr != nil {
		utils.ErrorLogger.Println(deleteDomainErr)
	}
	return deleteDomainErr
}
