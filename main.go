package main

import (
	"net/http"

	"github.com/ofonimefrancis/pixels/graphql"
)

func main() {

	router := graphql.NewGraphQLFacade()

	http.ListenAndServe(":8000", nil)
}
