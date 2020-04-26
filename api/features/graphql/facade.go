package graphql

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ofonimefrancis/pixels/pkg/datastore"
	"github.com/rinosukmandityo/hexagonal-login/services"
)

const BasePath = "/graphql"

var Schema graphql.Schema

type Facade struct {
	UserService services.UserService
}

func NewGraphQLFacade(userService datastore.UserRepository) *Facade {
	userService := serviceLogic.NewUserService(userService)
	return &Facade{UserService: userService}
}

func (f *Facade) handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		result := graphql.Do(graphql.Params{
			Schema:  schema,
			Context: ctx,
		})
		json.NewEncoder(w).Encode(result)
	}
}

func (f *Facade) RegisterRoutes(r *chi.Mux) {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: f.createQueryType(),
		//Mutation: f.createMutationType(),
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.Handle("/graphiql", h)

	r.HandleFunc("/graphql", f.handler(Schema))
}
