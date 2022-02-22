package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
			"latestPost": sampleField,
			"countries": countryQuery,
			"states": stateQuery,
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func GetHandle() http.Handler {
	var handle = gqlhandler.New(&gqlhandler.Config{
    Schema: &schema,
    Pretty: true,
  })
	return handle
}