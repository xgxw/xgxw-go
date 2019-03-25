package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/everywan/xgxw/internal/controllers"
	"github.com/everywan/xgxw/internal/middlewares"
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
		todoController := controllers.NewTodoController(boot.FileSvc, boot.TodoSvc)

		e := echo.New()
		jwtMiddleware := middlewares.NewJWTMiddlewares(boot.Logger, boot.Options.Auth)
		e.Use(middleware.Logger())

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "enjoy yourself!")
		})

		v1 := e.Group("/v1")

		{
			todo := v1.Group("/todo", jwtMiddleware)
			todo.GET("/:id", todoController.GetTodo)
			todo.GET("/s", todoController.GetTodos)
		}

		go func() {
			// 当程序较多/HTTP设置较多时, 可以单独封装Server组件, 在组件内计算这些值
			address := fmt.Sprintf("%s:%d", opts.Server.HTTP.Host, opts.Server.HTTP.Port)
			e.Start(address)
		}()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)
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
		HTTP HTTPOps `mapstructure:"http" yaml:"http"`
	}
)
