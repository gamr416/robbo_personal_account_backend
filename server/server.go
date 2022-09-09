package server

import (
	"context"
	"github.com/skinnykaen/robbo_student_personal_account.git/app/modules"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(lifecycle fx.Lifecycle, handlers modules.HandlerModule) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				router := gin.Default()
				router.Use(
					gin.Recovery(),
					gin.Logger(),
				)
				handlers.AuthHandler.InitAuthRoutes(router)
				handlers.ProjectsHandler.InitProjectRoutes(router)
				handlers.ProjectPageHandler.InitProjectRoutes(router)
				handlers.CoursesHandler.InitCourseRoutes(router)
				handlers.CohortsHandler.InitCohortRoutes(router)
				handlers.UsersHandler.InitUsersRoutes(router)
				handlers.RobboUnitsHandler.InitRobboUnitsRoutes(router)
				handlers.RobboGroupHandler.InitRobboGroupRoutes(router)
				handlers.CoursePacketHandler.InitCoursePacketRoutes(router)
				server := &http.Server{
					Addr: viper.GetString("server.address"),
					Handler: cors.New(
						// TODO make config
						cors.Options{
							AllowedOrigins:   []string{"http://0.0.0.0:3030", "http://0.0.0.0:8601", "http://localhost:3030"},
							AllowCredentials: true,
							AllowedMethods: []string{
								http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodOptions,
							},
							//AllowedHeaders: []string{"*"},
							AllowedHeaders: []string{
								"Origin", "X-Requested-With", "Content-Type", "Accept", "Set-Cookie", "Authorization",
							},
						},
					).Handler(router),
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}
				go func() {
					if err := server.ListenAndServe(); err != nil {
						log.Fatalf("Failed to listen and serve", err)
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}
