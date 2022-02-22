package db

var stateList = []State{
	{"BR", "RS", "Rio Grande do Sul"},
	{"BR", "SC", "Santa Catarina"},
	{"BR", "SP", "São Paulo"},
	{"US", "TX", "Texas"},
}

type State struct {
	CountryID string `json:"country_id,omitempty"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func GetStateList() []State { 
	return stateList
}