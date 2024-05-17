package main

import (
	"log"
	"net/http"

	"github.com/eulabs/back-end/api-product/internal/config"
	"github.com/eulabs/back-end/api-product/internal/config/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	hand_prd "github.com/eulabs/back-end/api-product/internal/handler/product"
	"github.com/eulabs/back-end/api-product/pkg/adapter/mysql"
	"github.com/eulabs/back-end/api-product/pkg/server"
	service_prd "github.com/eulabs/back-end/api-product/pkg/service/product"
)

var (
	VERSION = "0.1.0-dev"
	COMMIT  = "ABCDEFG-dev"
)

func main() {
	logger.Info("start Notification application")
	conf := config.NewConfig()
	conn_mysql := mysql.New(conf)

	prd_service := service_prd.NewProductService(conn_mysql)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", healthcheck)
	hand_prd.RegisterProductAPIHandlers(e, prd_service)

	srv := server.NewHTTPServer(e, conf)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Printf("Server Run on [Port: %s], [Mode: %s], [Version: %s], [Commit: %s]", conf.PORT, conf.Mode, VERSION, COMMIT)

	select {}
}

func healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"MSG":    "Server Ok",
		"codigo": 200,
	})
}
