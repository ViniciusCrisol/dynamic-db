package services

type assembleDomainURLService struct {
	rootStorageDir string
}

func NewAssembleDomainURLService(rootStorageDir string) *assembleDomainURLService {
	return &assembleDomainURLService{rootStorageDir}
}

func (svc *assembleDomainURLService) Execute(domainName string) string {
	return svc.rootStorageDir + "/" + domainName
}
