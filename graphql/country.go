package graphql

import (
	"github.com/adamatti/graphql-go/db"
	"github.com/graphql-go/graphql"
)

type countryFilter func (db.Country) bool

var countryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Country",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"states": stateQuery,
	},
})

var countryQuery = &graphql.Field{
	Type:        graphql.NewList(countryType),
	Description: "List of countries",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// FIXME move filters to DB layer
		idQuery, isOK := params.Args["id"].(string)
		list:= db.GetCountryList()
		if isOK {
			list = filterCountry(list, func(country db.Country) bool {
				return country.ID == idQuery
			});
		}
		return list, nil
	},
}

func filterCountry(countries []db.Country, f countryFilter) []db.Country {
	var filtered []db.Country
	for _, country := range countries {
		if f(country) {
			filtered = append(filtered, country)
		}
	}
	return filtered
}
