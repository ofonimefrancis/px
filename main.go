package main

import "github.com/ofonimefrancis/pixels/cmd"

func main() {

	// router := chi.NewMux()
	// graphQLFacade := graphql.NewGraphQLFacade()
	// graphQLFacade.RegisterRoutes(router)

	// log.Println("Application running on http://localhost:8000")
	// http.ListenAndServe(":8000", router)

	cmd.Execute()
}
