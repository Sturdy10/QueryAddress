package repositories

import (
	"database/sql"
	"log"
	"queryAddress/models"
)

type RepositoryPort interface {
	GetCountriesRepositories() ([]models.Country, error)
	GetProvinceAmphoeTambonZipcodeRepositories() ([]models.TambonData, error)
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) GetCountriesRepositories() ([]models.Country, error) {
	var countries []models.Country
	rows, err := r.db.Query("SELECT * FROM countries")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err := rows.Scan(&country.ID, &country.Name, &country.TopLevelDomain, &country.Alpha2Code, &country.Alpha3Code, &country.CallingCodes, &country.Capital, &country.Subregion, &country.Region, &country.Population, &country.Latitude, &country.Longitude, &country.Demonym, &country.Area, &country.Gini, &country.Timezones, &country.NativeName, &country.NumericCode, &country.FlagSVG, &country.FlagPNG, &country.CurrencyCode, &country.CurrencyName, &country.CurrencySymbol, &country.LanguageCode, &country.LanguageName, &country.LanguageNativeName, &country.IsIndependent)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func (r *repositoryAdapter) GetProvinceAmphoeTambonZipcodeRepositories() ([]models.TambonData, error) {
	var tambonData []models.TambonData

	rows, err := r.db.Query("SELECT * FROM \"thailandTambon\"") // even uppercase also need \"\"
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data models.TambonData
		if err := rows.Scan(
			&data.ID, &data.TambonID, &data.TambonThai, &data.TambonEng, &data.TambonThaiShort, &data.TambonEngShort,
			&data.DistrictID, &data.DistrictThai, &data.DistrictEng, &data.DistrictThaiShort, &data.DistrictEngShort,
			&data.ProvinceID, &data.ProvinceThai, &data.ProvinceEng, &data.OfficialRegions, &data.FourMainRegions,
			&data.TouristRegions, &data.GreaterBangkokProvinces, &data.PostalCodeRemark, &data.PostCodeMain, &data.PostCodeAll,
		); err != nil {
			log.Println(err)
			return nil, err
		}
		tambonData = append(tambonData, data)
	}
	return tambonData, nil
}
