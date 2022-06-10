package app

import (
	authdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/delegate"
	authgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/gateway"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	authusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/logger"
	prjdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/delegate"
	prjgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/gateway"
	prjhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/http"
	prjusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/server"
	"go.uber.org/fx"
	"log"
)

func InvokeWith(options ...fx.Option) *fx.App {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	var di = []fx.Option{
		fx.Provide(logger.NewLogger),
		fx.Provide(db_client.NewPostgresClient),
		fx.Provide(authgateway.SetupAuthGateway),
		fx.Provide(prjgateway.SetupProjectsGateway),
		fx.Provide(authusecase.SetupAuthUseCase),
		fx.Provide(prjusecase.SetupProjectUseCase),
		fx.Provide(authdelegate.SetupAuthDelegate),
		fx.Provide(prjdelegate.SetupProjectDelegate),
		fx.Provide(prjhttp.NewProjectsHandler),
		fx.Provide(authhttp.NewAuthHandler),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	InvokeWith(fx.Invoke(server.NewServer)).Run()
}
