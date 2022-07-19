package app

import (
	authdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/delegate"
	authgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/gateway"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	authusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	crsdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/delegate"
	crsgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/gateway"
	crshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/http"
	crsusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	edxapiusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/logger"
	ppagedelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/delegate"
	ppagegateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/gateway"
	ppagehttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/http"
	ppageusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/usecase"
	prjdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/delegate"
	prjgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/gateway"
	prjhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/http"
	prjusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/usecase"
	usrdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/users/delegate"
	usrgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/users/gateway"
	usrhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/users/http"
	usrusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/users/usecase"
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
		fx.Provide(ppagegateway.SetupProjectPageGateway),
		fx.Provide(authusecase.SetupAuthUseCase),
		fx.Provide(prjusecase.SetupProjectUseCase),
		fx.Provide(ppageusecase.SetupProjectPageUseCase),
		fx.Provide(edxapiusecase.SetupEdxApiUseCase),
		fx.Provide(authdelegate.SetupAuthDelegate),
		fx.Provide(prjdelegate.SetupProjectDelegate),
		fx.Provide(ppagedelegate.SetupProjectPageDelegate),
		fx.Provide(prjhttp.NewProjectsHandler),
		fx.Provide(ppagehttp.NewProjectPageHandler),
		fx.Provide(authhttp.NewAuthHandler),
		fx.Provide(crsdelegate.SetupCourseDelegate),
		fx.Provide(crsgateway.SetupCoursesGateway),
		fx.Provide(crshttp.NewCoursesHandler),
		fx.Provide(crsusecase.SetupCourseUseCase),
		fx.Provide(usrdelegate.SetupUsersDelegate),
		fx.Provide(usrgateway.SetupUsersGateway),
		fx.Provide(usrhttp.NewUsersHandler),
		fx.Provide(usrusecase.SetupUsersUseCase),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	InvokeWith(fx.Invoke(server.NewServer)).Run()
}
