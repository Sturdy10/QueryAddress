package services

import (
	"queryAddress/models"
	"queryAddress/repositories"
)

// type ServicePort interface {
// 	GetCountriesServices() ([]getQueryModels.CountryResponse, error)
// }

type ServicePort interface {
	GetCountriesServices() (*models.CountryResponse, error)
	GetProvinceAmphoeTambonZipcodeServices() (*models.TambonDataResponse, error)
}

type serviceAdapter struct {
	r repositories.RepositoryPort
}

func NewServiceAdapter(r repositories.RepositoryPort) ServicePort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) GetCountriesServices() (*models.CountryResponse, error) { /////// return single obeject complex and not flexible but very fast , the value pass references not copy into the funciton.

	countries, err := s.r.GetCountriesRepositories()
	if err != nil {
		return nil, err
	}
	countriesResponse := models.CountryResponse{
		Countries: countries,
	}
	return &countriesResponse, nil

}

func (s *serviceAdapter) GetProvinceAmphoeTambonZipcodeServices() (*models.TambonDataResponse, error) {
	tambonData, err := s.r.GetProvinceAmphoeTambonZipcodeRepositories()
	if err != nil {
		return nil, err
	}
	tambonDataResponse := models.TambonDataResponse{
		ProvinceAmphoeTambonZipcode: tambonData,
	}
	return &tambonDataResponse, nil
}
