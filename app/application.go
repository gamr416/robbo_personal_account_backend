package app

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/app/modules"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/logger"
	"github.com/skinnykaen/robbo_student_personal_account.git/server"
	"log"

	"go.uber.org/fx"
)

func InvokeWith(options ...fx.Option) *fx.App {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	var di = []fx.Option{
		fx.Provide(logger.NewLogger),
		fx.Provide(db_client.NewPostgresClient),
		fx.Provide(modules.SetupGateway),
		fx.Provide(modules.SetupUseCase),
		fx.Provide(modules.SetupDelegate),
		fx.Provide(modules.SetupHandler),
		fx.Provide(modules.SetupGraphQLModule),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	InvokeWith(
		fx.Invoke(server.NewHttpServer),
		fx.Invoke(server.NewGraphqlServer),
	).Run()
}
