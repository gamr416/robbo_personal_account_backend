package app

import (
	authdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/delegate"
	authgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/gateway"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	authusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/usecase"

	chrtdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts/delegate"
	chrtgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts/gateway"
	chrthttp "github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts/http"
	chrtusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts/usecase"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	crsdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/delegate"
	crsgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/gateway"
	crshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/http"
	crsusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/usecase"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	edxusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/edx/usecase"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/logger"
	usersdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/users/delegate"
	usersgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/users/gateway"
	usershtpp "github.com/skinnykaen/robbo_student_personal_account.git/package/users/http"
	usersusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/users/usecase"

	ppagedelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/delegate"
	ppagegateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/gateway"
	ppagehttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/http"
	ppageusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/usecase"

	prjdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/delegate"
	prjgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/gateway"
	prjhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/http"
	prjusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/usecase"

	robboUnitsdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits/delegate"
	robboUnitsgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits/gateway"
	robboUnitshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits/http"
	robboUnitsusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits/usecase"

	robboGroupdelegate "github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup/delegate"
	robboGroupgateway "github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup/gateway"
	robboGrouphttp "github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup/http"
	robboGroupusecase "github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup/usecase"
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
		fx.Provide(crsgateway.SetupCoursesGateway),
		fx.Provide(chrtgateway.SetupCohortsGateway),
		fx.Provide(usersgateway.SetupUsersGateway),
		fx.Provide(robboUnitsgateway.SetupRobboGroupGateway),
		fx.Provide(robboGroupgateway.SetupRobboGroupGateway),
		fx.Provide(authusecase.SetupAuthUseCase),
		fx.Provide(prjusecase.SetupProjectUseCase),
		fx.Provide(ppageusecase.SetupProjectPageUseCase),
		fx.Provide(crsusecase.SetupCourseUseCase),
		fx.Provide(chrtusecase.SetupCohortUseCase),
		fx.Provide(edxusecase.SetupEdxApiUseCase),
		fx.Provide(usersusecase.SetupUsersUseCase),
		fx.Provide(robboUnitsusecase.SetupRobboUnitsUseCase),
		fx.Provide(robboGroupusecase.SetupRobboGroupUseCase),
		fx.Provide(authdelegate.SetupAuthDelegate),
		fx.Provide(prjdelegate.SetupProjectDelegate),
		fx.Provide(ppagedelegate.SetupProjectPageDelegate),
		fx.Provide(crsdelegate.SetupCourseDelegate),
		fx.Provide(chrtdelegate.SetupCohortDelegate),
		fx.Provide(usersdelegate.SetupUsersDelegate),
		fx.Provide(robboUnitsdelegate.SetupRobboUnitsDelegate),
		fx.Provide(robboGroupdelegate.SetupRobboGroupDelegate),
		fx.Provide(prjhttp.NewProjectsHandler),
		fx.Provide(ppagehttp.NewProjectPageHandler),
		fx.Provide(authhttp.NewAuthHandler),
		fx.Provide(crshttp.NewCoursesHandler),
		fx.Provide(chrthttp.NewCohortsHandler),
		fx.Provide(usershtpp.NewUsersHandler),
		fx.Provide(robboUnitshttp.NewRobboUnitsHandler),
		fx.Provide(robboGrouphttp.NewRobboGroupHandler),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	InvokeWith(fx.Invoke(server.NewServer)).Run()
}
