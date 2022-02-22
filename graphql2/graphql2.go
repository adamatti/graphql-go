package graphql2

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// FIXME make array of countries work
type query struct{}

func (_ *query) Hello() string { 
	return "Hello, world!" 
}

var s = `
	type Query {
		hello: String!
	}
`
var schema = graphql.MustParseSchema(s, &query{})

func GetHandle() http.Handler {
	return &relay.Handler{Schema: schema}
}