package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"jwt_echo_graphql/controllers"
	"jwt_echo_graphql/controllers/middlwares"
	"jwt_echo_graphql/graph/generated"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middlewares.SetAuthContext())

	// set GraphQL with Echo framework
	resolver_ := &controllers.Resolver{}
	// since graphql is also same endpoint and POST method, we set as below
	e.POST("/graphql", func(c echo.Context) error {
		h := handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{Resolvers: resolver_},
			),
		)
		// pass response and request from the echo context to graphql server
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// This endpoint allows us to test API
	e.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))

	// Start server
	log.Fatal(e.Start(":8000"))
}
