package graphql

import "github.com/graphql-go/graphql"

func (f *Facade) createQueryType() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: availableFields,
	})
}
