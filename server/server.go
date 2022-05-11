package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(lifecycle fx.Lifecycle, handler authhttp.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				router := gin.Default()
				router.Use(
					gin.Recovery(),
					gin.Logger(),
				)
				handler.InitAuthRoutes(router)
				server := &http.Server{
					Addr:           viper.GetString("server.address"),
					Handler:        cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowCredentials: true}).Handler(router),
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
