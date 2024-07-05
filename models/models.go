package models



type Country struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	TopLevelDomain     string  `json:"topLevelDomain"`
	Alpha2Code         string  `json:"alpha2Code"`
	Alpha3Code         string  `json:"alpha3Code"`
	CallingCodes       string  `json:"callingCodes"`
	Capital            string  `json:"capital"`
	Subregion          string  `json:"subregion"`
	Region             string  `json:"region"`
	Population         int     `json:"population"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Demonym            string  `json:"demonym"`
	Area               float64 `json:"area"`
	Gini               float64 `json:"gini"`
	Timezones          string  `json:"timezones"`
	NativeName         string  `json:"nativeName"`
	NumericCode        string  `json:"numericCode"`
	FlagSVG            string  `json:"flag_svg"`
	FlagPNG            string  `json:"flag_png"`
	CurrencyCode       string  `json:"currency_code"`
	CurrencyName       string  `json:"currency_name"`
	CurrencySymbol     string  `json:"currency_symbol"`
	LanguageCode       string  `json:"language_code"`
	LanguageName       string  `json:"language_name"`
	LanguageNativeName string  `json:"language_nativeName"`
	IsIndependent      bool    `json:"isIndependent"`
}
type CountryResponse struct {
	Countries []Country `json:"countries"`
}

type TambonData struct {
	ID                      int    `json:"id"`
	TambonID                string `json:"TambonID"`
	TambonThai              string `json:"TambonThai"`
	TambonEng               string `json:"TambonEng"`
	TambonThaiShort         string `json:"TambonThaiShort"`
	TambonEngShort          string `json:"TambonEngShort"`
	DistrictID              string `json:"DistrictID"`
	DistrictThai            string `json:"DistrictThai"`
	DistrictEng             string `json:"DistrictEng"`
	DistrictThaiShort       string `json:"DistrictThaiShort"`
	DistrictEngShort        string `json:"DistrictEngShort"`
	ProvinceID              string `json:"ProvinceID"`
	ProvinceThai            string `json:"ProvinceThai"`
	ProvinceEng             string `json:"ProvinceEng"`
	OfficialRegions         string `json:"official_regions"`
	FourMainRegions         string `json:"four_main_regions"`
	TouristRegions          string `json:"tourist_regions"`
	GreaterBangkokProvinces string `json:"greater_bangkok_provinces"`
	PostalCodeRemark        string `json:"PostalCodeRemark"`
	PostCodeMain            string `json:"PostCodeMain"`
	PostCodeAll             string `json:"PostCodeAll"`
}
type TambonDataResponse struct {
	ProvinceAmphoeTambonZipcode []TambonData `json:"provinceAmphoeTambonZipcode"`
}

