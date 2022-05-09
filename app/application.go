package app

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/delegate"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/gateway"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/handlers"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/logger"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/server"
	"go.uber.org/fx"
)

func AppInvokeWith(options ...fx.Option) *fx.App {
	var di = []fx.Option{
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewConfig),
		fx.Provide(gateway.NewPostgresClient),
		fx.Provide(gateway.Setup),
		fx.Provide(usecase.Setup),
		fx.Provide(delegate.Setup),
		fx.Provide(handlers.NewHandler),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	AppInvokeWith(fx.Invoke(server.NewServer)).Run()
}
