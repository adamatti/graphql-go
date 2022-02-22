package graphql

import (
	"github.com/adamatti/graphql-go/db"
	"github.com/graphql-go/graphql"
)

type stateFilter func (db.State) bool

var stateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "State",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var stateQuery = &graphql.Field{
	Type:        graphql.NewList(stateType),
	Description: "List of states",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// FIXME move filters to DB layer
		list:= db.GetStateList()

		idQuery, isOK := params.Args["id"].(string)		
		if isOK {
			list = filterState(list, func (state db.State) bool {
				return state.ID == idQuery
			})
		}

		country, isOK := params.Source.(db.Country)
		if isOK {
			list = filterState(list, func (state db.State) bool {
				return state.CountryID == country.ID
			})
		}

		return list, nil
	},
}

func filterState(states []db.State, f stateFilter) []db.State {
	var filtered []db.State
	for _, state := range states {
		if f(state) {
			filtered = append(filtered, state)
		}
	}
	return filtered
}

