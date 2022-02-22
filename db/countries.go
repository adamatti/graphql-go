package db

var countryList = []Country{
	{"BR", "Brazil"},
	{"US", "United States"},
}

type Country struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func GetCountryList() []Country {
	return countryList
}