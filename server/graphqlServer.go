package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/skinnykaen/robbo_student_personal_account.git/app/modules"
	"github.com/skinnykaen/robbo_student_personal_account.git/graph/generated"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func NewGraphqlServer(lifecycle fx.Lifecycle, graphQLModule modules.GraphQLModule) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				port := viper.GetString("graphqlServer.port")

				srv := handler.NewDefaultServer(
					generated.NewExecutableSchema(generated.Config{
						Resolvers: &graphQLModule.UsersResolver,
					}),
				)

				http.Handle("/", playground.Handler("GraphQL playground", "/query"))
				http.Handle("/query", srv)

				log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
				go func() {
					if err = http.ListenAndServe(":"+port, nil); err != nil {
						log.Fatalf("Failed to listen adn serve")
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})

}
