package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/xgxw/xgxw-go/internal/controllers"
	"github.com/xgxw/xgxw-go/internal/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "http server",
	Long:  `服务端`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := loadApplocationOps()

		boot, err := newBootstrap(opts)
		handleInitError("bootstarp", err)
		fileCtrl := controllers.NewFileController(boot.Logger, boot.FileSvc)

		e := echo.New()
		jwtMiddleware := middlewares.NewJWTMiddlewares(boot.Logger, boot.Options.Auth)
		jwtMiddlewareFunc := jwtMiddleware.MiddlewareFunc
		e.Use(middleware.Logger())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     boot.Options.Server.CorsAllowOrigin,
			AllowCredentials: true,
		}))

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "enjoy yourself!")
		})

		v1 := e.Group("/v1")

		{
			v1.POST("/auth/token", jwtMiddleware.HandlerFunc)
		}

		{
			file := v1.Group("/file")
			file.GET("/public/*", fileCtrl.Get)
			file.GET("/*", fileCtrl.Get, jwtMiddlewareFunc)
			file.POST("/*", fileCtrl.Put, jwtMiddlewareFunc)
			file.PUT("/*", fileCtrl.Put, jwtMiddlewareFunc)
			file.DELETE("/*", fileCtrl.Del, jwtMiddlewareFunc)
		}

		{
			catalog := v1.Group("/catalog")
			catalog.GET("/public/*", fileCtrl.GetCatalog)
			catalog.GET("/*", fileCtrl.GetCatalog, jwtMiddlewareFunc)
		}

		quit := make(chan os.Signal, 1)
		go func() {
			// 当程序较多/HTTP设置较多时, 可以单独封装Server组件, 在组件内计算这些值
			address := fmt.Sprintf("%s:%d", opts.Server.HTTP.Host, opts.Server.HTTP.Port)
			err = e.Start(address)
			if err != nil {
				boot.Logger.Fatal("start echo error, error is ", err)
				quit <- os.Interrupt
			}
		}()
		signal.Notify(quit, os.Interrupt)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

type (
	// HTTPOps is ...
	HTTPOps struct {
		Host string `mapstructure:"host" yaml:"host"`
		Port uint   `mapstructure:"port" yaml:"port"`
	}
	// ServerOps is ...
	ServerOps struct {
		HTTP            HTTPOps  `mapstructure:"http" yaml:"http"`
		CorsAllowOrigin []string `mapstructure:"cors_allow_origin"`
	}
)
