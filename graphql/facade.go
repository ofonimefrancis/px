package graphql

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
)

const BasePath = "/graphql"

var Schema graphql.Schema

type Facade struct {
}

func NewGraphQLFacade() *Facade {
	return &Facade{}
}

func (f *Facade) handler(schema graphql.Schema) *chi.Mux {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		result := graphql.Do(graphql.Params{
			Schema:  schema,
			Context: ctx,
		})
	}
}

func (f *Facade) createQueryType() {

}

func (f *Facade) createMutationType() {

}

func (f *Facade) RegisterRoutes(r *chi.Mux) {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    f.createQueryType(),
		Mutation: f.createMutationType(),
	})

	if err != nil {
		panic(err)
	}

	r.Post("", f.handler(Schema))
}
