package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	courseshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/http"
	projectpagehttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/http"
	projectshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/http"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(lifecycle fx.Lifecycle, authhandler authhttp.Handler, projecthttp projectshttp.Handler, projectpagehttp projectpagehttp.Handler, coursehttp courseshttp.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				router := gin.Default()
				router.Use(
					gin.Recovery(),
					gin.Logger(),
				)
				authhandler.InitAuthRoutes(router)
				projecthttp.InitProjectRoutes(router)
				projectpagehttp.InitProjectRoutes(router)
				coursehttp.InitCourseRoutes(router)
				server := &http.Server{
					Addr: viper.GetString("server.address"),
					Handler: cors.New(
						// TODO make config
						cors.Options{
							AllowedOrigins:   []string{"http://0.0.0.0:3030", "http://0.0.0.0:8601", "localhost:3030"},
							AllowCredentials: true,
							AllowedMethods: []string{
								"PUT", "DELETE", "GET", "OPTIONS", "POST", "HEAD",
							},
							AllowedHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
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
