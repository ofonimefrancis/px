package graphql

import "github.com/graphql-go/graphql"

func (f *Facade) createQueryType() *graphql.Object {
	availableFields := graphql.Fields{
		"hello": &graphql.Field{
			Type:        graphql.String,
			Description: "Say Hello World to the world",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World.", nil
			},
		},
		"user": &graphql.Field{
			Type:        UserType,
			Description: "The only one user in our API",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := User{
					Name:  "Ofonime Francis",
					Email: "ofonimeusoro01!@gmail.com",
					Phone: "0780102786",
				}
				return user, nil
			},
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: availableFields,
	})
}

func (f *Facade) createMutationType() *graphql.Object {
	availableMutations := graphql.Fields{}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: availableMutations,
	})
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone_number": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone_number"`
}
