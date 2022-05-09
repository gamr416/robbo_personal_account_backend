package app

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth/delegate"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth/gateway"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/logger"
	"github.com/skinnykaen/robbo_student_personal_account.git/server"
	"go.uber.org/fx"
)

func AppInvokeWith(options ...fx.Option) *fx.App {
	var di = []fx.Option{
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewConfig),
		fx.Provide(db_client.NewPostgresClient),
		fx.Provide(gateway.SetupAuthGateway),
		fx.Provide(usecase.SetupAuthUseCase),
		fx.Provide(delegate.SetupAuthDelegate),
		fx.Provide(http.NewAuthHandler),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	AppInvokeWith(fx.Invoke(server.NewServer)).Run()
}
