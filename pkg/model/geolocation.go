package model

const (
	ProvinceLabel = "Province"
	DistrictLabel = "City District"
	WardLabel     = "Ward"

	LanguageEN = "en"
	LanguageVN = "vn"

	VNCode = 1
	ENCode = 2
)

type Language struct {
	Name string
	Code int
}

type GeolocateItem struct {
	Code int               `json:"code"`
	Name map[string]string `json:"name"`
}
